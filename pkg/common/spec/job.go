package spec

// Job describes a benchmark job, it's hardware requirement, how to run the workload
type Job struct {
	Name      string     `yaml:"name"`
	Reason    string     `yaml:"reason"`
	Framework string     `yaml:"framework"`
	Workload  string     `yaml:"workload"`
	Database  string     `yaml:"database"`
	Nodes     []Node     `yaml:"nodes"`
	Stages    []Stage    `yaml:"stages"`
	Pipelines []Pipeline `yaml:"pipelines"`
}
