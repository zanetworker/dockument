package labels

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	log "github.com/zanetworker/dockument/pkg/log"
)

func parseDependencies(dependencyLabels map[string]string) *Dependencies {
	dependencies := Dependencies{}
	for _, depMap := range divideMapByKey(dependencyLabels) {
		dependency := Dependency{}
		for dependencyLabel, value := range depMap {
			dependencyLabelStrings := strings.Split(dependencyLabel, ".")
			if len(dependencyLabelStrings) > 3 {
				dependencyParam := dependencyLabelStrings[3]
				switch dependencyParam {
				case "image":
					dependency.Image = value
				case "port":
					//TODO parse ports correctly
					dependency.Ports = []string{value}
				case "about":
					dependency.About = value
				case "mandatory":
					//TODO handle error correclty
					dependency.Mandatory, _ = strconv.ParseBool(value)
				}
			} else {
				dependency.Name = dependencyLabelStrings[2]
			}
		}
		dependencies = append(dependencies, dependency)
	}
	return &dependencies
}

func divideMapByKey(mapToDivide map[string]string) []map[string]string {
	mapsToReturn := []map[string]string{}
	for _, pattern := range getApplicationPatterns(mapToDivide) {
		tempMap := map[string]string{}
		r, err := regexp.Compile(fmt.Sprintf(`%s`, pattern))
		if err != nil {
			log.Error(err.Error())
		}
		for key, value := range mapToDivide {
			if r.MatchString(key) == true {
				tempMap[key] = value
			}
		}
		mapsToReturn = append(mapsToReturn, tempMap)
	}
	return mapsToReturn
}

func getApplicationPatterns(mapToDivideAndGroup map[string]string) []string {
	applicationsPattterns := []string{}
	for parentKey := range mapToDivideAndGroup {
		splittedKey := strings.Split(parentKey, ".")
		if len(splittedKey) == 3 {
			pattern := "^" + parentKey + ".*" + "$"
			applicationsPattterns = append(applicationsPattterns, pattern)
		}
	}
	return applicationsPattterns
}
