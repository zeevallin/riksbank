package cli

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/zeeraw/riksbank"

	"github.com/urfave/cli"
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
			t.flagGroup(),
		},
	}
}

func (t *Tool) actionSeries(c *cli.Context) error {
	req := &riksbank.SeriesRequest{
		Groups: c.StringSlice("group"),
	}
	res, err := t.Riksbank.Series(context.Background(), req)
	if err != nil {
		return err
	}
	for _, group := range res.Groups {
		t.renderSeries(group)
	}
	return nil
}

func (t *Tool) renderSeries(sg riksbank.SeriesGroup) {
	const (
		rowFmt = "%s\t %s\t %s\n"
	)
	w := tabwriter.NewWriter(os.Stdout, 14, 0, 1, ' ', 0)
	defer w.Flush()
	fmt.Fprintf(w, "[%s] %s:\n\n", sg.Group.ID, sg.Group.Name)
	fmt.Fprintf(w, rowFmt, "ID", "Name", "Description")
	for _, s := range sg.Series {
		fmt.Fprintf(w, rowFmt, s.ID, s.Name, s.Description)
	}
	fmt.Fprintf(w, "\n")
}
