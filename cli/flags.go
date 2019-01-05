package cli

import (
	"github.com/urfave/cli"
)

func (r *runner) flagFrom() *cli.StringFlag {
	return &cli.StringFlag{
		Name:        "from",
		Usage:       "date where the series start",
		Value:       defaultFrom,
		Destination: &r.from,
	}
}

func (r *runner) flagTo() *cli.StringFlag {
	return &cli.StringFlag{
		Name:        "to",
		Usage:       "date where the series end",
		Value:       defaultTo,
		Destination: &r.to,
	}
}

func (r *runner) flagLang() *cli.StringFlag {
	return &cli.StringFlag{
		Name:        "lang",
		Usage:       "en or sv",
		Destination: &r.lang,
		Value:       "en",
	}
}
