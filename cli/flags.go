package cli

import (
	"github.com/urfave/cli"
)

func (t *Tool) flagFrom() *cli.StringFlag {
	return &cli.StringFlag{
		Name:        "from",
		Usage:       "date where the series start",
		Value:       defaultFrom,
		Destination: &t.from,
	}
}

func (t *Tool) flagTo() *cli.StringFlag {
	return &cli.StringFlag{
		Name:        "to",
		Usage:       "date where the series end",
		Value:       defaultTo,
		Destination: &t.to,
	}
}

func (t *Tool) flagLang() *cli.StringFlag {
	return &cli.StringFlag{
		Name:        "lang",
		Usage:       "en or sv",
		Destination: &t.lang,
		Value:       "en",
	}
}

func (t *Tool) flagAggregate() *cli.StringFlag {
	return &cli.StringFlag{
		Name:        "aggregate",
		Usage:       "daily, weekly, monthly, quartery or yearly",
		Destination: &t.aggregate,
		Value:       "daily",
	}
}

func (t *Tool) flagAnalysis() *cli.StringFlag {
	return &cli.StringFlag{
		Name:        "analysis",
		Usage:       "real, mean, min, max or ultimo",
		Destination: &t.method,
		Value:       "real",
	}
}

func (t *Tool) flagCurrency() *cli.StringSliceFlag {
	return &cli.StringSliceFlag{
		Name:  "currency, c",
		Usage: "eg. GBP",
	}
}

func (t *Tool) flagSeries() *cli.StringSliceFlag {
	return &cli.StringSliceFlag{
		Name:  "series, s",
		Usage: "series id (eg. SETB1MBENCHC)",
	}
}
func (t *Tool) flagGroup() *cli.StringSliceFlag {
	return &cli.StringSliceFlag{
		Name:  "group, g",
		Usage: "group id (eg. 1), if none provided all groups will be shown",
	}
}
