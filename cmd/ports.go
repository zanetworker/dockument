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
This command is used to fetch exposed ports out of Dockerfiles Dockerfiles`

type portsCmd struct {
	dockerfile string
}

func newPortsCmd(out io.Writer) *cobra.Command {
	portsCmdParams := &portsCmd{}
	dockerCmd := &cobra.Command{
		Use:   "ports",
		Short: "fetches exposed ports from the Dockerfile",
		Long:  portsCmdDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			return portsCmdParams.run()
		},
	}

	f := dockerCmd.Flags()
	f.StringVar(&portsCmdParams.dockerfile, "dockerfile", "", "the path of the Dockerfile to Dockument")

	return dockerCmd
}
func (d *portsCmd) run() error {
	if len(d.dockerfile) != 0 {
		return printPorts(d.dockerfile)
	}
	return errors.New(utils.ColorString("red", "Please specfiy a path for the dockerfile to Dockument"))

}

func printPorts(dockerfile string) error {
	ports, err := labels.GetExposedPorts(dockerfile)
	for _, port := range *(ports) {
		fmt.Printf(utils.ColorString("blue", "### Port %s ### \n"), strings.ToUpper(port.Name))
		fmt.Printf("	%s: %s \n", utils.ColorString("green", "Name"), port.Name)
		fmt.Printf("	%s: %s\n", utils.ColorString("green", "Description"), port.About)
		fmt.Printf("	%s: %s\n", utils.ColorString("green", "Scheme"), port.Scheme)
		fmt.Printf("	%s: %s\n", utils.ColorString("green", "Protocol"), port.Protocol)
	}
	return err
}
