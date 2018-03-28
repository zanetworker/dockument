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

DOCKument Created! 
`

var dockerCreateCmdDesc = `
This command is used to DOCKument Dockerfiles & Docker Images`

var defaultDockumentLocation = utils.GetDir("root")

type dockerCreateCmd struct {
	dockerfile, image string
	outLocation       string
}

func newDockerCreateCmd(out io.Writer) *cobra.Command {
	dockerCreateCmdParams := &dockerCreateCmd{}
	dockerCmd := &cobra.Command{
		Use:   "create",
		Short: "create documentation for Dockerfiles or Docker images",
		Long:  dockerCreateCmdDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			return dockerCreateCmdParams.run()
		},
	}

	f := dockerCmd.Flags()

	f.StringVarP(&dockerCreateCmdParams.dockerfile, "dockerfile", "d", "", "the path of the Dockerfile to Dockument")
	f.StringVarP(&dockerCreateCmdParams.image, "image", "i", "", "the path of the Dockerfile to Dockument")
	f.StringVarP(&dockerCreateCmdParams.outLocation, "out", "o", defaultDockumentLocation, "the output location of documentation")

	return dockerCmd
}

func (d *dockerCreateCmd) run() error {
	if len(d.dockerfile) != 0 {
		commands.CreateDockument(d.dockerfile, d.outLocation)
		defer log.Infof("Dockument created for Dockerfile (%s) at ( %s )", d.dockerfile, d.outLocation)

	}

	if len(d.image) != 0 {
		commands.CreateImageDockument(d.image, d.outLocation)
		defer log.Infof("Dockument created for Docker Image (%s) at ( %s )", d.image, d.outLocation)

	}
	printLogo(thumbsUpLogo)
	return nil
}
