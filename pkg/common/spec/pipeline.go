package spec

import "github.com/dyweb/gommon/errors"

type Pipeline struct {
	Stages []string               `yaml:"stages"`
	XXX    map[string]interface{} `yaml:",inline"`
}

func (c *Pipeline) Validate() error {
	if c.XXX != nil {
		return errors.Errorf("undefined fields found in pipeline config %v", c.XXX)
	}
	return nil
}
