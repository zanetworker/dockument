package labels

//Dependency object for the Dockerfile
type Dependency struct {
	Name      string   `json:"name,omitempty"`
	Image     string   `json:"image,omitempty"`
	Ports     []string `json:"ports,omitempty"`
	About     string   `json:"about,omitempty"`
	Mandatory string   `json:"mandatory,omitempty"`
}

//Dependencies array of dependency objects
type Dependencies []Dependency

//CommandTest a command container structure test
type CommandTest struct {
	Name           string   `json:"name,omitempty"`
	Command        string   `json:"command,omitempty"`
	Setup          string   `json:"setup,omitempty"`
	Args           []string `json:"args,omitempty"`
	About          string   `json:"about,omitempty"`
	ExpectedOutput string   `json:"expectedOutput,omitempty"`
	ExpectedError  string   `json:"expectedError,omitempty"`
	ExcludedOutput string   `json:"excludedOutput,omitempty"`
	ExcludedError  string   `json:"excludedError,omitempty"`
}

//CommandTests array of command tests objects
type CommandTests []CommandTest

//FileContentTest a file content container structure test
type FileContentTest struct {
	Name             string `json:"name,omitempty"`
	Path             string `json:"path,omitempty"`
	ExpectedContents string `json:"expectedContents,omitempty"`
	ExcludedContents string `json:"excludedContents,omitempty"`
}

//FileContentTests array of file existence test objects
type FileContentTests []FileContentTest

//FileExistenceTest a file content container structure test
type FileExistenceTest struct {
	Name        string `json:"name,omitempty"`
	Path        string `json:"path,omitempty"`
	ShouldExist string `json:"shouldExist,omitempty"`
	Permissions string `json:"permissions,omitempty"`
}

//FileExistenceTests array of file exsistence tests
type FileExistenceTests []FileExistenceTest

//MetadataTest a test to make sure important meta-data exists in the container
type MetadataTest struct {
	Env          string `json:"env,omitempty"`
	ExposedPorts string `json:"exposedPorts,omitempty"`
	Volumes      string `json:"volumes,omitempty"`
	EntryPoint   string `json:"entrypoint,omitempty"`
	Cmd          string `json:"cmd,omitempty"`
	Workdir      string `json:"workdir,omitempty"`
}

//MetadataTests array of metadata tests
type MetadataTests []MetadataTest

//Env is the Environment Variable object for the Dockerfile
type Env struct {
	Name      string `json:"name,omitempty"`
	About     string `json:"about,omitempty"`
	Mandatory string `json:"mandatory,omitempty"`
}

//Envs array of environment variable objects
type Envs []Env

//Port is the exposed port object for the Dockerfile
type Port struct {
	Name     string `json:"name,omitempty"`
	About    string `json:"about,omitempty"`
	Scheme   string `json:"scheme,omitempty"`
	Protocol string `json:"protocol,omitempty"`
}

//Ports that should be exposed by the container application
type Ports []Port

//Resources is the exposed port object for the Dockerfile (only one resource object per dockerfile )
type Resources struct {
	CPU    string `json:"cpu,omitempty"`
	Memory string `json:"memory,omitempty"`
}

//Tags tags from dockerfile
type Tags map[string]string

//Others all other tags regardless of their type in the dockerfile
type Others map[string]string
