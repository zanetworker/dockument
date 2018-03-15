package main

import (
	"io"

	"github.com/spf13/cobra"
	"github.com/zanetworker/dockument/pkg/labels"
	"github.com/zanetworker/dockument/pkg/log"
	"github.com/zanetworker/dockument/pkg/utils"
)

var dockerCreateCmdDesc = `
This command is used to DOCKument Dockerfiles`

var defaultDockumentLocation = utils.GetDir("root")

type dockerCreateCmd struct {
	dockerfile  string
	outLocation string
}

func newDockerCreateCmd(out io.Writer) *cobra.Command {
	dockerCreateCmdParams := &dockerCreateCmd{}
	dockerCmd := &cobra.Command{
		Use:   "create",
		Short: "create documentation for Dockerfiles",
		Long:  dockerCreateCmdDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			return dockerCreateCmdParams.run()
		},
	}

	f := dockerCmd.Flags()

	f.StringVar(&dockerCreateCmdParams.dockerfile, "dockerfile", "", "the path of the Dockerfile to Document")
	f.StringVarP(&dockerCreateCmdParams.outLocation, "out", "o", defaultDockumentLocation, "the output location of documentation")

	return dockerCmd
}

func (d *dockerCreateCmd) run() error {
	jsonOut, err := labels.GetDepenedencies(d.dockerfile)
	log.Info(jsonOut)
	return err
}
