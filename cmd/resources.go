package main

import (
	"errors"
	"fmt"
	"io"

	"github.com/spf13/cobra"
	"github.com/zanetworker/dockument/pkg/labels"
	"github.com/zanetworker/dockument/pkg/utils"
)

var resourcesCmdDesc = `
This command is used to fetch resources out of Dockerfiles Dockerfiles`

type resourcesCmd struct {
	dockerfile string
}

func newResourcesCmd(out io.Writer) *cobra.Command {
	resourcesCmdParams := &resourcesCmd{}
	dockerCmd := &cobra.Command{
		Use:   "resources",
		Short: "fetches resources from the Dockerfile",
		Long:  resourcesCmdDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			return resourcesCmdParams.run()
		},
	}

	f := dockerCmd.Flags()
	f.StringVar(&resourcesCmdParams.dockerfile, "dockerfile", "", "the path of the Dockerfile to Document")

	return dockerCmd
}

func (d *resourcesCmd) run() error {
	if len(d.dockerfile) != 0 {
		return printResources(d.dockerfile)
	}
	return errors.New(utils.ColorString("red", "Please specfiy a path for the dockerfile to Dockument"))

}

func printResources(dockerfile string) error {
	resources, err := labels.GetResources(dockerfile)
	fmt.Println(utils.ColorString("blue", "### Resources ###"))
	fmt.Printf("	%s: %s \n", utils.ColorString("green", "CPU"), resources.CPU)
	fmt.Printf("	%s: %s\n", utils.ColorString("green", "Memoory"), resources.Memory)
	return err
}
