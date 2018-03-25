package main

import (
	"errors"
	"io"

	"github.com/spf13/cobra"
	"github.com/zanetworker/dockument/pkg/utils"
)

var allCmdDesc = `
This command is used to fetch all the important DOCKument labels out of Dockerfiles Dockerfiles`

type allCmd struct {
	dockerfile string
	imageName  string
}

func newAllCmd(out io.Writer) *cobra.Command {
	allCmdParams := &allCmd{}
	dockerCmd := &cobra.Command{
		Use:   "all",
		Short: "fetches all the important DOCKument labels from the Dockerfile",
		Long:  allCmdDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			return allCmdParams.run()
		},
	}

	f := dockerCmd.Flags()
	f.StringVarP(&allCmdParams.dockerfile, "dockerfile", "d", "", "the path of the Dockerfile to Dockument")
	f.StringVarP(&allCmdParams.imageName, "image", "i", "", "the name of the image to fetch labels from")

	return dockerCmd
}
func (d *allCmd) run() error {
	if len(d.dockerfile) != 0 {
		// TODO: add proper error handling here
		printDependencies(d.dockerfile, FILE)
		printPorts(d.dockerfile, FILE)
		printEnvs(d.dockerfile, FILE)
		printResources(d.dockerfile, FILE)
		printTags(d.dockerfile, FILE)
		printOthers(d.dockerfile, FILE)
		return nil
	}

	if len(d.imageName) != 0 {
		printDependencies(d.imageName, IMAGE)
		printPorts(d.imageName, IMAGE)
		printEnvs(d.imageName, IMAGE)
		printResources(d.imageName, IMAGE)
		printTags(d.imageName, IMAGE)
		printOthers(d.imageName, IMAGE)
		return nil
	}
	return errors.New(utils.ColorString("red", "Please specfiy a path for the dockerfile to Dockument"))
}

func returnIfErr(err error) {
	if err != nil {

	}
}
