package cli

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/urfave/cli"
	"github.com/zeeraw/riksbank/swea"
)

const (
	daysName  = "days"
	daysUsage = "Lists days with relevant banking information between two dates"
)

func (t *Tool) cmdDays() cli.Command {
	return cli.Command{
		Name:   daysName,
		Usage:  daysUsage,
		Action: t.actionDays,
		Flags: []cli.Flag{
			t.flagFrom(),
			t.flagTo(),
		},
	}
}

func (t *Tool) actionDays(c *cli.Context) error {
	ctx := context.Background()
	from, err := parseDate(c.String("from"))
	if err != nil {
		return err
	}
	to, err := parseDate(c.String("to"))
	if err != nil {
		return err
	}
	res, err := t.API.GetCalendarDays(ctx, &swea.GetCalendarDaysRequest{
		From: from,
		To:   to,
	})
	if err != nil {
		return nil
	}
	t.renderDays(res)
	return nil
}

func (t *Tool) renderDays(res *swea.GetCalendarDaysResponse) {
	fmt.Fprintf(os.Stdout, "Showing days between %s and %s\n\n", res.From.String(), res.To.String())
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintf(w, "Date\t Year\t Week\t Bank day\n")
	for _, day := range res.Days {
		fmt.Fprintf(w, "%s\t %d\t %d\t %s\n", day.Date.String(), day.WeekYear, day.Week, boolToYesNo(day.IsBankDay))
	}
	defer w.Flush()
}
