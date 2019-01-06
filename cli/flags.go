package cli

import (
	"github.com/urfave/cli"
)

func (t *Tool) flagFrom() *cli.StringFlag {
	return &cli.StringFlag{
		Name:  "from",
		Usage: "date where the series start",
		Value: defaultFrom,
	}
}

func (t *Tool) flagTo() *cli.StringFlag {
	return &cli.StringFlag{
		Name:  "to",
		Usage: "date where the series end",
		Value: defaultTo,
	}
}

func (t *Tool) flagAggregate() *cli.StringFlag {
	return &cli.StringFlag{
		Name:  "aggregate",
		Usage: "daily, weekly, monthly, quarterly or yearly",
		Value: "daily",
	}
}

func (t *Tool) flagAnalysis() *cli.StringFlag {
	return &cli.StringFlag{
		Name:  "analysis",
		Usage: "real, mean, min, max or ultimo",
		Value: "real",
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
