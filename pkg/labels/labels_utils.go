package labels

import (
	"fmt"
	"os"

	"github.com/docker/docker/builder/dockerfile/parser"
	log "github.com/zanetworker/dockument/pkg/log"
)

//DefaultEscapeToken escape token for dockerfile
const defaultEscapeToken = "\\"

func GetLabels(dockerfile string) error {
	labels := map[string]string{}
	nodeSearchFnc := func(f string, nodes []*parser.Node) error {
		for _, n := range nodes {
			searchFileFor("label", n, labels)
		}
		fmt.Println(len(labels))
		return nil
	}

	err := parseDockerfileNodes(nodeSearchFnc, dockerfile)
	if err != nil {
		return err
	}
	return nil
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
			fmt.Println(labels)
			n = n.Next.Next
		}

	}
	return labels
}
