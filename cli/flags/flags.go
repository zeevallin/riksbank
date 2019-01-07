package flags

import (
	"time"

	"github.com/urfave/cli"
)

const (
	dateLayout = "2006-01-02"
)

var (
	today = time.Now()
	// defaultTo should be today
	defaultTo = today.Format(dateLayout)
	// defaultFrom should be seven days ago
	defaultFrom = today.AddDate(0, 0, -7).Format(dateLayout)
)

var (
	// From represents a flag for retrieving a starting date
	From = &cli.StringFlag{
		Name:  "from",
		Usage: "date where the series start",
		Value: defaultFrom,
	}

	// To represents a flag for retrieving an end date
	To = &cli.StringFlag{
		Name:  "to",
		Usage: "date where the series end",
		Value: defaultTo,
	}

	// Aggregate represents a flag for picking an aggregate method
	Aggregate = &cli.StringFlag{
		Name:  "aggregate",
		Usage: "daily, weekly, monthly, quarterly or yearly",
		Value: "daily",
	}

	// Analysis represents a flag for picking an analysis method
	Analysis = &cli.StringFlag{
		Name:  "analysis",
		Usage: "real, mean, min, max or ultimo",
		Value: "real",
	}

	// Currency represents a flag for adding a currency pair
	Currency = &cli.StringSliceFlag{
		Name:  "currency, c",
		Usage: "eg. GBP",
	}

	// Series represents a flag for adding a series
	Series = &cli.StringSliceFlag{
		Name:  "series, s",
		Usage: "series id (eg. SETB1MBENCHC)",
	}

	// Group represents a flag for adding a group
	Group = &cli.StringSliceFlag{
		Name:  "group, g",
		Usage: "group id (eg. 1), if none provided all groups will be shown",
	}
)
