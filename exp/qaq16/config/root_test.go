package config_test

import (
	"io/ioutil"
	"testing"

	"github.com/benchhub/benchhub/exp/qaq16/config"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
)

func TestRoot(t *testing.T) {
	var cfg config.Root
	b, err := ioutil.ReadFile("testdata/qaq15.yml")
	require.Nil(t, err)
	require.Nil(t, yaml.Unmarshal(b, &cfg))

	require.Equal(t, "go", cfg.Contexts[0].Name)
	require.Equal(t, "port", cfg.Containers[0].Envs[0].Key)
	require.Equal(t, "8081", cfg.Containers[0].Envs[0].Value)
}
