package main

import (
	"io"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/zanetworker/dockument/pkg/log"
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
			log.Info("number of flags " + strconv.Itoa(len(args)))

			return dockerCmdParams.run()
		},
	}

	f := dockerCmd.Flags()
	f.StringVar(&dockerCmdParams.dockerfile, "dockerfile", "", "the path of the Dockerfile to Document")
	f.StringVarP(&dockerCmdParams.outLocation, "out", "o", defaultDockumentLocation, "the output location of documentation")

	return dockerCmd
}

func (n *dockerCmdParams) run() error {

	return nil
}
