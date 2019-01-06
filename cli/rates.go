package cli

import (
	"context"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"cloud.google.com/go/civil"
	"github.com/urfave/cli"
	"github.com/zeeraw/riksbank/swea"
)

const (
	ratesName  = "rates"
	ratesUsage = "Lists interest or exchange rates between two dates"
)

func (r *runner) cmdRates() cli.Command {
	return cli.Command{
		Name:   ratesName,
		Usage:  ratesUsage,
		Action: r.actionRates,
		Flags: []cli.Flag{
			r.flagFrom(),
			r.flagTo(),
			r.flagLang(),
			r.flagAggregate(),
			r.flagAnalysis(),
			r.flagSeries(),
		},
	}
}

func (r *runner) actionRates(c *cli.Context) error {
	ctx := context.Background()
	ss := c.StringSlice("series")
	if len(ss) < 1 {
		return fmt.Errorf("need to have at least one series")
	}
	series := make([]swea.SearchGroupSeries, len(ss))
	for idx, s := range ss {
		series[idx] = swea.SearchGroupSeries{
			SeriesID: s,
			GroupID:  "1", // Set this to 1 until we're able to fetch series and groups
		}
	}
	from, err := civil.ParseDate(c.String("from"))
	if err != nil {
		return err
	}
	to, err := civil.ParseDate(c.String("to"))
	if err != nil {
		return err
	}
	aggregate, err := swea.ParseAggregate(c.String("aggregate"))
	if err != nil {
		return err
	}
	analysis, err := swea.ParseAnalysisForAggregate(c.String("analysis"), aggregate)
	if err != nil {
		return err
	}
	req := &swea.GetInterestAndExchangeRatesRequest{
		From:            from,
		To:              to,
		Language:        swea.Language(r.lang),
		AggregateMethod: aggregate,
		Series:          series,
	}
	switch analysis {
	case swea.Mean:
		req.Average = true
	case swea.Min:
		req.Min = true
	case swea.Max:
		req.Max = true
	case swea.Ultimo:
		req.Ultimo = true
	}

	res, err := r.api.GetInterestAndExchangeRates(ctx, req)
	if err != nil {
		return err
	}

	return r.renderRates(analysis, res)
}

func (r *runner) renderRates(analysis swea.AnalysisMethod, res *swea.GetInterestAndExchangeRatesResponse) error {
	const (
		rowFmt = "%s\t %s\t %s\t %s\n"
	)
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	defer w.Flush()
	fmt.Fprintf(os.Stdout, "Ranging from %s to %s\n", res.From.String(), res.To.String())
	if res.AggregateMethod != swea.Daily {
		fmt.Fprintf(os.Stdout, "Aggregating %s %s value\n\n", swea.AggregateName(res.AggregateMethod), analysis)
	}
	var series = make([]string, len(res.Series))
	for idx, s := range res.Series {
		series[idx] = s.SeriesID
	}
	fmt.Fprintf(os.Stdout, "Series %s\n", strings.Join(series, ", "))
	fmt.Fprint(os.Stdout, "\n")

	var valueLabel string
	switch analysis {
	case swea.Mean:
		valueLabel = "Mean value"
	case swea.Min:
		valueLabel = "Min value"
	case swea.Max:
		valueLabel = "Max value"
	case swea.Ultimo:
		valueLabel = "Ultimo value"
	default:
		valueLabel = "Value"
	}
	fmt.Fprintf(w, rowFmt, "Period", "Series ID", "Series Name", valueLabel)
	for _, rate := range res.Rates {
		var value string
		switch analysis {
		case swea.Mean:
			value = rate.Average
		case swea.Min:
			value = rate.Min
		case swea.Max:
			value = rate.Max
		case swea.Ultimo:
			value = rate.Ultimo
		default:
			value = rate.Value
		}
		fmt.Fprintf(w, rowFmt, rate.Period, rate.SeriesID, rate.SeriesName, value)
	}
	return nil
}
