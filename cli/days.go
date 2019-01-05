package cli

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"

	"cloud.google.com/go/civil"
	"github.com/urfave/cli"
	"github.com/zeeraw/riksbank/swea"
)

func (r *runner) cmdDays() cli.Command {
	return cli.Command{
		Name:   "days",
		Usage:  "Lists days and information about them between two dates",
		Action: r.actionDays,
		Flags: []cli.Flag{
			r.flagFrom(),
			r.flagTo(),
		},
	}
}

func (r *runner) actionDays(c *cli.Context) error {
	ctx := context.Background()
	from, err := civil.ParseDate(r.from)
	if err != nil {
		return err
	}
	to, err := civil.ParseDate(r.to)
	if err != nil {
		return err
	}
	res, err := r.api.GetCalendarDays(ctx, &swea.GetCalendarDaysRequest{
		From: from,
		To:   to,
	})
	if err != nil {
		return nil
	}
	r.renderDays(res)
	return nil
}

func (r *runner) renderDays(res *swea.GetCalendarDaysResponse) {
	fmt.Fprintf(os.Stdout, "Showing calendar days for dates:\n%s to %s\n\n", res.From.String(), res.To.String())
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintf(w, "Date\t Year\t Week\t Bank day\n")
	for _, day := range res.Days {
		fmt.Fprintf(w, "%s\t %d\t %d\t %s\n", day.Date.String(), day.WeekYear, day.Week, boolToYesNo(day.IsBankDay))
	}
	defer w.Flush()
}
