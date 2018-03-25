package main

import (
	"errors"
	"fmt"
	"io"
	"reflect"

	"github.com/spf13/cobra"
	"github.com/zanetworker/dockument/pkg/labels"
	"github.com/zanetworker/dockument/pkg/utils"
)

var tagsCmdDesc = `
This command is used to fetch tags out of Dockerfiles or images`

type tagsCmd struct {
	dockerfile string
	imageName  string
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
	f.StringVarP(&tagCmdParams.dockerfile, "dockerfile", "d", "", "the path of the Dockerfile to Document")
	f.StringVarP(&tagCmdParams.imageName, "image", "i", "", "the name of the image to fetch labels from")

	return dockerCmd
}

func (d *tagsCmd) run() error {
	if len(d.dockerfile) != 0 {
		return printTags(d.dockerfile, FILE)
	}
	if len(d.imageName) != 0 {
		return printTags(d.imageName, IMAGE)
	}
	return errors.New(utils.ColorString("red", "Please specfiy a path for the dockerfile to Dockument"))

}

func printTags(name, inputType string) error {
	var tags *labels.Tags
	var err error

	switch inputType {
	case FILE:
		tags, err = labels.GetTags(name)
	case IMAGE:
		tags, err = labels.GetImageTags(name)
	}

	nilTags := &labels.Tags{}
	if !reflect.DeepEqual(nilTags, tags) {
		fmt.Println(utils.ColorString("blue", "### Tags / Metadata ###"))
		for tag, value := range *tags {
			fmt.Printf("	%s: %s \n", utils.ColorString("green", tag), value)

		}
	}
	return err
}
