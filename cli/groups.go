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
	groupsName  = "groups"
	groupsUsage = "Lists all interest and exchange groups"
)

func (r *runner) cmdGroups() cli.Command {
	return cli.Command{
		Name:   groupsName,
		Usage:  groupsUsage,
		Action: r.actionGroups,
		Flags: []cli.Flag{
			r.flagLang(),
		},
	}
}

func (r *runner) actionGroups(c *cli.Context) error {
	ctx := context.Background()
	req := &swea.GetInterestAndExchangeGroupNamesRequest{
		Language: swea.Language(r.lang),
	}

	res, err := r.api.GetInterestAndExchangeGroupNames(ctx, req)
	if err != nil {
		return err
	}

	return r.renderGroups(res)
}

func (r *runner) renderGroups(res *swea.GetInterestAndExchangeGroupNamesResponse) error {
	const (
		rowFmt = "%s\t %s\n"
	)
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	defer w.Flush()
	fmt.Fprintf(w, rowFmt, "ID", "Name")
	for _, group := range res.Groups {
		fmt.Fprintf(w, rowFmt, group.ID, group.Name)
	}

	return nil
}
