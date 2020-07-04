package gobench

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/benchhub/benchhub/bhpb"

	"github.com/dyweb/gommon/errors"
	"golang.org/x/tools/benchmark/parse"
)

// TODO: there is no standard format that can be consumed by other program directly
// Ref:
// - standard format? https://github.com/golang/go/issues/14313
// - custom output? https://stackoverflow.com/questions/38038717/can-golang-benchmark-give-a-custom-output
// - https://github.com/aclements/go-misc it even has plot

func ParseFile(p string) ([]*bhpb.GoBenchmarkResult, error) {
	f, err := os.Open(p)
	if err != nil {
		return nil, errors.Wrap(err, "error open go benchmark output")
	}
	defer f.Close()
	return Parse(f)
}

// Parse calls parse.ParseSet and convert parse.Benchmark to bhpb.
// It applies the following converters
// - ExtractCPU
func Parse(r io.Reader) ([]*bhpb.GoBenchmarkResult, error) {
	var parsed []*parse.Benchmark
	scan := bufio.NewScanner(r)
	ord := 0
	for scan.Scan() {
		if b, err := parse.ParseLine(scan.Text()); err == nil {
			b.Ord = ord
			ord++
			parsed = append(parsed, b)
		}
	}
	if err := scan.Err(); err != nil {
		return nil, err
	}
	var converted []*bhpb.GoBenchmarkResult
	for _, b := range parsed {
		converted = append(converted, &bhpb.GoBenchmarkResult{
			Package:             "",
			Name:                b.Name,
			NOp:                 int64(b.N),
			NsPerOp:             b.NsPerOp,
			AllocPerOp:          b.AllocsPerOp,
			BytesAllocatedPerOp: b.AllocedBytesPerOp,
			MbPerSecond:         b.MBPerS,
			Measured:            int64(b.Measured),
			Ord:                 uint32(b.Ord),
			Duration:            int64(b.NsPerOp * float64(b.N)),
			Cpu:                 0,
		})
	}
	err := Convert(converted, ExtractCPU)
	return converted, err
}

// Converter modifies a benchmark result in place.
type Converter func(b *bhpb.GoBenchmarkResult) error

// ExtractCPU converts BenchmarkSort-8 to BenchmarkSort and save 8 in b.Cpu
func ExtractCPU(b *bhpb.GoBenchmarkResult) error {
	if b.Name == "" {
		return errors.New("can't extract cpu core from empty benchmark name")
	}
	i := strings.LastIndex(b.Name, "-")
	if i == -1 {
		return errors.New("-<cpu> not found")
	}
	cpu, err := strconv.Atoi(b.Name[i+1:])
	if err != nil {
		return errors.Wrap(err, "error convert cpu number to int")
	}
	b.Name = b.Name[:i]
	b.Cpu = uint32(cpu)
	return nil
}

// Convert calls all converters and bails out on first error.
func Convert(benchmarks []*bhpb.GoBenchmarkResult, converters ...Converter) error {
	for i, benchmark := range benchmarks {
		for j, converter := range converters {
			if err := converter(benchmark); err != nil {
				return errors.Wrapf(err, "convert %d benchmark failed on %d converter", i, j)
			}
		}
	}
	return nil
}
