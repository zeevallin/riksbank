package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
	"github.com/zeeraw/riksbank/swea"
)

type runner struct {
	api *swea.Swea

	lang string
	from string
	to   string
}

func main() {
	r := &runner{
		api: swea.New(swea.Config{}),
	}
	app := cli.NewApp()
	app.Name = "riksbank"
	app.Usage = ""
	app.Description = "riksbank.se command line client"
	app.Author = "Philip Vieira"
	app.Commands = []cli.Command{
		r.cmdCalendarDays(),
		r.cmdCrossNames(),
	}
	app.Before = func(c *cli.Context) error {
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
