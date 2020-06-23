# v0.1.0 Micro

## Motivation

Have BenchHub up and running for micro benchmarks. It should be easier to use than database benchmark.
We can test it in [libtsdb](https://github.com/libtsdb), [gegecece](https://github.com/at15/gegecece) and [gommon](https://github.com/dyweb/gommon).

Originally we only want to support go but decided to support more than one languages so we can have a more generic framework interface.

- go, the built in benchmark
- rust, criterion
- cpp, google/benchmark
- java, jmh

Skipped

- machine information
- metrics collected from system, e.g. total mem, cpu usage

## Specs

- [ ] TODO: general workflow

## Features

## Components

### Core Storage

- RDBMS (might use code generator as well for generating query...)
- in memory test database (could use code generator)
- log
- profile data, e.g. pprof, perf

### Misc CI

- have Travis CI/github workflow working
- remove deps on cobra and go.ice and replace in go mod
  
### Framework Go

- deal w/ sub benchmark, will this be the case for other benchmark framework?
- skip label in sub benchmark `BenchmarkXXX/a=123,b=456/`
- pprof?