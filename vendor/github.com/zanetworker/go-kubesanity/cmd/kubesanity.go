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

	kubesanity_env "github.com/zanetworker/go-kubesanity/pkg/environment"
	"github.com/zanetworker/go-kubesanity/pkg/log"
)

var (
	settings      kubesanity_env.EnvSettings
	documentation string
	// cfgFile  string
)

const globalUsage = `
Kubesanity is a kubernetes validation tool, it can be used to validate the configuration of various kubernetes objects, for example pods, services,
volumes, etc. 

Environment:
$KUBESANITY_HOME          set an alternative Kubesanity location for files. By default, these are stored in ~/.kubesanity
`

var kubesanityLogo = ` 
##    ## ##     ## ########  ########  ######     ###    ##    ## #### ######## ##    ## 
##   ##  ##     ## ##     ## ##       ##    ##   ## ##   ###   ##  ##     ##     ##  ##  
##  ##   ##     ## ##     ## ##       ##        ##   ##  ####  ##  ##     ##      ####   
#####    ##     ## ########  ######    ######  ##     ## ## ## ##  ##     ##       ##    
##  ##   ##     ## ##     ## ##             ## ######### ##  ####  ##     ##       ##    
##   ##  ##     ## ##     ## ##       # #    ## ##     ## ##   ###  ##     ##       ##    
##    ##  #######  ########  ########  ######  ##     ## ##    ## ####    ##       ##
                                                                      
`

func newRootCmd(args []string) *cobra.Command {
	// rootCmd represents the base command when called without any subcommands
	kubesanityCmd := &cobra.Command{
		Use:   "kubesanity",
		Short: "kubesanity is a kubernetes validation tool",
		Long:  globalUsage,
		Run:   runKubesanity,
	}

	flags := kubesanityCmd.PersistentFlags()
	settings.AddFlags(flags)
	out := kubesanityCmd.OutOrStdout()

	// out := kubesanityCmd.OutOrStdout()
	kubesanityCmd.AddCommand(
		newNetworkCmd(out),
		newVersionCmd(out),
	)

	// set defaults from environment
	settings.Init(flags)

	return kubesanityCmd
}

func printLogo() {
	figletColoured := aec.BlueF.Apply(kubesanityLogo)
	if runtime.GOOS == "windows" {
		figletColoured = aec.BlueF.Apply(kubesanityLogo)
	}
	if _, err := fmt.Println(figletColoured); err != nil {
		log.ErrorS("Failed to print kubesanity figlet", err)
	}
}

func runKubesanity(cmd *cobra.Command, args []string) {
	printLogo()
	if len(args) == 0 {
		cmd.Help()
	}
}

//Execute command for kubesanity CLI
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
