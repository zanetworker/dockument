// Copyright © 2018 Adel Zaalouk adel.zalok.89@gmail.com
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

package environment

import (
	"os"
	"path/filepath"

	"github.com/spf13/pflag"
	"github.com/zanetworker/dockument/pkg/log"
	"github.com/zanetworker/dockument/pkg/path"
	"k8s.io/client-go/util/homedir"
)

const (
	//DockumentHomeEnvVar location of Dockument configuration file
	DockumentHomeEnvVar = "DOCKUMENT_HOME"
)

// DefaultDockumentHome is the default DOCKUMENT_HOME.
var DefaultDockumentHome = filepath.Join(homedir.HomeDir(), ".dockument")

// DefaultOutPath is the default value for
var DefaultOutPath, _ = os.Getwd()

//EnvSettings are the global environment settings
type EnvSettings struct {
	// Home is the local path to the dockument home directory.
	Home    path.Home
	Outpath string
}

// AddFlags binds flags to the given flagset.
func (s *EnvSettings) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar((*string)(&s.Home), "home", DefaultDockumentHome, "location of your dockumeny config. Overrides $DOCKUMENT_HOME ")
	fs.StringVar(&s.Outpath, "outpath", DefaultOutPath, "target location for Dockerfile documentation")
}

// EnvMap maps flag names to envvars
var EnvMap = map[string]string{
	"home": "DOCKUMENT_PATH",
}

// Init sets values from the environment.
func (s *EnvSettings) Init(fs *pflag.FlagSet) {
	for name, envar := range EnvMap {
		setFlagFromEnv(name, envar, fs)

		value, err := fs.GetString(name)
		if err != nil {
			log.Fatal(err)
		}
		_, ok := os.LookupEnv(envar)
		if !ok {
			os.Setenv(envar, value)
		}
	}
}

func setFlagFromEnv(name, envar string, fs *pflag.FlagSet) {
	//Check if this parameter was passed as a flag
	if fs.Changed(name) {
		return
	}
	if v, ok := os.LookupEnv(envar); ok {
		if err := fs.Set(name, v); err != nil {
			log.ErrorS("Failed to Set Env variable", err)
		}
	}
}
