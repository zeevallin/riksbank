package cli

import "github.com/urfave/cli"

const (
	exchangeName  = "exchange"
	exchangeUsage = "Shows a list of commands for currency exchange"
)

func (r *runner) cmdExchange() cli.Command {
	return cli.Command{
		Name:  exchangeName,
		Usage: exchangeUsage,
		Subcommands: []cli.Command{
			r.cmdExchangeSeries(),
			r.cmdExchangeRates(),
		},
	}
}
