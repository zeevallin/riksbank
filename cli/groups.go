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
	groupsName  = "groups"
	groupsUsage = "Lists all interest and exchange groups"
)

func (t *Tool) cmdGroups() cli.Command {
	return cli.Command{
		Name:   groupsName,
		Usage:  groupsUsage,
		Action: t.actionGroups,
	}
}

func (t *Tool) actionGroups(c *cli.Context) error {
	ctx := context.Background()
	req := &riksbank.GroupsRequest{}
	res, err := t.Riksbank.Groups(ctx, req)
	if err != nil {
		return err
	}
	return t.renderGroups(req, res)
}

func (t *Tool) renderGroups(req *riksbank.GroupsRequest, res *riksbank.GroupsResponse) error {
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
