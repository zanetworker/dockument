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
This command is used to fetch dependencies out of Dockerfiles or images`

type dependenciesCmd struct {
	dockerfile string
	imageName  string
}

func newDependenciesCmd(out io.Writer) *cobra.Command {
	dependenciesCmdParams := &dependenciesCmd{}
	dockerCmd := &cobra.Command{
		Use:   "deps",
		Short: "fetches dependencies from the Dockerfile or Docker image",
		Long:  dependenciesCmdDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			return dependenciesCmdParams.run()
		},
	}

	f := dockerCmd.Flags()
	f.StringVarP(&dependenciesCmdParams.dockerfile, "dockerfile", "d", "", "the path of the Dockerfile to Dockument")
	f.StringVarP(&dependenciesCmdParams.imageName, "image", "i", "", "the name of the image to fetch labels from")

	return dockerCmd
}
func (d *dependenciesCmd) run() error {
	if len(d.dockerfile) != 0 {
		return printDependencies(d.dockerfile, FILE)
	}

	if len(d.imageName) != 0 {
		return printDependencies(d.imageName, IMAGE)
	}
	return errors.New(utils.ColorString("red", "Please specfiy a path for the dockerfile to Dockument"))

}

func printDependencies(name, inputType string) error {
	var dependencies *labels.Dependencies
	var err error

	switch inputType {
	case FILE:
		dependencies, err = labels.GetDepenedencies(name)
	case IMAGE:
		dependencies, err = labels.GetImageDepenedencies(name)
	}

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
