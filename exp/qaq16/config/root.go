package config

// Root is top level config.
type Root struct {
	Data       Data
	Contexts   []Context
	Score      Score
	Parameters []Parameter
	Containers []Container
}

type Data struct {
	DB  string `yaml:"db"`
	Dir string
}

type Context struct {
	Name  string
	Image string
}

type Score struct {
	Timeout string
	Capture string // The regex that capture the score
}

type Parameter struct {
	Name    string
	Default int
}

type Container struct {
	Image    string
	Resource Resource
	Envs     []Env
}

type Resource struct {
	CPU int
	RAM string
}

type Env struct {
	Key   string
	Value string
}
