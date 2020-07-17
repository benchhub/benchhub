# v0.1.0 Micro

## TODO

- [ ] shift version number because split register user and register test & benchmark

## Related

- Children:
  - [v0.0.4 Register User](../v0.0.4-register-user)
  - [v0.0.5 Register Test](../v0.0.5-register-test)

## Motivation

Have BenchHub up and running for micro benchmarks. It is used more often than database benchmarks.
We can utilize and test features it in [libtsdb](https://github.com/libtsdb), [gegecece](https://github.com/at15/gegecece) and [gommon](https://github.com/dyweb/gommon).

Originally we only want to support go but decided to support more than one language, so we don't tied to go specific logic too much.
We plan to support the following languages in v0.1.0, the go framework would have most features while other frameworks only have partial implementation.

- go, builtin test & benchmark
- rust, builtin test & criterion
- cpp, catch2? & google/benchmark
- java, junit & jmh

## Specs

The general workflow of benchmark driven development is:

- find a problem to solve
- write code
- write test
- run test to make sure the code compiles and looks correct (w/ race, memory leak detector etc.)
- write benchmark
- write benchhub config
- run benchmark with a set of parameters
- view the report that contains metrics, log and profile
- compare with previous runs/related benchmarks
- look up interesting results from knowledge base (by human and/or by tool)
- annotate report, write notes, highlight/tag result and update the knowledge base
- iterate again with different code/parameters

Proposed examples are:

- sort
  - based on pingcap talent plan v1
  - mergesort
  - quicksort (e.g. locality compared w/ mergesort)
- strstr
  - based on ali contest
  - index byte, last index byte e.g. go asm
- io
  - based on ali contest
  - read file and count the lines
  - http download file and count the lines
  - external merge sort?

## Features

Order by implementation order.

### v0.0.4 Register Benchmark

See [v0.0.4](../v0.0.4-register) for latest doc.

NOTE: pick 0.0.4 because 0.0.3 is the latest archive tag

Description

Before user submit a benchmark result (or a request to run on managed runtime), they need to put the code in a place that can unique identify the benchmarks.
It's natural to follow github's layout of `org|user/project`, but there can be many benchmarks under one project.
Further more some benchmarks contains (conditional) sub-benchmark. We need to allow user to specify or scan the benchmarks.
We must consider migrating the benchmark to a different location without losing data e.g. `benchhub/benchhub/_example` -> `benchhub/example`

Components

- `core/config`
  - can have a `META` file under each folder, used for describing project and its tests and benchmarks
  - configuration to describe a project
  - filter out/in specific benchmarks to run/report
  - package specific config ?what did I mean by that ....
- `core/storage/rdbms`
  - meta for project and benchmark, it should be general and same for all the benchmarks
- `cmd/bh`
  - the cli to register a benchmark e.g. `bh register .` `bh register ./...`
- `ui`
  - a project page to show registered projects, wow, js ... wow.... 呜呜呜 react redux 太难了 ...
- `test`
  - framework and util for running the entire stack locally
  
Internal Dependencies

- `core/storage`
  - a query builder, avoid using raw SQL in all places, define schema in go code? e.g. like `entgo`
  - this could go to `go.ice` though ...

External Dependencies

- `gommon`
  - `dcli` I don't want to use cobra and drag in a huge `go.sum` file ..., it's lucky that viper didn't k8s apiserver as kv store.
  - `noodle` need to bundle assets
- `reikla`
  - `rcl/rcl-go` Can use YAML for now ... just need a JSON with comment support

### v0.0.5 Run Test

Description

Before running benchmark, it's good to run some tests and see if things start falling.
Additional tools like memory leak, race detector should be enabled when running test.
When test failed, we may still go ahead and run benchmarks if that's what user wants.

Estimator is helpful when running test, if the test is going to take a long time, user can go brew some coffee.

Components

- `core/config`
  - specify what command/tests to run
- `frameworks/[go|rust]`
  - run tests locally
  - parse output for success/failure
- `core/estimator`
  - guess duration of entire test and individual test based on historical result (simply give an avg is enough)
- `core/storage/rdbms`
  - register tests
  - save test results, pass/fail, duration, log (if short)?
  - estimator related table
- `core/storage/log`
  - save metadata of log in rdbms or es?
  - save full log in a local fs, or a remote object store e.g. s3
  - maybe es?
- `cmd/bh`
  - `bh test .`
- `ui`
  - `/path/to/project/tests?id=xxx`
- `example/sort`
  - correct merge sort implementation in go and rust

### v0.0.6 Run Benchmarks

Description

BenchHub should support running BenchHub like GitHub supports Git protocol.
Without parameters and metrics (in general), benchmark is almost same as test.

Having parameters allows comparing different implementations in the same code base with different resource.
e.g. `nested loop join` vs `hash join`, `mergeSortParallel cpu=1, data=4g`, `mergeSortParalle cpu=2, ram=2g data=4g`
Estimator is more complex when using parameter. e.g. `mergeSort data=400m` and `mergeSort data=4g` should have very different estimation.
It is even more complex for multiple dimension, one naive way is to calculate average on all the historical data that match all dimensions.

Metrics like CPU usage, memory allocation gives more insight on why some code path is slow/fast.
Metrics can be extended to include profile, traces etc. Which give more inside and increase the cost for capturing and storage.
If user already know where the problem is, simple general metrics like CPU is good enough.
However, if users are exploring optimization opportunities, or the result is opposite from expectation.
Detailed information like profile and traces may lead to the right direction.  

Components

- `core/config`
  - specify parameters, changing schema of parameters may result in database schema change
  - specify benchmarks to run, some benchmarks may only use some parameters ...
- `frameworks/[go/rust]`
  - run benchmarks locally
  - parse output for CPU and RAM usage.
- `core/estimator`
  - estimate with parameters into consideration, naive avg ... (don't want to ML)
- `core/storage/rdbms`
  - register benchmarks
  - save benchmark results with different parameters (the data size and number of tables might explode ...)
  - parameter table? a standard set of parameters that can import? e.g. cpu, ram etc.
  - estimator table w/ parameters
- `core/storage/log`
  - should be same as test, parameters may have effect on file naming
  - filter log by parameters should be supported
- `core/storage/pprof`
  - save go pprof data, naive way is to save all the files w/ tgz ... though I think it's possible to compress it in columnar way
- `cmd/bh`
  - `bh bench .`
- `ui`
  - `/path/to/project/benchmarks?id=xxx&p1=xx&p2=xxx`
- `example/io`
  - go read file by line w/ pprof
  - rust read file by line, using pprof-rs
  
### ???

- share result to public i.e. there need to be a public hosted version of bh
  - authentication
- knowledge base
- annotate and refer to knowledge base 

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