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

func (r *runner) cmdExchangeRates() cli.Command {
	return cli.Command{
		Name:   "exchange_rates",
		Usage:  "Lists the exchange rates per day",
		Action: r.actionExchangeRates,
		Flags: []cli.Flag{
			r.flagFrom(),
			r.flagTo(),
			r.flagLang(),
			r.flagAggregate(),
			r.flagCurrency(),
		},
	}
}

func (r *runner) actionExchangeRates(c *cli.Context) error {
	ctx := context.Background()
	cs := c.StringSlice("currency")
	if len(cs) <= 0 {
		return fmt.Errorf("need to provide at least one currency pair")
	}
	pairs := make([]swea.CrossPair, len(cs))
	for idx, c := range cs {
		pairs[idx] = swea.ParseCurrencyPair(c).ToCrossPair()
	}
	from, err := civil.ParseDate(r.from)
	if err != nil {
		return err
	}
	to, err := civil.ParseDate(r.to)
	if err != nil {
		return err
	}
	method, err := swea.ParseAggregate(r.aggregate)
	if err != nil {
		return err
	}
	req := &swea.GetCrossRatesRequest{
		CrossPairs:      pairs,
		From:            from,
		To:              to,
		Language:        swea.Language(r.lang),
		AggregateMethod: method,
	}
	res, err := r.api.GetCrossRates(ctx, req)
	if err != nil {
		return err
	}

	r.renderExchangeRates(res)

	return nil
}

func (r *runner) renderExchangeRates(res *swea.GetCrossRatesResponse) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	defer w.Flush()
	fmt.Fprintf(os.Stdout, "Ranging from %s to %s\n", res.From.String(), res.To.String())
	if res.AggregateMethod != swea.Daily {
		fmt.Fprintf(os.Stdout, "Aggregating %s avarage\n", swea.AggregateName(res.AggregateMethod))
	}
	var pairs = make([]string, len(res.CrossPairs))
	for idx, cp := range res.CrossPairs {
		pairs[idx] = cp.ToCurrencyPair().String()
	}
	fmt.Fprintf(os.Stdout, "Series %s\n", strings.Join(pairs, ", "))
	fmt.Fprint(os.Stdout, "\n")
	fmt.Fprintf(w, "Period\t Base currency\t Counter currency\t Exchange rate\n")
	for _, rate := range res.CrossRates {
		fmt.Fprintf(w, "%s\t %s\t %s\t %s\n", rate.Period, rate.Base, rate.Counter, rate.Value)
	}
}
