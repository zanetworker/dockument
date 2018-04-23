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
	Dependencies       *labels.Dependencies
	Envs               *labels.Envs
	Ports              *labels.Ports
	Resources          *labels.Resources
	Tags               *labels.Tags
	CommandTests       *labels.CommandTests
	MetadataTests      *labels.MetadataTests
	FileContentTests   *labels.FileContentTests
	FileExistenceTests *labels.FileExistenceTests
	Others             *labels.Others
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
	checkAndCreatePath(dockumentPath)

	validDockerfilePath := isValidPath(dockerfile)
	if !validDockerfilePath {
		log.Fatalf("Could not create Dockument, dockerfile path  ( %s ) is invalid, please check command help", dockerfile)
	}

	f, err := os.Create(dockumentPath)
	if err != nil {
		log.Errorf("failed to create a dockument, error: %s ", err.Error())
	}
	defer f.Close()

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
	// Fetch Environment variables
	resources, err := labels.GetResources(dockerfile)
	if err != nil {
		log.Fatalf("Failed to retrieve resources, error: %s", err.Error())
	}

	// Get tag labels
	tags, err := labels.GetTags(dockerfile)
	if err != nil {
		log.Fatalf("Failed to retrieve tags, error: %s", err.Error())
	}

	// Get Misc labels
	commandTests, err := labels.GetCommandTests(dockerfile)
	if err != nil {
		log.Fatalf("Failed to retrieve command tests, error: %s", err.Error())
	}

	// Get Misc labels
	fileExistenceTests, err := labels.GetFileExistenceTests(dockerfile)
	if err != nil {
		log.Fatalf("Failed to retrieve file existence tests, error: %s", err.Error())
	}

	// Get Misc labels
	fileContentTests, err := labels.GetFileContentTests(dockerfile)
	if err != nil {
		log.Fatalf("Failed to retrieve file content tests, error: %s", err.Error())
	}

	// Get Misc labels
	metadataTests, err := labels.GetMetadataTests(dockerfile)
	if err != nil {
		log.Fatalf("Failed to retrieve file metadata tests, error: %s", err.Error())
	}

	// Get Misc labels
	others, err := labels.GetOtherTags(dockerfile)
	if err != nil {
		log.Fatalf("Failed to retrieve tags, error: %s", err.Error())
	}

	templateData := dockumentTemplateParams{
		Dependencies:       deps,
		Envs:               envs,
		Ports:              ports,
		Resources:          resources,
		Tags:               tags,
		CommandTests:       commandTests,
		FileContentTests:   fileContentTests,
		FileExistenceTests: fileExistenceTests,
		MetadataTests:      metadataTests,
		Others:             others,
	}

	tplDockument.ExecuteTemplate(f, "dockument", templateData)
}

//CreateImageDockument creates the dockerfile dock-umentation
func CreateImageDockument(imageName, dockumentPath string) {
	validImageName := labels.ImageExists(imageName)
	if !validImageName {
		log.Fatalf("Could not create Dockument, image path  ( %s ) is invalid, please check command help", imageName)
	}

	f, err := os.Create(dockumentPath)
	if err != nil {
		log.Errorf("failed to create a dockument, error: %s ", err.Error())
	}

	// Fetch dependencies
	deps, err := labels.GetImageDepenedencies(imageName)
	if err != nil {
		log.Fatalf("Failed to retrieve dependencies, error: %s", err.Error())
	}

	// Fetch Environment variables
	envs, err := labels.GetImageEnvs(imageName)
	if err != nil {
		log.Fatalf("Failed to retrieve envs, error: %s", err.Error())
	}

	// Fetch Environment variables
	ports, err := labels.GetImageExposedPorts(imageName)
	if err != nil {
		log.Fatalf("Failed to retrieve envs, error: %s", err.Error())
	}
	// Fetch Environment variables
	resources, err := labels.GetImageResources(imageName)
	if err != nil {
		log.Fatalf("Failed to retrieve resources, error: %s", err.Error())
	}

	// Get tag labels
	tags, err := labels.GetImageTags(imageName)
	if err != nil {
		log.Fatalf("Failed to retrieve tags, error: %s", err.Error())
	}

	// Get Misc labels
	commandTests, err := labels.GetImageCommandTests(imageName)
	if err != nil {
		log.Fatalf("Failed to retrieve command tests, error: %s", err.Error())
	}

	// Get Misc labels
	fileExistenceTests, err := labels.GetImageFileExistenceTests(imageName)
	if err != nil {
		log.Fatalf("Failed to retrieve file existence tests, error: %s", err.Error())
	}

	// Get Misc labels
	fileContentTests, err := labels.GetImageFileContentTests(imageName)
	if err != nil {
		log.Fatalf("Failed to retrieve file content tests, error: %s", err.Error())
	}

	// Get Misc labels
	metadataTests, err := labels.GetImageMetadataTests(imageName)
	if err != nil {
		log.Fatalf("Failed to retrieve file metadata tests, error: %s", err.Error())
	}

	// Get Misc labels
	others, err := labels.GetOtherImageTags(imageName)
	if err != nil {
		log.Fatalf("Failed to retrieve tags, error: %s", err.Error())
	}

	templateData := dockumentTemplateParams{
		Dependencies:       deps,
		Envs:               envs,
		Ports:              ports,
		Resources:          resources,
		Tags:               tags,
		CommandTests:       commandTests,
		FileContentTests:   fileContentTests,
		FileExistenceTests: fileExistenceTests,
		MetadataTests:      metadataTests,
		Others:             others,
	}

	err = tplDockument.ExecuteTemplate(f, "dockument", templateData)
	if err != nil {
		log.Fatalf("Can not execute template %s", err.Error())
	}
}

//ValidPath is a helper function to assert validity of a file path
func ValidPath(dockerfile string) bool {
	return isValidPath(dockerfile)
}
