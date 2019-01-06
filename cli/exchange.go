package cli

import "github.com/urfave/cli"

const (
	exchangeName  = "exchange"
	exchangeUsage = "Shows a list of commands for currency exchange"
)

func (t *Tool) cmdExchange() cli.Command {
	return cli.Command{
		Name:  exchangeName,
		Usage: exchangeUsage,
		Subcommands: []cli.Command{
			t.cmdExchangeSeries(),
			t.cmdExchangeRates(),
		},
	}
}
