package labels

import (
	"errors"
	"os"
	"regexp"
	"strings"

	"github.com/docker/docker/builder/dockerfile/parser"
	log "github.com/sirupsen/logrus"
	"github.com/xeipuuv/gojsonschema"
	"github.com/zanetworker/dockument/pkg/utils"
)

//DefaultEscapeToken escape token for dockerfile
const defaultEscapeToken = "\\"

func getLabels(dockerfile string) (map[string]string, error) {
	labels := map[string]string{}
	nodeSearchFnc := func(f string, nodes []*parser.Node) error {
		for _, n := range nodes {
			searchFileFor("label", n, labels)
		}
		return nil
	}

	err := parseDockerfileNodes(nodeSearchFnc, dockerfile)
	if err != nil {
		return nil, err
	}

	return labels, nil
}

func getFrom(dockerfile string) (map[string]string, error) {
	from := map[string]string{}
	nodeSearchFnc := func(f string, nodes []*parser.Node) error {
		for _, n := range nodes {
			searchFileFor("from", n, from)
		}
		return nil
	}
	err := parseDockerfileNodes(nodeSearchFnc, dockerfile)
	if err != nil {
		return nil, err
	}

	return from, nil

}

func parseDockerfileNodes(searchFnc func(string, []*parser.Node) error, filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	d := parser.Directive{LookingForDirectives: true}
	parser.SetEscapeToken(defaultEscapeToken, &d)
	result, err := parser.Parse(f, &d)
	if err != nil {
		return err
	}
	nodes := []*parser.Node{result}
	if result.Children != nil {
		nodes = append(nodes, result.Children...)
	}
	if err := searchFnc(filename, nodes); err != nil {
		return err
	}
	return nil
}

func searchFileFor(search string, n *parser.Node, items map[string]string) map[string]string {
	defer func() map[string]string {
		if r := recover(); r != nil {
			return items
		}
		return nil
	}()

	if n.Value == search {
		switch n.Value {
		case "label":
			for {
				// the end of this loop will panic and catched to return (one way to loop this data structure)
				key, value := n.Next.Value, n.Next.Next.Value
				items[key] = value
				n = n.Next.Next
			}
		case "from":
			imageValue := strings.Split(n.Next.Value, " ")[0]
			items["image"] = imageValue
			return items
		}

	}
	return items
}

func fetchLabelsFor(labelType string, labelMap map[string]string) (map[string]string, error) {
	fetchedLabelsToReturn := map[string]string{}
	for key, value := range labelMap {
		switch labelType {
		case DEPENDENCY:
			r, err := regexp.Compile(`^api.DEPENDENCY.*$`)
			if err != nil {
				return nil, err
			}
			if r.MatchString(key) == true {
				fetchedLabelsToReturn[key] = value
			}
		case ENVS:
			r, err := regexp.Compile(`^api.ENV.*$`)
			if err != nil {
				return nil, err
			}
			if r.MatchString(key) == true {
				fetchedLabelsToReturn[key] = value
			}
		case EXPOSED:
			r, err := regexp.Compile(`^api.EXPOSE.*$`)
			if err != nil {
				return nil, err
			}
			if r.MatchString(key) == true {
				fetchedLabelsToReturn[key] = value
			}
		case RESOURCES:
			r, err := regexp.Compile(`^api.RESOURCES.*$`)
			if err != nil {
				return nil, err
			}
			if r.MatchString(key) == true {
				fetchedLabelsToReturn[key] = value
			}
		case TAGS:
			r, err := regexp.Compile(`^api.TAGS.*$`)
			if err != nil {
				return nil, err
			}
			if r.MatchString(key) == true {
				fetchedLabelsToReturn[key] = value
			}
		case COMMAND_TESTS:
			r, err := regexp.Compile(`^api.TEST.command.*$`)
			if err != nil {
				return nil, err
			}
			if r.MatchString(key) == true {
				fetchedLabelsToReturn[key] = value
			}
		case FILE_CONTENT_TESTS:
			r, err := regexp.Compile(`^api.TEST.fileContent.*$`)
			if err != nil {
				return nil, err
			}
			if r.MatchString(key) == true {
				fetchedLabelsToReturn[key] = value
			}

		case FILE_EXISTENCE_TESTS:
			r, err := regexp.Compile(`^api.TEST.fileExistence.*$`)
			if err != nil {
				return nil, err
			}
			if r.MatchString(key) == true {
				fetchedLabelsToReturn[key] = value
			}

		case METADATA_TESTS:
			r, err := regexp.Compile(`^api.TEST.metadata.*$`)
			if err != nil {
				return nil, err
			}
			if r.MatchString(key) == true {
				fetchedLabelsToReturn[key] = value
			}
		}
	}
	return fetchedLabelsToReturn, nil
}

func fetchOthers(labelMap map[string]string) (Others, error) {
	fetchedLabelsToReturn := Others{}
	for key, value := range labelMap {
		if !keyMatchesPatterns(key) {
			fetchedLabelsToReturn[key] = value
		}
	}
	return fetchedLabelsToReturn, nil
}

func keyMatchesPatterns(keyToMatch string) bool {
	//TODO: handle errors correctly
	rDep, _ := regexp.Compile(`^api.DEPENDENCY.*$`)
	rEnv, _ := regexp.Compile(`^api.ENV.*$`)
	rExpose, _ := regexp.Compile(`^api.EXPOSE.*$`)
	rResources, _ := regexp.Compile(`^api.RESOURCES.*$`)
	rTags, _ := regexp.Compile(`^api.TAGS.*$`)
	rCommandTests, _ := regexp.Compile(`^api.TEST.command.*$`)
	rFileContentTests, _ := regexp.Compile(`^api.TEST.fileContent.*$`)
	rFileExistenceTests, _ := regexp.Compile(`^api.TEST.fileExistence.*$`)
	rMetadataTests, _ := regexp.Compile(`^api.TEST.metadata.*$`)

	if rDep.MatchString(keyToMatch) {
		return true
	}
	if rEnv.MatchString(keyToMatch) {
		return true
	}
	if rExpose.MatchString(keyToMatch) {
		return true
	}
	if rResources.MatchString(keyToMatch) {
		return true
	}
	if rTags.MatchString(keyToMatch) {
		return true
	}
	if rCommandTests.MatchString(keyToMatch) {
		return true
	}
	if rFileContentTests.MatchString(keyToMatch) {
		return true
	}
	if rFileExistenceTests.MatchString(keyToMatch) {
		return true
	}
	if rMetadataTests.MatchString(keyToMatch) {
		return true
	}
	return false

}

func validateMyObjectWithSchema(schema string, jsonTovalidate string, jsonContentType string) (bool, error) {
	var result *gojsonschema.Result
	var err error

	switch jsonContentType {
	case "file":
		schemaLoader := gojsonschema.NewReferenceLoader("file://" + utils.GetDir("api") + "/" + schema)
		documentLoader := gojsonschema.NewReferenceLoader("file://" + utils.GetDir("api") + "/" + jsonTovalidate)
		result, err = gojsonschema.Validate(schemaLoader, documentLoader)
		if err != nil {
			log.Error(err.Error())
			return false, err
		}
	case "raw":
		schemaLoader := gojsonschema.NewReferenceLoader("file://" + utils.GetDir("api") + "/" + schema)
		documentLoader := gojsonschema.NewStringLoader(jsonTovalidate)
		result, err = gojsonschema.Validate(schemaLoader, documentLoader)
		if err != nil {
			log.Error(err.Error())
			return false, err
		}
	default:
		return false, errors.New("Invalid type used, valid types are: \"file\" or \"raw\"")
	}

	return result.Valid(), nil
}
