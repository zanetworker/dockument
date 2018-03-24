package main

import (
	"errors"
	"io"

	"github.com/spf13/cobra"
	"github.com/zanetworker/dockument/pkg/utils"
)

var allCmdDesc = `
This command is used to fetch all the important DOCKument labels out of Dockerfiles Dockerfiles`

const (
	FILE  = "file"
	IMAGE = "image"
)

type allCmd struct {
	dockerfile string
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
	f.StringVar(&allCmdParams.dockerfile, "dockerfile", "", "the path of the Dockerfile to Dockument")

	return dockerCmd
}
func (d *allCmd) run() error {
	if len(d.dockerfile) != 0 {
		// TODO: add proper error handling here
		err := printDependencies(d.dockerfile, FILE)
		err = printPorts(d.dockerfile, FILE)
		err = printEnvs(d.dockerfile, FILE)
		err = printResources(d.dockerfile)
		err = printTags(d.dockerfile)
		err = printOthers(d.dockerfile)
		return err
	}
	return errors.New(utils.ColorString("red", "Please specfiy a path for the dockerfile to Dockument"))
}

func returnIfErr(err error) {
	if err != nil {

	}
}
