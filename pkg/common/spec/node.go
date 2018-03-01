package spec

import "github.com/dyweb/gommon/errors"

// Node describes requirement for different types of nodes and individual nodes
type Node struct {
	// Name is used in config to refer specific node, it should be unique in job
	Name string `yaml:"name"`
	// Type specify type of nodes, currently only two types, database and loader
	Type string `yaml:"type"`
	// TODO: resource requirement
	// TODO: tags?
	XXX map[string]interface{} `yaml:",inline"`
}

// NodeSelector is used by stages to select nodes to run task on
type NodeSelector struct {
	Name string                 `yaml:"name"`
	Type string                 `yaml:"type"`
	XXX  map[string]interface{} `yaml:",inline"`
}

func (c *Node) Validate() error {
	if c.XXX != nil {
		return errors.Errorf("undefined fields found in node config %v", c.XXX)
	}
	return nil
}

func (c *NodeSelector) Validate() error {
	if c.XXX != nil {
		return errors.Errorf("undefined fields found in node selector config %v", c.XXX)
	}
	return nil
}
