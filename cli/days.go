package cli

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/urfave/cli"
	"github.com/zeeraw/riksbank"
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
	req := &riksbank.DaysRequest{
		From: from,
		To:   to,
	}
	res, err := t.Riksbank.Days(ctx, req)
	if err != nil {
		return nil
	}
	t.renderDays(req, res)
	return nil
}

func (t *Tool) renderDays(req *riksbank.DaysRequest, res *riksbank.DaysResponse) {
	const (
		rowFmt = "%s\t %s\t %s\n"
	)
	fmt.Fprintf(os.Stdout, "Showing days between %s and %s\n\n", formatDate(req.From), formatDate(req.To))
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintf(w, rowFmt, "Date", "Week", "Bank day")
	for _, day := range res.Days {
		fmt.Fprintf(w, rowFmt, formatDate(day.Date), formatInt(day.Week), boolToYesNo(day.IsBankDay))
	}
	defer w.Flush()
}
