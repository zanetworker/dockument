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
	var commandTests *labels.CommandTests
	var fexTests *labels.FileExistenceTests

	var err error

	switch inputType {
	case FILE:
		commandTests, err = labels.GetCommandTests(name)
		fexTests, err = labels.GetFileExistenceTests(name)
	case IMAGE:
		commandTests, err = labels.GetCommandTests(name)
		fexTests, err = labels.GetFileExistenceTests(name)
	}

	printCommandTests(commandTests)
	printFileExistenceTests(fexTests)
	return err
}

func printCommandTests(cmdTests *labels.CommandTests) {
	nilCommandTests := &labels.CommandTests{}
	if !reflect.DeepEqual(nilCommandTests, cmdTests) {
		for _, test := range *(cmdTests) {
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
}

func printFileExistenceTests(fexTests *labels.FileExistenceTests) {
	nilFexTests := &labels.FileExistenceTests{}
	if !reflect.DeepEqual(nilFexTests, fexTests) {
		for _, test := range *(fexTests) {
			fmt.Printf(utils.ColorString("blue", "### File Existence Tests %s ### \n"), strings.ToUpper(test.Name))
			fmt.Printf("	%s: %s \n", utils.ColorString("green", "Path"), test.Path)
			fmt.Printf("	%s: %s\n", utils.ColorString("green", "Should Exist"), test.ShouldExist)
			fmt.Printf("	%s: %s\n", utils.ColorString("green", "Permissions"), test.Permissions)
		}
	}
}
