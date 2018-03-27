package labels

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

//GetImageLabels asds
func GetImageLabels(imageName string) (map[string]string, error) {
	return getImageLabels(imageName)
}

//GetImageDepenedencies fetch the container dependencies from the Dockerfile
func GetImageDepenedencies(imageName string) (*Dependencies, error) {
	labels, err := getImageLabels(imageName)
	if err != nil {
		return nil, err
	}

	labelsToReturn, err := fetchLabelsFor(DEPENDENCY, labels)
	if err != nil {
		return nil, err
	}

	imageDependencies := parseDependencies(labelsToReturn)

	// Make sure that the dependencies conform with the schema
	json, _ := json.Marshal(imageDependencies)
	valid, err := validateMyObjectWithSchema("dependencies.json", string(json), "raw")
	if err != nil {
		log.Error(err.Error())
	}

	//TODO improve or decouple validation
	if !valid {
		log.Warn("your dependencies are not in a valid json format thus can't be used as is")
	}

	return imageDependencies, nil
}

//GetImageCommandTests fetch the command tests from the container image
func GetImageCommandTests(imageName string) (*CommandTests, error) {
	labels, err := getImageLabels(imageName)
	if err != nil {
		return nil, err
	}

	labelsToReturn, err := fetchLabelsFor(COMMAND_TESTS, labels)
	if err != nil {
		return nil, err
	}

	imageCommandTests := parseCommandTests(labelsToReturn)

	// Make sure that the dependencies conform with the schema
	json, _ := json.Marshal(imageCommandTests)
	valid, err := validateMyObjectWithSchema("commandTests.json", string(json), "raw")
	if err != nil {
		log.Error(err.Error())
	}

	//TODO improve or decouple validation
	if !valid {
		log.Warn("your Command Tests are not in a valid json format thus can't be used as is")
	}

	return imageCommandTests, nil
}

//GetImageFileExistenceTests fetch the container file existence tests from the container Image
func GetImageFileExistenceTests(imageName string) (*FileExistenceTests, error) {
	labels, err := getImageLabels(imageName)
	if err != nil {
		return nil, err
	}

	labelsToReturn, err := fetchLabelsFor(FILE_EXISTENCE_TESTS, labels)
	if err != nil {
		return nil, err
	}
	imageFETests := parseFileExistenceTests(labelsToReturn)

	// Make sure that the dependencies conform with the schema
	json, _ := json.Marshal(imageFETests)
	valid, err := validateMyObjectWithSchema("fileExistenceTests.json", string(json), "raw")
	if err != nil {
		log.Error(err.Error())
	}

	//TODO improve or decouple validation
	if !valid {
		log.Warn("your File Existence Tests are not in a valid json format thus can't be used as is")
	}

	return imageFETests, nil
}

//GetImageMetadataTests fetches the image metadata tests from the container image
func GetImageMetadataTests(imageName string) (*MetadataTests, error) {
	labels, err := getImageLabels(imageName)
	if err != nil {
		return nil, err
	}
	labelsToReturn, err := fetchLabelsFor(METADATA_TESTS, labels)
	if err != nil {
		return nil, err
	}
	imageMetaTests := parseMetadataTests(labelsToReturn)

	// Make sure that the dependencies conform with the schema
	json, _ := json.Marshal(imageMetaTests)
	valid, err := validateMyObjectWithSchema("metadataTests.json", string(json), "raw")
	if err != nil {
		log.Error(err.Error())
	}

	//TODO improve or decouple validation
	if !valid {
		log.Warn("your Metadata Tests are not in a valid json format thus can't be used as is")
	}

	return imageMetaTests, nil
}

//GetImageFileContentTests fetches the me
func GetImageFileContentTests(imageName string) (*FileContentTests, error) {
	labels, err := getImageLabels(imageName)
	if err != nil {
		return nil, err
	}

	labelsToReturn, err := fetchLabelsFor(FILE_CONTENT_TESTS, labels)
	if err != nil {
		return nil, err
	}
	imageContentTests := parseFileContentTests(labelsToReturn)

	// Make sure that the dependencies conform with the schema
	json, _ := json.Marshal(imageContentTests)
	valid, err := validateMyObjectWithSchema("fileContentTests.json", string(json), "raw")
	if err != nil {
		log.Error(err.Error())
	}

	//TODO improve or decouple validation
	if !valid {
		log.Warn("your File Content Tests are not in a valid json format thus can't be used as is")
	}

	return imageContentTests, nil
}

//GetImageEnvs fetch cotnainer environment variables from the docker image
func GetImageEnvs(imageName string) (*Envs, error) {
	labels, err := getImageLabels(imageName)
	if err != nil {
		return nil, err
	}

	labelsToReturn, err := fetchLabelsFor(ENVS, labels)
	if err != nil {
		return nil, err
	}

	imageEnvs := parseEnvs(labelsToReturn)
	// Make sure that the envs conform with the schema
	json, _ := json.Marshal(imageEnvs)
	valid, err := validateMyObjectWithSchema("envs.json", string(json), "raw")
	if err != nil {
		log.Error(err.Error())
	}

	//TODO improve or decouple validation
	if !valid {
		log.Warn("your ENVs are not in a valid json format thus can't be used as is")
	}

	return imageEnvs, nil
}

//GetImageResources fetch the container resources from the docker image
func GetImageResources(imageName string) (*Resources, error) {
	labels, err := getImageLabels(imageName)
	if err != nil {
		return nil, err
	}

	resourcesToReturn, err := fetchLabelsFor(RESOURCES, labels)
	if err != nil {
		return nil, err
	}

	imageResources := parseResources(resourcesToReturn)
	json, _ := json.Marshal(imageResources)
	valid, err := validateMyObjectWithSchema("resources.json", string(json), "raw")
	if err != nil {
		log.Error(err.Error())
	}

	//TODO improve or decouple validation
	if !valid {
		log.Warn("your Resources are not in a valid json format thus can't be used as is")
	}
	return imageResources, nil
}

//GetImageExposedPorts fetch the texposed ports from the docker image
func GetImageExposedPorts(imageName string) (*Ports, error) {
	//TODO
	labels, err := getImageLabels(imageName)
	if err != nil {
		return nil, err
	}

	labelsToReturn, err := fetchLabelsFor(EXPOSED, labels)
	if err != nil {
		return nil, err
	}
	imagePorts := parsePorts(labelsToReturn)

	// Make sure that the envs conform with the schema
	json, _ := json.Marshal(imagePorts)
	valid, err := validateMyObjectWithSchema("ports.json", string(json), "raw")
	if err != nil {
		log.Error(err.Error())
	}

	//TODO improve or decouple validation
	if !valid {
		log.Warn("your Ports are not in a valid json format thus can't be used as is")
	}

	return imagePorts, nil

}

//GetImageTags fetch the tags used in the docker image
func GetImageTags(imageName string) (*Tags, error) {
	labels, err := getImageLabels(imageName)
	if err != nil {
		return nil, err
	}

	tagsToReturn, err := fetchLabelsFor(TAGS, labels)
	if err != nil {
		return nil, err
	}

	imageTags := parseTags(tagsToReturn)
	return imageTags, nil
}

//GetOtherImageTags gets all the docker image labels
func GetOtherImageTags(imageName string) (*Others, error) {
	labels, err := getImageLabels(imageName)
	if err != nil {
		return nil, err
	}
	otherLabels, err := fetchOthers(labels)
	return &otherLabels, err
}
