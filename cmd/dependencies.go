package main

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
	"github.com/zanetworker/dockument/pkg/labels"
	"github.com/zanetworker/dockument/pkg/utils"
)

var dependenciesCmdDesc = `
This command is used to fetch dependencies out of Dockerfiles Dockerfiles`

type dependenciesCmd struct {
	dockerfile string
}

func newDependenciesCmd(out io.Writer) *cobra.Command {
	dependenciesCmdParams := &dependenciesCmd{}
	dockerCmd := &cobra.Command{
		Use:   "deps",
		Short: "fetches dependencies from the Dockerfile",
		Long:  dockerCreateCmdDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			return dependenciesCmdParams.run()
		},
	}

	f := dockerCmd.Flags()
	f.StringVar(&dependenciesCmdParams.dockerfile, "dockerfile", "", "the path of the Dockerfile to Document")

	return dockerCmd
}
func (d *dependenciesCmd) run() error {
	dependencies, err := labels.GetDepenedencies(d.dockerfile)
	for _, dependency := range *(dependencies) {
		fmt.Println(utils.ColorString("blue", "### Dependency ###"))
		fmt.Printf("	%s: %s \n", utils.ColorString("green", "Application"), dependency.Name)
		fmt.Printf("	%s: %s\n", utils.ColorString("green", "Image"), dependency.Image)
		fmt.Printf("	%s: %s\n", utils.ColorString("green", "Description"), dependency.About)
		fmt.Printf("	%s: %s\n", utils.ColorString("green", "Ports"), dependency.Ports)
		fmt.Printf("	%s: %s\n", utils.ColorString("green", "Required"), dependency.Mandatory)
	}
	return err
}
