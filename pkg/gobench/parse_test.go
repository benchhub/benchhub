package gobench_test

import (
	"testing"

	"github.com/benchhub/benchhub/pkg/gobench"
	"github.com/stretchr/testify/assert"
)

func TestParseFile(t *testing.T) {
	results, err := gobench.ParseFile("testdata/sort.txt")
	assert.Nil(t, err)
	assert.NotNil(t, results)
	assert.Equal(t, "BenchmarkStd-8", results[0].Name)
	assert.Equal(t, uint64(64), results[0].AllocedBytesPerOp)
}
