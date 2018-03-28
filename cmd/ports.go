package main

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zanetworker/dockument/pkg/labels"
	"github.com/zanetworker/dockument/pkg/utils"
)

var portsCmdDesc = `
This command is used to fetch exposed ports out of Dockerfiles or images`

type portsCmd struct {
	dockerfile string
	imageName  string
}

func newPortsCmd(out io.Writer) *cobra.Command {
	portsCmdParams := &portsCmd{}
	dockerCmd := &cobra.Command{
		Use:   "ports",
		Short: "fetches exposed ports from the Dockerfile or Docker images",
		Long:  portsCmdDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			return portsCmdParams.run()
		},
	}

	f := dockerCmd.Flags()
	f.StringVarP(&portsCmdParams.dockerfile, "dockerfile", "d", "", "the path of the Dockerfile to Dockument")
	f.StringVarP(&portsCmdParams.imageName, "image", "i", "", "the name of the image to fetch labels from")

	return dockerCmd
}
func (d *portsCmd) run() error {
	if len(d.dockerfile) != 0 {
		return printPorts(d.dockerfile, FILE)
	}
	if len(d.imageName) != 0 {
		return printPorts(d.imageName, IMAGE)
	}
	return errors.New(utils.ColorString("red", "Please specfiy a path for the dockerfile to Dockument"))

}

func printPorts(name, inputType string) error {
	var ports *labels.Ports
	var err error

	switch inputType {
	case FILE:
		ports, err = labels.GetExposedPorts(name)
	case IMAGE:
		ports, err = labels.GetImageExposedPorts(name)
	}

	for _, port := range *(ports) {
		fmt.Printf(utils.ColorString("blue", "### Port %s ### \n"), strings.ToUpper(port.Name))
		fmt.Printf("	%s: %s \n", utils.ColorString("green", "Name"), port.Name)
		fmt.Printf("	%s: %s\n", utils.ColorString("green", "Description"), port.About)
		fmt.Printf("	%s: %s\n", utils.ColorString("green", "Scheme"), port.Scheme)
		fmt.Printf("	%s: %s\n", utils.ColorString("green", "Protocol"), port.Protocol)
	}
	return err
}
