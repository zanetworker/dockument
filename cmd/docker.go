package main

import (
	"io"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/zanetworker/dockument/internal/commands"
	"github.com/zanetworker/dockument/pkg/utils"
)

var thumbsUpLogo = `
‎________$$$$
_______$$__$
_______$___$$
_______$___$$
_______$$___$$
________$____$$
________$$____$$$
_________$$_____$$
_________$$______$$
__________$_______$$
____$$$$$$$________$$
__$$$_______________$$$$$$
_$$____$$$$____________$$$
_$___$$$__$$$____________$$
_$$________$$$____________$
__$$____$$$$$$____________$
__$$$$$$$____$$___________$
__$$_______$$$$___________$
___$$$$$$$$$__$$_________$$
____$________$$$$_____$$$$
____$$____$$$$$$____$$$$$$
_____$$$$$$____$$__$$
_______$_____$$$_$$$
________$$$$$$$$$$
`

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
	commands.CreateDockument(d.dockerfile, d.outLocation)
	log.Infof("Dockument created at ( %s )", d.outLocation)
	printLogo(thumbsUpLogo)
	return nil
}
