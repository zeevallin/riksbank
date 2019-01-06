package cli

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/zeeraw/riksbank"

	"github.com/urfave/cli"
)

const (
	exchangeCurrenciesName  = "currencies"
	exchangeCurrenciesUsage = "Lists currencies available for currency exchange rates to SEK"
)

func (t *Tool) cmdExchangeCurrencies() cli.Command {
	return cli.Command{
		Name:   exchangeCurrenciesName,
		Usage:  exchangeCurrenciesUsage,
		Action: t.actionExchangeCurrencies,
	}
}

func (t *Tool) actionExchangeCurrencies(c *cli.Context) error {
	req := &riksbank.ExchangeCurrenciesRequest{}
	res, err := t.Riksbank.ExchangeCurrencies(context.Background(), req)
	if err != nil {
		return err
	}
	t.renderExchangeCurrencies(req, res)
	return nil
}

func (t *Tool) renderExchangeCurrencies(req *riksbank.ExchangeCurrenciesRequest, res *riksbank.ExchangeCurrenciesResponse) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	defer w.Flush()
	fmt.Fprintf(w, "ID\t Name\t Description\n")
	for _, ec := range res.Currencies {
		fmt.Fprintf(w, "%s\t %s\t %s\n", ec.SeriesID, ec.Code, ec.Description)
	}
}
