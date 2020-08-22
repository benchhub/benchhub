package generator_test

import (
	"testing"

	"github.com/benchhub/benchhub/lib/tqbuilder/generator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExtractPath(t *testing.T) {
	ep, err := generator.ExtractPath("core/services/user/db/spec/ddl", "github.com/benchhub/benchhub")
	require.Nil(t, err)
	assert.Equal(t, "user", ep.Package)
	assert.Equal(t, "core/services/user/db/generated", ep.OutputPath)
}
