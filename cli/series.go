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
	seriesName  = "series"
	seriesUsage = "Lists all interest and exchange series"
)

func (t *Tool) cmdSeries() cli.Command {
	return cli.Command{
		Name:   seriesName,
		Usage:  seriesUsage,
		Action: t.actionSeries,
		Flags: []cli.Flag{
			t.flagLang(),
			t.flagGroup(),
		},
	}
}

func (t *Tool) actionSeries(c *cli.Context) error {
	ctx := context.Background()
	lang := swea.Language(c.String("lang"))

	greq := &swea.GetInterestAndExchangeGroupNamesRequest{
		Language: lang,
	}
	gres, err := t.API.GetInterestAndExchangeGroupNames(ctx, greq)
	if err != nil {
		return err
	}
	gs := c.StringSlice("group")

	for _, g := range gres.Groups {
		if len(gs) < 1 || isInSlice(gs, g.ID) {
			req := &swea.GetInterestAndExchangeNamesRequest{
				Language: lang,
				GroupID:  g.ID,
			}
			res, err := t.API.GetInterestAndExchangeNames(ctx, req)
			if err != nil {
				return err
			}
			t.renderSeries(g, res)
		}
	}
	return nil
}

func (t *Tool) renderSeries(group swea.GroupInfo, res *swea.GetInterestAndExchangeNamesResponse) {
	const (
		rowFmt = "%s\t %s\t %s\n"
	)
	w := tabwriter.NewWriter(os.Stdout, 14, 0, 1, ' ', 0)
	defer w.Flush()
	fmt.Fprintf(w, "(%s) %s:\n\n", group.ID, group.Name)
	fmt.Fprintf(w, rowFmt, "ID", "Name", "Description")
	for _, s := range res.Series {
		fmt.Fprintf(w, rowFmt, s.ID, s.Name, s.Description)
	}
	fmt.Fprintf(w, "\n")
}

func isInSlice(items []string, s string) bool {
	for _, item := range items {
		if item == s {
			return true
		}
	}
	return false
}
