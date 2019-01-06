package cli

import (
	"fmt"

	"github.com/zeeraw/riksbank"

	"github.com/urfave/cli"
)

// Tool represents the command line tool for riksbank
type Tool struct {
	Riksbank *riksbank.Riksbank
}

var (
	defaultFrom string
	defaultTo   string
)

// New returns a new Tool with the live API
func New() *Tool {
	return &Tool{
		Riksbank: riksbank.New(riksbank.Config{}),
	}
}

// Run will run the command line tool
func (t *Tool) Run(args []string) error {
	if len(args) <= 0 {
		return fmt.Errorf("should at least have one argument")
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
		t.cmdRates(),
		t.cmdSeries(),
		t.cmdGroups(),
		t.cmdDays(),
		t.cmdExchange(),
	}
	app.Before = func(c *cli.Context) error {
		return nil
	}
	return app.Run(args)
}
