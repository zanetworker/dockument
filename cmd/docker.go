package main

import (
	"io"

	"github.com/spf13/cobra"
	"github.com/zanetworker/DOCKument/pkg/labels"
	"github.com/zanetworker/dockument/pkg/utils"
)

var dockerCmdDesc = `
This command is used to DOCKument Dockerfiles`

var defaultDockumentLocation = utils.GetDir("root")

type dockerCmdParams struct {
	dockerfile  string
	outLocation string
}

func newDockerCmd(out io.Writer) *cobra.Command {
	dockerCmdParams := &dockerCmdParams{}

	dockerCmd := &cobra.Command{
		Use:   "create",
		Short: "create documentation for Dockerfiles",
		Long:  globalUsage,
		RunE: func(cmd *cobra.Command, args []string) error {
			return dockerCmdParams.run()
		},
	}

	f := dockerCmd.Flags()
	f.StringVar(&dockerCmdParams.dockerfile, "dockerfile", "", "the path of the Dockerfile to Document")
	f.StringVarP(&dockerCmdParams.outLocation, "out", "o", defaultDockumentLocation, "the output location of documentation")

	return dockerCmd
}

func (d *dockerCmdParams) run() error {
	err := labels.GetLabels(d.dockerfile)
	return err
}
