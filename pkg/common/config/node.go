package config

type NodeProvider struct {
	Name     string `yaml:"name"`
	Region   string `yaml:"region"`
	Instance string `yaml:"instance"`
}

type NodeConfig struct {
	// Role is preferred role of this node, should be set based on instance type
	Role     string       `yaml:"role"`
	Provider NodeProvider `yaml:"provider"`
}
