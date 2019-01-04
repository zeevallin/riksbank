package main

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"cloud.google.com/go/civil"
	"github.com/urfave/cli"
	"github.com/zeeraw/riksbank/swea"
)

func (r *runner) cmdCalendarDays() cli.Command {
	return cli.Command{
		Name:   "calendar_days",
		Action: r.actionCalnedarDays,
		Flags: []cli.Flag{
			r.flagFrom(),
			r.flagTo(),
		},
	}
}

func (r *runner) actionCalnedarDays(c *cli.Context) error {
	ctx := context.Background()
	res, err := r.api.GetCalendarDays(ctx, &swea.GetCalendarDaysRequest{
		From: civil.Date{Year: 2018, Month: time.January, Day: 1},
		To:   civil.Date{Year: 2018, Month: time.February, Day: 1},
	})
	if err != nil {
		return nil
	}
	r.renderCalendarDays(res)
	return nil
}

func (r *runner) renderCalendarDays(res *swea.GetCalendarDaysResponse) {
	fmt.Fprintf(os.Stdout, "Showing calendar days for dates:\n%s to %s\n\n", res.From.String(), res.To.String())
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintf(w, "Date\t Year\t Week\t Bank day\n")
	for _, day := range res.Days {
		fmt.Fprintf(w, "%s\t %d\t %d\t %s\n", day.Date.String(), day.WeekYear, day.Week, boolToYesNo(day.IsBankDay))
	}
	defer w.Flush()
}
