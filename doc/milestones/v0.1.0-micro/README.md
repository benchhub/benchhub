# v0.1.0 Micro

## Motivation

Have BenchHub up and running for micro benchmarks. It is used more often than database benchmarks.
We can test it in [libtsdb](https://github.com/libtsdb), [gegecece](https://github.com/at15/gegecece) and [gommon](https://github.com/dyweb/gommon).

Originally we only want to support go but decided to support more than one language, so we don't tied to go specific logic too much.

- go, the built in benchmark
- rust, criterion
- cpp, google/benchmark
- java, jmh

Skipped

- machine information
- metrics collected from system, e.g. total mem, cpu usage

## Specs

The general workflow of benchmark driven development is:

- find (or make up) a problem to solve
- write code
- write benchmark
- write benchhub config
- run test to make sure the code compiles and looks correct (it's test, not proof)
- run benchmark with a set of parameters
- view the report that contains metrics, log and profile
- compare with previous runs/related benchmarks
- annotate report, write notes, highlight/tag result
- iterate again with different code/parameters

## Features

### Register Benchmark

Description

Before user submit a benchmark result (or a request to run on managed runtime), they need to put the code in a place that can unique identify the benchmarks.
It's natural to follow github's layout of `org|user/project`, but there can be many benchmarks under one project.
Further more some benchmarks contains (conditional) sub-benchmark. We need to allow user to specify or scan the benchmarks.

Components

- config
  - configuration to describe a project
  - filter out/in specific benchmarks to run/report
  - package specific config
- storage
  - meta for project and benchmark, it should be general and same for all the benchmarks

## Components

### Core Storage

Structured

- RDBMS

Text

- log

Semi-Structured

- profile data, e.g. pprof, perf

Time series

- (long) metrics

### Framework Go

- parse output of go benchmark
  - sub benchmark, does other frameworks have things like this?
  - label in sub benchmark, I never used them `BenchmarkXXX/a=123,b=456/`
- save pprof data
  - **extremely** useful

### Ops CI

- have Travis CI/github workflow working
- remove deps on cobra and go.ice and replace in go mod