package utils

import (
	"os"
	"path"
)

//GetDir gets diretory by name
func GetDir(dirToGet string) string {
	var projectPath = "/src/github.com/zanetworker/dockument/"

	switch dirToGet {
	case "root":
		return path.Join(os.Getenv("GOPATH") + projectPath)
	case "api":
		return path.Join(os.Getenv("GOPATH")+projectPath, "api")

	}
	return ""
}
