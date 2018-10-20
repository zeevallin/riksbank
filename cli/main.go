package main

import (
	"log"
	"os"
	"time"

	"github.com/urfave/cli"
	"github.com/zeeraw/riksbank/api"
	"github.com/zeeraw/riksbank/api/calc"
	"github.com/zeeraw/riksbank/api/series"
)

func main() {
	app := cli.NewApp()
	app.Name = "riksbank"
	app.Usage = ""
	app.Description = "riksbank.se command line client"
	app.Author = "Philip Vieira"

	// /exhange/gbp/day/2018-01-01
	// /exhange/gbp/days/2018-01-01..2018-02-01/average,max,min

	// /exhange/gbp/week/2018-1/average,max,min
	// /exhange/gbp/weeks/2018-1..2018-4/average,max,min

	// /exhange/gbp/month/2018-1/average,max,min,ultimo
	// /exhange/gbp/months/2018-1..2018-2/average,max,min,ultimo

	// /exhange/gbp/quarter/2018-1/average,max,min
	// /exhange/gbp/quarters/2018-1..2018-2/average,max,min

	// /exhange/gbp/year/2017/average,max,min
	// /exhange/gbp/years/2017..2018/average,max,min

	app.Action = func(c *cli.Context) error {
		req := &api.RatesRequest{
			From:   api.Date{Year: 2014, Month: time.January, Day: 1},
			To:     api.Date{Year: 2016, Month: time.February, Day: 1},
			Period: calc.Year,
			Methods: []calc.Method{
				calc.Average,
				calc.Max,
				calc.Min,
			},
			Series: []series.Series{
				series.SEKGBPPMI,
				series.SEKAUDPMI,
				series.SEDP6MSTIBOR,
			},
		}
		_, err := api.GetRates(req)

		return err
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	// r := mux.NewRouter()
	// http.ListenAndServe(":8080", r)
}
