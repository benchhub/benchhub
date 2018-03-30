package config

import (
	"time"
)

type JobConfig struct {
	PollInterval time.Duration `yaml:"pollInterval"`
}
