// Copyright © 2018 Adek Zaalouk <adel.zalok.89@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/morikuni/aec"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"

	log "github.com/sirupsen/logrus"
	dockument_env "github.com/zanetworker/dockument/pkg/environment"
)

var (
	settings      dockument_env.EnvSettings
	documentation string
	// cfgFile  string
)

const globalUsage = `
DOCKument is a tool to auto-generate Documentation for container images (Dockerfiles).

Environment:
$DOCKUMENT_HOME          set an alternative DOCKument location for files. By default, these are stored in ~/.dockument
`

const (
	//FILE is used to tell dockument to fetch labels from a Dockerfile
	FILE = "file"
	//IMAGE is used to tell dockument to fetch labels from a Docker image
	IMAGE = "image"
)

var dockumentLogo = `
######  #######  #####  #    #                                   
#     # #     # #     # #   #  #    # #    # ###### #    # ##### 
#     # #     # #       #  #   #    # ##  ## #      ##   #   #   
#     # #     # #       ###    #    # # ## # #####  # #  #   #   
#     # #     # #       #  #   #    # #    # #      #  # #   #   
#     # #     # #     # #   #  #    # #    # #      #   ##   #   
######  #######  #####  #    #  ####  #    # ###### #    #   #                                                               
`

func newRootCmd(args []string) *cobra.Command {
	// rootCmd represents the base command when called without any subcommands
	dockumentCmd := &cobra.Command{
		Use:   "dockument",
		Short: "DOCKument is a tool that facilitate the documentation of Docker images",
		Long:  globalUsage,
		Run:   runDockument,
	}

	flags := dockumentCmd.PersistentFlags()
	settings.AddFlags(flags)
	out := dockumentCmd.OutOrStdout()

	dockumentCmd.AddCommand(
		newDockerCreateCmd(out),
		newDependenciesCmd(out),
		newResourcesCmd(out),
		newPortsCmd(out),
		newTagsCmd(out),
		newEnvsCmd(out),
		newAllCmd(out),
		newOtherCmd(out),
		newTestsCmd(out),
		newVersionCmd(out),
	)

	// set defaults from environment
	settings.Init(flags)

	return dockumentCmd
}

func printLogo(logoToPrint string) {
	figletColoured := aec.GreenF.Apply(logoToPrint)
	if runtime.GOOS == "windows" {
		figletColoured = aec.BlueF.Apply(logoToPrint)
	}
	if _, err := fmt.Println(figletColoured); err != nil {
		log.Errorf("Failed to print dockument figlet, error: %s", err.Error())
	}
}

func returnWithError(err error) error {
	if err != nil {
		return err
	}
	return nil
}

func runDockument(cmd *cobra.Command, args []string) {
	printLogo(dockumentLogo)
	if len(args) == 0 {
		cmd.Help()
	}
}

//Execute command for dockument CLI
func main() {
	cmd := newRootCmd(os.Args[1:])
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}

	if len(documentation) > 0 {
		err := doc.GenMarkdownTree(cmd, "./doc")
		if err != nil {
			log.Fatal(err)
		}
	}
}
