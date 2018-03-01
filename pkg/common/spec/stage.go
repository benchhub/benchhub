package spec

type Stage struct {
	Name      string `yaml:"name"`
	Selectors []Node `yaml:"selectors"`
	// Background specify if there are background task in this stage, a stage with background does not stop until specified
	Background bool   `yaml:"background"`
	Tasks      []Task `yaml:"tasks"`
}
