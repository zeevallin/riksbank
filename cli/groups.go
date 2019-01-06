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

func (t *Tool) cmdGroups() cli.Command {
	return cli.Command{
		Name:   groupsName,
		Usage:  groupsUsage,
		Action: t.actionGroups,
		Flags: []cli.Flag{
			t.flagLang(),
		},
	}
}

func (t *Tool) actionGroups(c *cli.Context) error {
	ctx := context.Background()
	req := &swea.GetInterestAndExchangeGroupNamesRequest{
		Language: swea.Language(t.lang),
	}

	res, err := t.API.GetInterestAndExchangeGroupNames(ctx, req)
	if err != nil {
		return err
	}

	return t.renderGroups(res)
}

func (t *Tool) renderGroups(res *swea.GetInterestAndExchangeGroupNamesResponse) error {
	const (
		rowFmt = "%s\t %s\n"
	)
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight)
	defer w.Flush()
	fmt.Fprintf(w, rowFmt, "ID", "Name")
	for _, group := range res.Groups {
		fmt.Fprintf(w, rowFmt, group.ID, group.Name)
	}

	return nil
}
