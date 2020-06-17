package config_test

import (
	"testing"

	"github.com/benchhub/benchhub/exp/qaq16/config"
	"github.com/stretchr/testify/require"
)

func TestRoot(t *testing.T) {
	cfg, err := config.Read("testdata/qaq15.yml")
	require.Nil(t, err)

	require.Equal(t, "go", cfg.Contexts[0].Name)
	require.Equal(t, "port", cfg.Containers[0].Envs[0].Key)
	require.Equal(t, "8081", cfg.Containers[0].Envs[0].Value)
	require.Equal(t, 20_000, cfg.Parameters[0].Default)
}
