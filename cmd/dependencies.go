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
		Long:  dependenciesCmdDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			return dependenciesCmdParams.run()
		},
	}

	f := dockerCmd.Flags()
	f.StringVar(&dependenciesCmdParams.dockerfile, "dockerfile", "", "the path of the Dockerfile to Dockument")

	return dockerCmd
}
func (d *dependenciesCmd) run() error {
	if len(d.dockerfile) != 0 {
		return printDependencies(d.dockerfile)
	}

	return errors.New(utils.ColorString("red", "Please specfiy a path for the dockerfile to Dockument"))

}

func printDependencies(dockerfile string) error {
	dependencies, err := labels.GetDepenedencies(dockerfile)
	for _, dependency := range *(dependencies) {
		fmt.Printf(utils.ColorString("blue", "### Dependency %s ### \n"), strings.ToUpper(dependency.Name))
		fmt.Printf("	%s: %s \n", utils.ColorString("green", "Application"), dependency.Name)
		fmt.Printf("	%s: %s\n", utils.ColorString("green", "Image"), dependency.Image)
		fmt.Printf("	%s: %s\n", utils.ColorString("green", "Description"), dependency.About)
		fmt.Printf("	%s: %s\n", utils.ColorString("green", "Ports"), dependency.Ports)
		fmt.Printf("	%s: %s\n", utils.ColorString("green", "Required"), dependency.Mandatory)
	}
	return err
}
