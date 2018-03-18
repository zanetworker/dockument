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
This command is used to fetch important ENVs out of Dockerfiles Dockerfiles`

type envsCmd struct {
	dockerfile string
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
	f.StringVar(&envsCmdParams.dockerfile, "dockerfile", "", "the path of the Dockerfile to Dockument")

	return dockerCmd
}
func (d *envsCmd) run() error {
	if len(d.dockerfile) != 0 {
		return printEnvs(d.dockerfile)
	}
	return errors.New(utils.ColorString("red", "Please specfiy a path for the dockerfile to Dockument"))

}

func printEnvs(dockerfile string) error {
	envs, err := labels.GetEnvs(dockerfile)
	for _, env := range *envs {
		fmt.Printf(utils.ColorString("blue", "### ENV %s ### \n"), strings.ToUpper(env.Name))
		fmt.Printf("	%s: %s \n", utils.ColorString("green", "Name"), env.Name)
		fmt.Printf("	%s: %s\n", utils.ColorString("green", "Description"), env.About)
		fmt.Printf("	%s: %s\n", utils.ColorString("green", "Mandatory"), env.Mandatory)
	}
	return err
}
