package labels

import "fmt"

// func GetLabels() {
// }

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
func GetDepenedencies(dockerfile string) (map[string]string, error) {
	labels, err := getLabels(dockerfile)
	if err != nil {

	}
	labelsToReturn, err := fetchLabelsFor(DEPENDENCY, labels)
	if err != nil {
		return nil, err
	}
	return labelsToReturn, nil
}

//GetEnvs fetch cotnainer environment variables from the Dockerfile
func GetEnvs() {

	schemaLoader := gojsonschema.NewReferenceLoader("file:///home/me/schema.json")
	documentLoader := gojsonschema.NewReferenceLoader("file:///home/me/document.json")

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		panic(err.Error())
	}

	if result.Valid() {
		fmt.Printf("The document is valid\n")
	} else {
		fmt.Printf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}
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
