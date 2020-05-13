# Roadmap

## Upcoming

### v0.0.4

Support for multiple micro benchmark framework so core logic is more generic.
The adapter implementation could be simply shell out and collect structured output.

- go, the built in benchmark
- rust, criterion
- cpp, google/benchmark
- java, jmh

Skipped

- machine information
- metrics collected from system, e.g. total mem, cpu usage

- [ ] go
  - deal w/ sub benchmark, will this be the case for other benchmark framework?
  - skip label in sub benchmark `BenchmarkXXX/a=123,b=456/`
- [ ] storage
  - relational database (might use code generator as well for generating query...)
  - in memory test database (could use code generator)
- [ ] have Travis CI/github workflow working
  - remove deps on cobra and go.ice

### v0.0.5

- [ ] YCSB/tpc-c, maybe use go-ycsb instead of the java one
- [ ] save time series data in ?

Dependencies

- [ ] libtsdb-go if we want to store process resource usage like mem, cpu etc.

### v0.0.6

- [ ] initial support for xephon-b

## Finished
