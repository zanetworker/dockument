package main

import (
	"errors"
	"fmt"
	"io"

	"github.com/spf13/cobra"
	"github.com/zanetworker/dockument/pkg/labels"
	"github.com/zanetworker/dockument/pkg/utils"
)

var tagsCmdDesc = `
This command is used to fetch tags out of Dockerfiles Dockerfiles`

type tagsCmd struct {
	dockerfile string
}

func newTagsCmd(out io.Writer) *cobra.Command {
	tagCmdParams := &tagsCmd{}
	dockerCmd := &cobra.Command{
		Use:   "tags",
		Short: "fetches tags from the Dockerfile",
		Long:  tagsCmdDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			return tagCmdParams.run()
		},
	}

	f := dockerCmd.Flags()
	f.StringVar(&tagCmdParams.dockerfile, "dockerfile", "", "the path of the Dockerfile to Document")

	return dockerCmd
}

func (d *tagsCmd) run() error {
	if len(d.dockerfile) != 0 {
		return printTags(d.dockerfile)
	}
	return errors.New(utils.ColorString("red", "Please specfiy a path for the dockerfile to Dockument"))

}

func printTags(dockerfile string) error {
	tags, err := labels.GetTags(dockerfile)
	fmt.Println(utils.ColorString("blue", "### Tags / Metadata ###"))
	for tag, value := range *tags {
		fmt.Printf("	%s: %s \n", utils.ColorString("green", tag), value)
	}
	return err
}
