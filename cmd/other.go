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

var otherCmdDesc = `
This command is used to fetch all other non-DOCKument labels out of Dockerfiles or images`

type otherCmd struct {
	dockerfile string
	imageName  string
}

func newOtherCmd(out io.Writer) *cobra.Command {
	otherCmdParams := &otherCmd{}
	dockerCmd := &cobra.Command{
		Use:   "other",
		Short: "fetch all other non-DOCKument labels out of Dockerfiles or Docker images",
		Long:  otherCmdDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			return otherCmdParams.run()
		},
	}

	f := dockerCmd.Flags()
	f.StringVarP(&otherCmdParams.dockerfile, "dockerfile", "d", "", "the path of the Dockerfile to Dockument")
	f.StringVarP(&otherCmdParams.imageName, "image", "i", "", "the name of the image to fetch labels from")

	return dockerCmd
}
func (d *otherCmd) run() error {
	if len(d.dockerfile) != 0 {
		err := printOthers(d.dockerfile, FILE)
		return err
	}
	if len(d.imageName) != 0 {
		err := printOthers(d.imageName, IMAGE)
		return err
	}
	return errors.New(utils.ColorString("red", "Please specfiy a path for the dockerfile to Dockument"))
}

func printOthers(name, inputType string) error {
	var others *labels.Others
	var err error

	switch inputType {
	case FILE:
		others, err = labels.GetOtherTags(name)
	case IMAGE:
		others, err = labels.GetOtherImageTags(name)
	}

	nilOthers := &labels.Others{}
	if !reflect.DeepEqual(nilOthers, others) {
		fmt.Println(utils.ColorString("blue", "### Misc Labels ###"))
		for other, value := range *others {
			fmt.Printf("	%s: %s \n", utils.ColorString("green", other), value)
		}
	}
	return err
}
