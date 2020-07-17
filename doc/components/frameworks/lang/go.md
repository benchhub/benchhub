# Go micro benchmark

NOTE: the following content is moved from old gobench.md under design.

## Workflow

Client

- allocate id
  - register spec and job
- run command
- parse go benchmark output
- send result back

Server

- allocate id
  - hash job spec so job with same spec can be compared
  - if spec does not exist
    - insert new spec
  - else
    - update last used time
  - insert new job
  - return id
- do nothing when client is running benchmark
- update database when result is sent back

## Spec

Spec is a client config defined in [proto](../../bhpb/gobench.proto)

## Schema

It requires ? tables

### Spec Table

This is a general table, i.e. applies to all benchmark frameworks.

- id, PK
- type, int, enum number in proto TODO: rename to adapter or framework?
- payload_hash, string, hash of spec, used to check
- payload, string, actual content of spec, encoded proto bytes, maybe use TEXT? (haven't use RDBMS for too long...)
- create_time, datetime, when is the spec first created
- last_used, datetime, most recent usage of this spec

### Job Table

This is a general table, i.e. applies to all benchmark frameworks.

- id, PK
- spec_id, int, should be foreign key but I am not sure if I want to use foreign key
- type, int
- create_time, datetime
- report_time, datetime
- start_time, datetime
- end_time, datetime
- duration, int, nanosecond?

### Gobench Table

This table stores actual benchmark result from go micro benchmark and maps to benchmark result output from `go test -bench`

- id, PK
- job_id, int
- spec_id, int
- package, string, full package path, e.g. `github.com/benchhub/benchhub/_example/sort`
- package_id, int

Fields from decoded result

- name, string, benchmark name e.g. `BenchmarkXXXXX`
- name_id, int
- core, int, TODO: does all benchmark has the num core suffix?
- ns_per_op, double, TODO: I think int should work ...
- allocated_bytes_per_op, int
- allocs_per_op, int
- mb_per_s, int
- measured, bitmap to keep track what is measured
- ord, int
- [ ] TODO: is it possible save duration? I think so ... just N * NsPerOp?

```go
// x/tools/benchmark/parse/parse.go

// Benchmark is one run of a single benchmark.
type Benchmark struct {
	Name              string  // benchmark name
	N                 int     // number of iterations
	NsPerOp           float64 // nanoseconds per iteration
	AllocedBytesPerOp uint64  // bytes allocated per iteration
	AllocsPerOp       uint64  // allocs per iteration
	MBPerS            float64 // MB processed per second
	Measured          int     // which measurements were recorded
	Ord               int     // ordinal position within a benchmark run
}
```

### Go Package Table

- id, PK
- package, string

### Go Benchmark Name Table

- id, PK
- name, string

## Reference

- https://github.com/golang/perf contains `benchstat`, `benchsave` and a small dashboard
  - uses https://github.com/golang/perf/blob/master/storage/benchfmt/benchfmt.go for read
  - [ ] didn't find if it has used golang.org/x/tools/benchmark/parse
    - https://github.com/golang/perf/blob/36b577b0eb03b831f9f591c1338a115cafcb56a7/benchstat/data.go#L212 seems it ignores header and decode everything to float64
  - it uses SQL database to keep uploaded record https://github.com/golang/perf/blob/master/storage/db/db.go
- https://github.com/aclements/go-perf-v2
- https://github.com/golang/benchmarks
- https://github.com/golang/proposal/blob/master/design/14313-benchmark-format.md
- https://github.com/knqyf263/cob GitHub Action that compares two recent commits
- [gRPC go dashboard](https://performance-dot-grpc-testing.appspot.com/explore?dashboard=5652536396611584&widget=490377658&container=1286539696)
