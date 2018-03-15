package labels

import (
	"os"
	"regexp"

	"github.com/docker/docker/builder/dockerfile/parser"
	log "github.com/zanetworker/dockument/pkg/log"
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

func parseDockerfileNodes(searchFnc func(string, []*parser.Node) error, filename string) error {
	log.Info("Parsing Dockerfile:" + filename)

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

func searchFileFor(search string, n *parser.Node, labels map[string]string) map[string]string {
	defer func() map[string]string {
		if r := recover(); r != nil {
			return labels
		}
		return nil
	}()
	if n.Value == search {
		for {
			key, value := n.Next.Value, n.Next.Next.Value
			labels[key] = value
			n = n.Next.Next
		}
	}

	return labels
}

func fetchLabelsFor(labelType string, labelMap map[string]string) (map[string]string, error) {
	fetchedLabelsToReturn := map[string]string{}
	for key, value := range labelMap {
		switch labelType {
		case "DEPENDENCY":
			r, err := regexp.Compile(`^api.DEPENDENCY.*$`)
			if err != nil {
				return nil, err
			}
			if r.MatchString(key) == true {
				fetchedLabelsToReturn[key] = value
			}
		case "ENVS":
			r, err := regexp.Compile(`^api.ENVS.*$`)
			if err != nil {
				return nil, err
			}
			if r.MatchString(key) == true {
				fetchedLabelsToReturn[key] = value
			}
		case "EXPOSED":
			r, err := regexp.Compile(`^api.EXPOSED.*$`)
			if err != nil {
				return nil, err
			}
			if r.MatchString(key) == true {
				fetchedLabelsToReturn[key] = value
			}
		case "RESOURCES":
			r, err := regexp.Compile(`^api.RESOURCES.*$`)
			if err != nil {
				return nil, err
			}
			if r.MatchString(key) == true {
				fetchedLabelsToReturn[key] = value
			}
		case "TAGS":
			r, err := regexp.Compile(`^api.TAGS.*$`)
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
