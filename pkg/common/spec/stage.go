package spec

import "github.com/dyweb/gommon/errors"

type Stage struct {
	Name      string `yaml:"name"`
	Selectors []Node `yaml:"selectors"`
	// Background specify if there are background task in this stage, a stage with background does not stop until specified
	Background bool                   `yaml:"background"`
	Tasks      []Task                 `yaml:"tasks"`
	XXX        map[string]interface{} `yaml:",inline"`
}

func (c *Stage) Validate() error {
	if c.XXX != nil {
		return errors.Errorf("undefined fields found in stage config %v", c.XXX)
	}
	merr := errors.NewMultiErr()
	for i := range c.Selectors {
		merr.Append(c.Selectors[i].Validate())
	}
	for i := range c.Tasks {
		merr.Append(c.Tasks[i].Validate())
	}
	return merr.ErrorOrNil()
}
