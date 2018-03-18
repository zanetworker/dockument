package labels

import (
	"fmt"
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"
)

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
