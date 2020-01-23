package store_test

import (
	"context"
	"testing"

	"github.com/benchhub/benchhub/bhpb"
	"github.com/benchhub/benchhub/pkg/store"
	"github.com/dyweb/gommon/util/testutil"
	"github.com/stretchr/testify/assert"
)

func TestMetaMem_RegisterGoBenchmark(t *testing.T) {
	var spec bhpb.GoBenchmarkSpec
	testutil.ReadYAMLToStrict(t, "testdata/sort.yml", &spec)

	bkg := context.Background()
	mem := store.NewMetaMem()
	res, err := mem.RegisterGoBenchmark(bkg, &spec)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), res.SpecId)
	assert.Equal(t, int64(1), res.JobId)

	res, err = mem.RegisterGoBenchmark(bkg, &spec)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), res.SpecId)
	assert.Equal(t, int64(2), res.JobId)

	var spec2 bhpb.GoBenchmarkSpec
	testutil.ReadYAMLToStrict(t, "testdata/sort2.yml", &spec2)
	res, err = mem.RegisterGoBenchmark(bkg, &spec2)
	assert.Nil(t, err)
	assert.Equal(t, int64(2), res.SpecId)
	assert.Equal(t, int64(3), res.JobId)
}
