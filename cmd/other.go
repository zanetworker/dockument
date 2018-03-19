package main

import (
	"errors"
	"fmt"
	"io"

	"github.com/spf13/cobra"
	"github.com/zanetworker/dockument/pkg/labels"
	"github.com/zanetworker/dockument/pkg/utils"
)

var otherCmdDesc = `
This command is used to fetch all other non-DOCKument labels out of Dockerfiles Dockerfiles`

type otherCmd struct {
	dockerfile string
}

func newOtherCmd(out io.Writer) *cobra.Command {
	otherCmdParams := &otherCmd{}
	dockerCmd := &cobra.Command{
		Use:   "other",
		Short: "fetch all other non-DOCKument labels out of Dockerfiles Dockerfile",
		Long:  otherCmdDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			return otherCmdParams.run()
		},
	}

	f := dockerCmd.Flags()
	f.StringVar(&otherCmdParams.dockerfile, "dockerfile", "", "the path of the Dockerfile to Dockument")

	return dockerCmd
}
func (d *otherCmd) run() error {
	if len(d.dockerfile) != 0 {
		err := printOthers(d.dockerfile)
		return err
	}
	return errors.New(utils.ColorString("red", "Please specfiy a path for the dockerfile to Dockument"))
}

func printOthers(dockerfile string) error {
	others, err := labels.GetOtherTags(dockerfile)
	fmt.Println(utils.ColorString("blue", "### Misc Labels ###"))
	for other, value := range *others {
		fmt.Printf("	%s: %s \n", utils.ColorString("green", other), value)
	}
	return err
}
