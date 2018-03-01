package spec

import "github.com/dyweb/gommon/errors"

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
	// TODO: re unmarshal based on driver
	Config interface{} `yaml:"config"`
}

type taskAlias struct {
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
	// TODO: re unmarshal based on driver
	Config interface{} `yaml:"config"`
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
	// TODO: ports
}

func (t *Task) UnmarshalYAML(unmarshal func(interface{}) error) error {
	log.Info("task unmarshal called")
	// NOTE: use alias to avoid endless loop
	var alias taskAlias
	if err := unmarshal(&alias); err != nil {
		return err
	}
	switch alias.Driver {
	// FIXME: this does not work, still go back to map interface ...
	case "shell":
		alias.Config = ShellConfig{}
	case "docker":
		alias.Config = DockerConfig{}
	default:
		return errors.New("unknown driver")
	}
	if err := unmarshal(&alias); err != nil {
		return err
	}
	// copy everything
	t.Background = alias.Background
	t.Driver = alias.Driver
	t.Ready = alias.Ready
	t.Shutdown = alias.Shutdown
	t.Config = alias.Config
	return nil
}
