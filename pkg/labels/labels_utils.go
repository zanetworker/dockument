package labels

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/docker/docker/builder/dockerfile/parser"
)

func getLabel(c *cli.Context) error {
	labels := map[string]string 

	nodeSearch := func(f string, nodes []*parser.Node) error {
		for _, n := range nodes {
			labels = searchFor("label", n, labels)
			log.Println(images)
		}
		return nil
	}

	err := forFile(c, nodeSearch)
	if err != nil {
		return err
	}

	// setup the tab writer
	w := tabwriter.NewWriter(os.Stdout, 20, 1, 3, ' ', 0)

	// print header
	fmt.Fprintln(w, "BASE\tCOUNT")

	pl := rank(images)
	for _, p := range pl {
		fmt.Fprintf(w, "%s\t%d\n", p.Key, p.Value)
	}

	w.Flush()
	return nil
}


func forFile(c *cli.Context, fnc func(string, []*parser.Node) error) error {
	for _, fn := range c.Args() {
		logrus.Debugf("File: %s", fn)
		f, err := os.Open(fn)
		if err != nil {
			return err
		}

		result, err := parser.Parse(f)
		if err != nil {
			return err
		}
		ast := result.AST
		nodes := []*parser.Node{ast}
		if ast.Children != nil {
			nodes = append(nodes, ast.Children...)
		}
		if err := fnc(fn, nodes); err != nil {
			return err
		}
	}
	return nil
}

func searchFor(search string, n *parser.Node) map[string]string {

	if n.Value == search {
		if v, ok := a[n.Next.Value]; ok {
			a[n.Next.Value] = v + 1
		} else {
			log.Println(n.Next.Next.Value)
			a[n.Next.Value] = 1

		}
	}
	return a
}