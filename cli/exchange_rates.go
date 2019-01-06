package cli

import (
	"context"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/zeeraw/riksbank"

	"github.com/urfave/cli"
	"github.com/zeeraw/riksbank/currency"
)

const (
	exchangeRatesName  = "rates"
	exchangeRatesUsage = "Lists the exchange rates between two dates"
)

func (t *Tool) cmdExchangeRates() cli.Command {
	return cli.Command{
		Name:   exchangeRatesName,
		Usage:  exchangeRatesUsage,
		Action: t.actionExchangeRates,
		Flags: []cli.Flag{
			t.flagFrom(),
			t.flagTo(),
			t.flagAggregate(),
			t.flagCurrency(),
		},
	}
}

func (t *Tool) actionExchangeRates(c *cli.Context) error {
	ctx := context.Background()
	cs := c.StringSlice("currency")
	if len(cs) <= 0 {
		return fmt.Errorf("need to provide at least one currency pair")
	}
	pairs := make([]currency.Pair, len(cs))
	for idx, c := range cs {
		pairs[idx] = currency.ParsePair(c)
	}
	from, err := parseDate(c.String("from"))
	if err != nil {
		return err
	}
	to, err := parseDate(c.String("to"))
	if err != nil {
		return err
	}
	method, err := riksbank.ParseAggregate(c.String("aggregate"))
	if err != nil {
		return err
	}
	req := &riksbank.ExchangeRatesRequest{
		CurrencyPairs:   pairs,
		AggregateMethod: method,
		From:            from,
		To:              to,
	}
	res, err := t.Riksbank.ExchangeRates(ctx, req)
	if err != nil {
		return err
	}
	return t.renderExchangeRates(req, res)
}

func (t *Tool) renderExchangeRates(req *riksbank.ExchangeRatesRequest, res *riksbank.ExchangeRatesResponse) error {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	defer w.Flush()
	fmt.Fprintf(os.Stdout, "Ranging from %s to %s\n", formatDate(req.From), formatDate(req.To))
	if req.AggregateMethod != riksbank.Daily {
		fmt.Fprintf(os.Stdout, "Aggregating %s avarage\n", req.AggregateMethod.Name())
	}
	var pairs = make([]string, len(req.CurrencyPairs))
	for idx, cp := range req.CurrencyPairs {
		pairs[idx] = cp.String()
	}
	fmt.Fprintf(os.Stdout, "Series %s\n", strings.Join(pairs, ", "))
	fmt.Fprint(os.Stdout, "\n")
	fmt.Fprintf(w, "Period\t Base currency\t Counter currency\t Exchange rate\n")
	for _, rate := range res.ExchangeRates {
		fmt.Fprintf(w, "%s\t %s\t %s\t %s\n", rate.Period, rate.Base, rate.Counter, formatFloat(rate.Value))
	}
	return nil
}
