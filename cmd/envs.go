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

var envsCmdDesc = `
This command is used to fetch important ENVs out of Dockerfiles or images`

type envsCmd struct {
	dockerfile string
	imageName  string
}

func newEnvsCmd(out io.Writer) *cobra.Command {
	envsCmdParams := &envsCmd{}
	dockerCmd := &cobra.Command{
		Use:   "envs",
		Short: "fetches exposed envs from the Dockerfile",
		Long:  envsCmdDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			return envsCmdParams.run()
		},
	}

	f := dockerCmd.Flags()
	f.StringVarP(&envsCmdParams.dockerfile, "dockerfile", "d", "", "the path of the Dockerfile to Dockument")
	f.StringVarP(&envsCmdParams.imageName, "image", "i", "", "the name of the image to fetch labels from")

	return dockerCmd
}
func (d *envsCmd) run() error {
	if len(d.dockerfile) != 0 {
		return printEnvs(d.dockerfile, FILE)
	}
	if len(d.imageName) != 0 {
		return printEnvs(d.imageName, IMAGE)

	}
	return errors.New(utils.ColorString("red", "Please specfiy a path for the dockerfile to Dockument"))

}

func printEnvs(name, inputType string) error {
	var envs *labels.Envs
	var err error

	switch inputType {
	case FILE:
		envs, err = labels.GetEnvs(name)
	case IMAGE:
		envs, err = labels.GetImageEnvs(name)

	}
	for _, env := range *envs {
		fmt.Printf(utils.ColorString("blue", "### ENV %s ### \n"), strings.ToUpper(env.Name))
		fmt.Printf("	%s: %s \n", utils.ColorString("green", "Name"), env.Name)
		fmt.Printf("	%s: %s\n", utils.ColorString("green", "Description"), env.About)
		fmt.Printf("	%s: %s\n", utils.ColorString("green", "Mandatory"), env.Mandatory)
	}
	return err
}
