package spec

// Node describes requirement for different types of nodes and individual nodes
type Node struct {
	// Name is used in config to refer specific node, it should be unique in job
	Name string `yaml:"name"`
	// Type specify type of nodes, currently only two types, database and loader
	Type string `yaml:"type"`
	// TODO: resource requirement
	// TODO: tags?
}

// NodeSelector is used by stages to select nodes to run task on
type NodeSelector struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
}
