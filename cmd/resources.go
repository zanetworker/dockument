package main

import (
	"errors"
	"fmt"
	"io"
	"reflect"

	"github.com/spf13/cobra"
	"github.com/zanetworker/dockument/pkg/labels"
	"github.com/zanetworker/dockument/pkg/utils"
)

var resourcesCmdDesc = `
This command is used to fetch resources out of Dockerfiles or images`

type resourcesCmd struct {
	dockerfile string
	imageName  string
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
	f.StringVarP(&resourcesCmdParams.dockerfile, "dockerfile", "d", "", "the path of the Dockerfile to Document")
	f.StringVarP(&resourcesCmdParams.imageName, "image", "i", "", "the name of the image to fetch labels from")

	return dockerCmd
}

func (d *resourcesCmd) run() error {
	if len(d.dockerfile) != 0 {
		return printResources(d.dockerfile, "file")
	}
	if len(d.imageName) != 0 {
		return printResources(d.imageName, "image")
	}
	return errors.New(utils.ColorString("red", "Please specfiy a path for the dockerfile to Dockument"))
}

func printResources(name, inputType string) error {
	var resources *labels.Resources
	var err error

	switch inputType {
	case FILE:
		resources, err = labels.GetResources(name)
	case IMAGE:
		resources, err = labels.GetImageResources(name)
	}
	nilResources := &labels.Resources{}
	if !reflect.DeepEqual(nilResources, resources) {
		fmt.Println(utils.ColorString("blue", "### Resources ###"))
		fmt.Printf("	%s: %s \n", utils.ColorString("green", "CPU"), resources.CPU)
		fmt.Printf("	%s: %s\n", utils.ColorString("green", "Memoory"), resources.Memory)
	}
	return err
}
