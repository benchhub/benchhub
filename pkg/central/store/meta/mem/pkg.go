package mem

import (
	"github.com/benchhub/benchhub/pkg/central/store/meta"

	"github.com/benchhub/benchhub/pkg/util/logutil"
)

var log = logutil.NewPackageLogger()

func init() {
	meta.RegisterProviderFactory("mem", func() meta.Provider {
		return NewMetaStore()
	})
}
