package labels

//Dependency object for the Dockerfile
type Dependency struct {
	About     string   `json:"about,omitempty"`
	Image     string   `json:"image,omitempty"`
	Mandatory bool     `json:"mandatory,omitempty"`
	Name      string   `json:"name,omitempty"`
	Ports     []string `json:"ports,omitempty"`
}

//Dependencies array of dependency objects
type Dependencies []Dependency
