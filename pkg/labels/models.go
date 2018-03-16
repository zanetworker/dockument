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

//Env is the Environment Variable object for the Dockerfile
type Env struct {
	Name      string `json:"name,omitempty"`
	About     string `json:"about,omitempty"`
	Mandatory string `json:"mandatory,omitempty"`
}

//Envs array of environment variable objects
type Envs []Env
