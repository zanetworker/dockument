package commands

import (
	"os"
	"text/template"

	log "github.com/sirupsen/logrus"
	"github.com/zanetworker/dockument/pkg/labels"
	"github.com/zanetworker/dockument/pkg/utils"
)

// Functions to use in the template
var fm = template.FuncMap{}
var tplDockument *template.Template

type dockumentTemplateParams struct {
	Dependencies *labels.Dependencies
	Envs         *labels.Envs
	Ports        *labels.Ports
}

func init() {
	dockerfileDockument := utils.GetDir("root") + "/configs/" + "dockumentation.tpl"
	validTemplatePath := isValidPath(dockerfileDockument)
	if !validTemplatePath {
		log.Fatalf("Could not create Dockument, dockertemplate path  ( %s ) is invalid", dockerfileDockument)
	}
	tplDockument = template.Must(template.New("").Funcs(fm).ParseFiles(dockerfileDockument))
}

//CreateDockument creates the dockerfile dock-umentation
func CreateDockument(dockerfile, dockumentPath string) {
	validOutPath := isValidPath(dockumentPath)
	if !validOutPath {
		//TODO replace with proper error type
		log.Fatalf("Could not create Dockument, output path  ( %s ) is invalid, please check command help", dockumentPath)
	}

	validDockerfilePath := isValidPath(dockerfile)
	if !validDockerfilePath {
		log.Fatalf("Could not create Dockument, dockerfile path  ( %s ) is invalid, please check command help", dockerfile)
	}

	f, err := os.Create(dockumentPath)
	if err != nil {
		log.Errorf("failed to create a dockument, error: %s ", err.Error())
	}

	// Fetch dependencies
	deps, err := labels.GetDepenedencies(dockerfile)
	if err != nil {
		log.Fatalf("Failed to retrieve dependencies, error: %s", err.Error())
	}

	// Fetch Environment variables
	envs, err := labels.GetEnvs(dockerfile)
	if err != nil {
		log.Fatalf("Failed to retrieve envs, error: %s", err.Error())
	}

	// Fetch Environment variables
	ports, err := labels.GetExposedPorts(dockerfile)
	if err != nil {
		log.Fatalf("Failed to retrieve envs, error: %s", err.Error())
	}

	templateData := dockumentTemplateParams{
		Dependencies: deps,
		Envs:         envs,
		Ports:        ports,
	}

	tplDockument.ExecuteTemplate(f, "dockument", templateData)
}
