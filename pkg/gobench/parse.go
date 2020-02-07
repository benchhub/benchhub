package gobench

import (
	"bufio"
	"io"
	"os"

	"github.com/dyweb/gommon/errors"
	"golang.org/x/tools/benchmark/parse"
)

// TODO: there is no standard format that can be consumed by other program directly
// Ref:
// - standard format? https://github.com/golang/go/issues/14313
// - custom output? https://stackoverflow.com/questions/38038717/can-golang-benchmark-give-a-custom-output
// - https://github.com/aclements/go-misc it even has plot

func ParseFile(p string) ([]*parse.Benchmark, error) {
	f, err := os.Open(p)
	if err != nil {
		return nil, errors.Wrap(err, "error open go benchmark output")
	}
	defer f.Close()
	return Parse(f), nil
}

// TODO: change to ParseSet? Also can we keep labels of test output, e.g. go version etc.
func Parse(rr io.Reader) []*parse.Benchmark {
	var res []*parse.Benchmark
	r := bufio.NewReader(rr)
	for {
		l, err := r.ReadSlice('\n')
		// EOF
		if err != nil {
			break
		}
		b, err := parse.ParseLine(string(l))
		// Other test output
		if err != nil {
			continue
		}
		res = append(res, b)
	}
	return res
}
