package main

import (
	"errors"
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zanetworker/dockument/pkg/labels"
	"github.com/zanetworker/dockument/pkg/utils"
)

var testsCmdDesc = `
This command is used to fetch tests out of Dockerfiles or images`

type testsCmd struct {
	dockerfile string
	imageName  string
}

func newTestsCmd(out io.Writer) *cobra.Command {
	testsCmdParams := &testsCmd{}
	dockerCmd := &cobra.Command{
		Use:   "tests",
		Short: "fetches tests from the Dockerfile",
		Long:  testsCmdDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			return testsCmdParams.run()
		},
	}

	f := dockerCmd.Flags()
	f.StringVarP(&testsCmdParams.dockerfile, "dockerfile", "d", "", "the path of the Dockerfile to Document")
	f.StringVarP(&testsCmdParams.imageName, "image", "i", "", "the name of the image to fetch labels from")

	return dockerCmd
}

func (d *testsCmd) run() error {
	if len(d.dockerfile) != 0 {
		return printTests(d.dockerfile, FILE)
	}
	if len(d.imageName) != 0 {
		return printTests(d.imageName, IMAGE)
	}
	return errors.New(utils.ColorString("red", "Please specfiy a path for the dockerfile to Dockument"))

}

func printTests(name, inputType string) error {
	switch inputType {
	case FILE:
		printCommandTests(name, FILE)
		printFileExistenceTests(name, FILE)
		printMetadataTests(name, FILE)
		printFileContentTests(name, FILE)
	case IMAGE:
		printCommandTests(name, IMAGE)
		printFileExistenceTests(name, IMAGE)
		printMetadataTests(name, IMAGE)
		printFileContentTests(name, IMAGE)
	}

	return nil
}

func printCommandTests(name, inputType string) error {
	var commandTests *labels.CommandTests
	var err error

	switch inputType {
	case FILE:
		commandTests, err = labels.GetCommandTests(name)
	case IMAGE:
		commandTests, err = labels.GetImageCommandTests(name)
	}

	nilCommandTests := &labels.CommandTests{}
	if !reflect.DeepEqual(nilCommandTests, commandTests) {
		for _, test := range *(commandTests) {
			fmt.Printf(utils.ColorString("blue", "### Command Test %s ### \n"), strings.ToUpper(test.Name))
			fmt.Printf("	%s: %s \n", utils.ColorString("green", "Name"), test.Name)
			fmt.Printf("	%s: %s\n", utils.ColorString("green", "Command"), test.Command)
			fmt.Printf("	%s: %s\n", utils.ColorString("green", "Args"), test.Args)
			fmt.Printf("	%s: %s\n", utils.ColorString("green", "Expected Ouput"), test.ExpectedOutput)
			fmt.Printf("	%s: %s\n", utils.ColorString("green", "Expected Error"), test.ExpectedError)
			fmt.Printf("	%s: %s\n", utils.ColorString("green", "Excluded Ouput"), test.ExcludedOutput)
			fmt.Printf("	%s: %s\n", utils.ColorString("green", "Excluded Error"), test.ExcludedError)
		}
	}

	return err
}

func printFileExistenceTests(name, inputType string) error {
	var fexTests *labels.FileExistenceTests
	var err error

	switch inputType {
	case FILE:
		fexTests, err = labels.GetFileExistenceTests(name)
	case IMAGE:
		fexTests, err = labels.GetImageFileExistenceTests(name)
	}

	nilFexTests := &labels.FileExistenceTests{}
	if !reflect.DeepEqual(nilFexTests, fexTests) {
		for _, test := range *(fexTests) {
			fmt.Printf(utils.ColorString("blue", "### File Existence Tests %s ### \n"), strings.ToUpper(test.Name))
			fmt.Printf("	%s: %s \n", utils.ColorString("green", "Path"), test.Path)
			fmt.Printf("	%s: %s\n", utils.ColorString("green", "Should Exist"), test.ShouldExist)
			fmt.Printf("	%s: %s\n", utils.ColorString("green", "Permissions"), test.Permissions)
		}
	}

	return err
}

func printMetadataTests(name, inputType string) error {
	var metadataTests *labels.MetadataTests
	var err error

	switch inputType {
	case FILE:
		metadataTests, err = labels.GetMetadataTests(name)
	case IMAGE:
		metadataTests, err = labels.GetImageMetadataTests(name)

	}

	nilMetaTests := &labels.MetadataTests{}
	if !reflect.DeepEqual(nilMetaTests, metadataTests) {
		for _, test := range *(metadataTests) {
			fmt.Printf(utils.ColorString("blue", "### Metadata Tests Tests ### \n"))
			fmt.Printf("	%s: %s \n", utils.ColorString("green", "Envs"), test.Env)
			fmt.Printf("	%s: %s\n", utils.ColorString("green", "Volumes"), test.Volumes)
			fmt.Printf("	%s: %s\n", utils.ColorString("green", "ExposedPorts"), test.ExposedPorts)
			fmt.Printf("	%s: %s\n", utils.ColorString("green", "Entry Point"), test.EntryPoint)
			fmt.Printf("	%s: %s\n", utils.ColorString("green", "CMD"), test.Cmd)
			fmt.Printf("	%s: %s\n", utils.ColorString("green", "Workdir"), test.Workdir)

		}
	}
	return err
}

func printFileContentTests(name, inputType string) error {
	var fileContentTests *labels.FileContentTests
	var err error

	switch inputType {
	case FILE:
		fileContentTests, err = labels.GetFileContentTests(name)
	case IMAGE:
		fileContentTests, err = labels.GetImageFileContentTests(name)

	}

	nilMetaTests := &labels.FileContentTests{}
	if !reflect.DeepEqual(nilMetaTests, fileContentTests) {
		for _, test := range *(fileContentTests) {
			fmt.Printf(utils.ColorString("blue", "### File Content Tests %s ### \n"), strings.ToUpper(test.Name))
			fmt.Printf("	%s: %s \n", utils.ColorString("green", "Path"), test.Path)
			fmt.Printf("	%s: %s\n", utils.ColorString("green", "Expected Contents"), test.ExpectedContents)
			fmt.Printf("	%s: %s\n", utils.ColorString("green", "Exluded Contents"), test.ExcludedContents)

		}
	}
	return err
}
