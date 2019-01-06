package cli

import (
	"time"

	"cloud.google.com/go/civil"
	"github.com/urfave/cli"
	"github.com/zeeraw/riksbank/swea"
)

type runner struct {
	api *swea.Swea

	lang      string
	from      string
	to        string
	aggregate string
	method    string
}

var (
	defaultFrom string
	defaultTo   string
)

func init() {
	today := civil.DateOf(time.Now())
	// Default to should be today
	defaultTo = today.String()
	// Default from should be seven days ago
	defaultFrom = today.AddDays(-7).String()
}

// Run will run the command line tool
func Run(args []string) error {
	r := &runner{
		api: swea.New(swea.Config{}),
	}
	app := cli.NewApp()
	app.Name = "riksbank"
	app.Usage = ""
	app.UsageText = "riksbank [global options] command [command options]"
	app.Version = "0.0.1"
	app.Action = nil
	app.Description = "Command line client for the swedish central bank (riksbank.se)"
	app.Author = "Philip Vieira"
	app.Email = "zee@vall.in"
	app.Commands = []cli.Command{
		r.cmdRates(),
		r.cmdDays(),
		r.cmdSeries(),
		r.cmdGroups(),
		r.cmdExchanges(),
		r.cmdExchangeRates(),
	}
	app.Before = func(c *cli.Context) error {
		return nil
	}
	return app.Run(args)
}
