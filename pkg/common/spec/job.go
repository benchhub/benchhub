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
	for i := range c.Nodes {
		merr.Append(c.Nodes[i].Validate())
	}
	for i := range c.Stages {
		merr.Append(c.Stages[i].Validate())
	}
	for i := range c.Pipelines {
		merr.Append(c.Pipelines[i].Validate())
	}
	return merr.ErrorOrNil()
}
