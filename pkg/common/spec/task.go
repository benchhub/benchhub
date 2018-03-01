package spec

type Task struct {
	// Background decides if the task is a long running process
	Background bool   `yaml:"background"`
	Driver     string `yaml:"driver"`
	// Ready defines commands to check the server is up and running
	// a background task is finished when all tasks in Ready can run without error
	// there can NOT be background task in Ready
	Ready []Task `yaml:"ready"`
	// Shutdown defines commands to run after the long running process is killed
	// TODO: though you can use docker stop to stop a background started by docker start ...
	Shutdown []Task `yaml:"shutdown"`

	// Drivers
	Shell  *ShellConfig  `yaml:"shell"`
	Exec   *ExecConfig   `yaml:"exec"`
	Docker *DockerConfig `yaml:"docker"`
}

// ExecConfig run task using os/exec where Command is path of binary or can be found in PATH
type ExecConfig struct {
	Command string   `yaml:"command"`
	Args    []string `yaml:"args"`
	// TODO: environment variable
}

// ShellConfig run using sh -c
type ShellConfig struct {
	Command string `yaml:"command"`
}

// DockerConfig run using docker
type DockerConfig struct {
	Image string `yaml:"image"`
	// Action can be pull, run
	Action string `yaml:"action"`
	Ports  []Port `yaml:"ports"`
}

type Port struct {
	Guest int `yaml:"guest"`
	Host  int `yaml:"host"`
}
