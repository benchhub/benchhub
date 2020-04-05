package gobench_test

import (
	"testing"

	"github.com/benchhub/benchhub/bhpb"
	"github.com/benchhub/benchhub/pkg/gobench"
	"github.com/stretchr/testify/assert"
)

func TestParseFile(t *testing.T) {
	results, err := gobench.ParseFile("testdata/sort.txt")
	assert.Nil(t, err)
	assert.NotNil(t, results)
	assert.Equal(t, "BenchmarkStd", results[0].Name)
	assert.Equal(t, uint64(64), results[0].BytesAllocatedPerOp)
}

func TestExtractCPU(t *testing.T) {
	b := &bhpb.GoBenchmarkResult{Name: "BenchmarkStd-8"}
	assert.Nil(t, gobench.ExtractCPU(b))
	assert.Equal(t, "BenchmarkStd", b.Name)
	assert.Equal(t, uint32(8), b.Cpu)
}
