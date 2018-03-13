package version

import (
	"encoding/json"
	"runtime"
	"time"
)

// Options are the options for the version command
type Options struct {
	Version   string `json:"Version"`
	GitCommit string `json:"GitCommit"`
	BuildDate string `json:"BuildDate,omitempty"`
	GOVersion string `json:"GOVersion,omitempty"`
	GOARCH    string `json:"GOARCH,omitempty"`
	GOOS      string `json:"GOOS,omitempty"`
	// out       io.Writer
}

//BuildVersion returns a formatted version
func (v *Options) BuildVersion(short bool) (string, error) {

	var versionToReturn string
	v.Version = "v" + getVersion()

	if short {
		v.GitCommit = getGitCommit()[:7]
		marshalledVersionShort, err := json.Marshal(v)
		if err != nil {
			return "", err
		}
		versionToReturn = string(marshalledVersionShort)
	} else {
		v.GitCommit = getGitCommit()
		v.BuildDate = time.Now().UTC().String()
		v.GOVersion = runtime.Version()
		v.GOARCH = runtime.GOARCH
		v.GOOS = runtime.GOOS
		marshalledVersionLong, err := json.Marshal(v)
		if err != nil {
			return "", err
		}
		versionToReturn = string(marshalledVersionLong)
	}

	return versionToReturn, nil
}

// //BuildVersion buils version
// func BuildVersion() string {
// 	if len(Version) == 0 {
// 		return "dev"
// 	}
// 	return Version
// }
