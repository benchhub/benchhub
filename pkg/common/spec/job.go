package spec

import "github.com/dyweb/gommon/errors"

// Job describes a benchmark job, it's hardware requirement, how to run the workload
type Job struct {
	Name      string                 `yaml:"name"`
	Reason    string                 `yaml:"reason"`
	Framework string                 `yaml:"framework"`
	Workload  string                 `yaml:"workload"`
	Database  string                 `yaml:"database"`
	Nodes     []Node                 `yaml:"nodes"`
	Stages    []Stage                `yaml:"stages"`
	Pipelines []Pipeline             `yaml:"pipelines"`
	XXX       map[string]interface{} `yaml:",inline"`
}

func (c *Job) Validate() error {
	if c.XXX != nil {
		return errors.Errorf("undefined fields found in job config %v", c.XXX)
	}
	merr := errors.NewMultiErr()
	// check node name dup
	nodeNames := make(map[string]bool, len(c.Nodes))
	for i := range c.Nodes {
		if nodeNames[c.Nodes[i].Name] {
			merr.Append(errors.Errorf("dup node name %s", c.Nodes[i].Name))
		}
		nodeNames[c.Nodes[i].Name] = true
		merr.Append(c.Nodes[i].Validate())
	}
	// check stage name dup
	stageNames := make(map[string]bool, len(c.Stages))
	for i := range c.Stages {
		if stageNames[c.Stages[i].Name] {
			merr.Append(errors.Errorf("dup stage name %s", c.Stages[i].Name))
		}
		stageNames[c.Stages[i].Name] = true
		merr.Append(c.Stages[i].Validate())
	}
	for i := range c.Pipelines {
		merr.Append(c.Pipelines[i].Validate())
	}
	return merr.ErrorOrNil()
}
