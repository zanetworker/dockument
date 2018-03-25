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
	EXPOSED = "EXPOSED"

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

	// Make sure that the dependencies conform with the schema
	json, _ := json.Marshal(dockerfileDependencies)
	valid, err := validateMyObjectWithSchema("dependencies.json", string(json), "raw")
	if err != nil {
		log.Error(err.Error())
	}

	//TODO improve or decouple validation
	if !valid {
		log.Warn("your dependencies are not in a valid json format thus can't be used as is")
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

	labelsToReturn, err := fetchLabelsFor(ENVS, labels)
	if err != nil {
		return nil, err
	}

	dockerfileEnvs := parseEnvs(labelsToReturn)
	// Make sure that the envs conform with the schema
	json, _ := json.Marshal(dockerfileEnvs)
	valid, err := validateMyObjectWithSchema("envs.json", string(json), "raw")
	if err != nil {
		log.Error(err.Error())
	}

	//TODO improve or decouple validation
	if !valid {
		log.Warn("your ENVs are not in a valid json format thus can't be used as is")
	}

	return dockerfileEnvs, nil
}

//GetResources fetch the container resources from the Dockerfile
func GetResources(dockerfile string) (*Resources, error) {
	//TODO
	labels, err := getLabels(dockerfile)
	if err != nil {
		return nil, err
	}

	resourcesToReturn, err := fetchLabelsFor(RESOURCES, labels)
	if err != nil {
		return nil, err
	}

	dockerfileResources := parseResources(resourcesToReturn)
	json, _ := json.Marshal(dockerfileResources)
	valid, err := validateMyObjectWithSchema("resources.json", string(json), "raw")
	if err != nil {
		log.Error(err.Error())
	}

	//TODO improve or decouple validation
	if !valid {
		log.Warn("your Resources are not in a valid json format thus can't be used as is")
	}
	return dockerfileResources, nil
}

//GetExposedPorts fetch the texposed ports from the Dockerfile
func GetExposedPorts(dockerfile string) (*Ports, error) {
	//TODO
	labels, err := getLabels(dockerfile)
	if err != nil {
		return nil, err
	}

	labelsToReturn, err := fetchLabelsFor(EXPOSED, labels)
	if err != nil {
		return nil, err
	}
	dockerfilePorts := parsePorts(labelsToReturn)

	// Make sure that the envs conform with the schema
	json, _ := json.Marshal(dockerfilePorts)
	valid, err := validateMyObjectWithSchema("ports.json", string(json), "raw")
	if err != nil {
		log.Error(err.Error())
	}

	//TODO improve or decouple validation
	if !valid {
		log.Warn("your Ports are not in a valid json format thus can't be used as is")
	}

	return dockerfilePorts, nil

}

//GetTags fetch the tags used in the Dockerfile
func GetTags(dockerfile string) (*Tags, error) {
	labels, err := getLabels(dockerfile)
	if err != nil {
		return nil, err
	}

	tagsToReturn, err := fetchLabelsFor(TAGS, labels)
	if err != nil {
		return nil, err
	}

	dockerfileTags := parseTags(tagsToReturn)
	return dockerfileTags, nil
}

//GetOtherTags gets all the dockerfile labels
func GetOtherTags(dockerfile string) (*Others, error) {
	labels, err := getLabels(dockerfile)
	if err != nil {
		return nil, err
	}
	otherLabels, err := fetchOthers(labels)
	return &otherLabels, err
}
