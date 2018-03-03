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

	// Drivers
	Shell  *Shell                 `yaml:"shell"`
	Exec   *Exec                  `yaml:"exec"`
	Docker *Docker                `yaml:"docker"`
	XXX    map[string]interface{} `yaml:",inline"`
}

// Exec run task using os/exec where Command is path of binary or can be found in PATH
type Exec struct {
	Command string   `yaml:"command"`
	Args    []string `yaml:"args"`
	// TODO: environment variable
	XXX map[string]interface{} `yaml:",inline"`
}

// Shell run using sh -c
type Shell struct {
	Command string                 `yaml:"command"`
	XXX     map[string]interface{} `yaml:",inline"`
}

const (
	DockerPull = "pull"
	DockerRun  = "run"
)

// Docker run using docker
type Docker struct {
	Image string `yaml:"image"`
	// Action can be pull, run
	Action string                 `yaml:"action"`
	Ports  []Port                 `yaml:"ports"`
	XXX    map[string]interface{} `yaml:",inline"`
}

type Port struct {
	Guest int                    `yaml:"guest"`
	Host  int                    `yaml:"host"`
	XXX   map[string]interface{} `yaml:",inline"`
}

func (c *Task) Validate() error {
	if c.XXX != nil {
		return errors.Errorf("undefined fields found in task config %v", c.XXX)
	}
	merr := errors.NewMultiErr()
	for i := range c.Ready {
		merr.Append(c.Ready[i].Validate())
	}
	for i := range c.Shutdown {
		merr.Append(c.Shutdown[i].Validate())
	}
	switch c.Driver {
	case "":
		merr.Append(errors.New("must specify driver for task"))
	case "shell":
		if c.Shell == nil {
			merr.Append(errors.New("shell is chosen but config is empty"))
		} else {
			merr.Append(c.Shell.Validate())
		}
	case "exec":
		if c.Exec == nil {
			merr.Append(errors.New("exec is chosen but config is empty"))
		} else {
			merr.Append(c.Exec.Validate())
		}
	case "docker":
		if c.Docker == nil {
			merr.Append(errors.New("docker is chosen but config is empty"))
		} else {
			merr.Append(c.Docker.Validate())
		}
	default:
		merr.Append(errors.Errorf("unknown driver %s", c.Driver))
	}
	return merr.ErrorOrNil()
}

func (c *Shell) Validate() error {
	if c.XXX != nil {
		return errors.Errorf("undefined fields found in shell config %v", c.XXX)
	}
	return nil
}

func (c *Exec) Validate() error {
	if c.XXX != nil {
		return errors.Errorf("undefined fields found in exec config %v", c.XXX)
	}
	return nil
}

func (c *Docker) Validate() error {
	if c.XXX != nil {
		return errors.Errorf("undefined fields found in docker config %v", c.XXX)
	}
	return nil
}
