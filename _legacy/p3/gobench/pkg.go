// Package gobench implements go benchmark specific logic like parsing output
package gobench

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// TODO: this should be moved to gommon or at least go.ice
func LoadYAMLTo(f string, v interface{}) error {
	b, err := ioutil.ReadFile(f)
	if err != nil {
		return err
	}
	if err := yaml.UnmarshalStrict(b, v); err != nil {
		return err
	}
	return nil
}
