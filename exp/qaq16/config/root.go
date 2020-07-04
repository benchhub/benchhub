package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func Read(p string) (Root, error) {
	b, err := ioutil.ReadFile(p)
	if err != nil {
		return Root{}, err
	}
	var cfg Root
	if err := yaml.Unmarshal(b, &cfg); err != nil {
		return Root{}, err
	}
	return cfg, nil
}

// Root is top level config.
type Root struct {
	Data       Data
	Contexts   []Context
	Score      Score
	Parameters []Parameter
	Containers []Container
}

type Data struct {
	DB  string `yaml:"db"`
	Dir string
}

type Context struct {
	Name  string
	Image string
}

type Score struct {
	Command Command
	Timeout string
	Capture string // The regex that capture the score
}

type Command struct {
	Shell string
	Dir   string
	Envs  []Env
}

type Parameter struct {
	Name    string
	Default int
}

type Container struct {
	Name     string
	Abstract bool
	Extends  string
	Image    string
	Resource Resource
	Envs     []Env
	Mounts   []Mount
}

type Resource struct {
	CPU int
	RAM string
}

type Env struct {
	Key   string
	Value string
}

// TODO: only bind mount is supported
type Mount struct {
	Src string
	Dst string
}
