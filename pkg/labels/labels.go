package labels

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

// Functions to use in the template
// var fm = template.FuncMap{}

// func init() {
// 	dockerfileDockument := utils.GetDir("root") + "/templates/" + "dockumentation.md"
// }

const (
	// DEPENDENCY the dependencies tag
	DEPENDENCY = "DEPENDENCY"

	// ENVS the environment variables tag
	ENVS = "ENVS"

	// EXPOSED the exposed ports tag
	EXPOSED = "EXPOSE"

	// RESOURCES the used resources tag
	RESOURCES = "RESOURCES"

	// TAGS the extra tags used for our images
	TAGS = "TAGS"
)

//GetDepenedencies fetch the container dependencies from the Dockerfile
func GetDepenedencies(dockerfile string) (*Dependencies, error) {
	labels, err := getLabels(dockerfile)
	if err != nil {
		return nil, err
	}

	labelsToReturn, err := fetchLabelsFor(DEPENDENCY, labels)
	if err != nil {
		return nil, err
	}

	dockerfileDependencies := parseDependencies(labelsToReturn)
	json, _ := json.Marshal(dockerfileDependencies)
	valid, err := validateMyObjectWithSchema("dependencies.json", string(json), "raw")
	if err != nil {
		log.Error(err.Error())
	}

	//TODO improve or decouple validation
	if !valid {
		log.Warn("your the dependencies are not in a valid json format thus can't be used as is")
	}

	return dockerfileDependencies, nil
}

//GetEnvs fetch cotnainer environment variables from the Dockerfile
func GetEnvs(dockerfile string) (*Envs, error) {
	//TODO
	labels, err := getLabels(dockerfile)
	if err != nil {
		return nil, err
	}

	// labelsToReturn, err := fetchLabelsFor(ENVS, labels)
	// if err != nil {
	// 	return nil, err
	// }
	return nil, nil
}

//GetResources fetch the container resources from the Dockerfile
func GetResources() {

}

//GetExposedPorts fetch the texposed ports from the Dockerfile
func GetExposedPorts() {

}

//GetTags fetch the tags used in the Dockerfile
func GetTags() {

}

//GetAllLabels gets all the dockerfile labels
func GetAllLabels() {
}
