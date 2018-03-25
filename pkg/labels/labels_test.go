package labels

import "testing"

func TestGetDependencies(t *testing.T) {

	// testMap := map[string]string{"api.DEPENDENCY.redis": "",
	// 	"api.DEPENDENCY.redis.image": "redis:latest",
	// 	"api.DEPENDENCY.adel":        "",
	// 	"api.DEPENDENCY.adel.image":  "redis:latest",
	// }
	// fmt.Printf("%#v", parseDependencies(testMap))
	// return nil, nil
}

func TestGetEnvVars(t *testing.T) {
	//TODO:
}

func TestGetExposedPorts(t *testing.T) {
	//TODO:
}

func TestGetTags(t *testing.T) {
	//TODO:
}

func TestGetAllLabels(t *testing.T) {
	//TODO:
}

// func TestVersion(t *testing.T) {
// 	tests := []struct {
// 		name     string
// 		id       string
// 		expected string
// 	}{
// 		{
// 			name:     "Empty Version Means BuildVersion Returns Dev",
// 			id:       "EmptyVersionMeansBuildVersionReturnsDev",
// 			expected: "dev",
// 		},
// 		{
// 			name:     "Version Returned FromFrom Build Version",
// 			id:       "VersionReturnedFromBuildVersion",
// 			expected: "testing-manual",
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			var actual, errMsg string

// 			switch tt.id {
// 			case "EmptyVersionMeansBuildVersionReturnsDev":
// 				actual = BuildVersion()
// 				errMsg = fmt.Sprintf("Version is not from Build - want: %s, got: %s", tt.expected, actual)

// 			case "VersionReturnedFromBuildVersion":
// 				Version = "testing-manual"
// 				actual = BuildVersion()
// 				errMsg = fmt.Sprintf("Version is not from Build - want: %s, got: %s", tt.expected, actual)
// 			}

// 			testutils.Equals(t, tt.expected, actual, errMsg)
// 		})
// 	}
// }
