# v0.0.5 Register Test

## TODO

Done

- [x] move register test from v0.0.4 register user

## Related

- Parent: [v0.1.0 Micro](../v0.1.0-micro)
- Prev: [v0.0.4 Register User](../v0.0.4-register-user)

## Motivation

Split from [v0.0.4 Register User](../v0.0.4-register-user). User registration logic already has enough boilerplate on databases, rpc etc.
Register test and benchmark focus more how to extract test name from test output. We may also take a bazel like approach, specify test target.

## Specs

- define test and benchmark target in `bh.rcl`
- get test and benchmark name by running test and extract name from output

## Implementation

- [ ] list actual impl order, though not in a rush cause the previous one is not done yet.
  - but gotest does not depend that much on previous work and _example/mergesort can be updated as well

## Features

- [Register Test](#register-test)
- [Register Benchmark](#register-benchmark)

### Register Test

Description

It's impractical to have people manually enter all the test/benchmark names.
The easiest solution is to run the test and parse the output.
The name of the test should be unique inside the project and relies on the framework.

Known issues

- conditional test won't work i.e. `go test -short`, this is a special case of parameter. We will consider that when we start actually running test/benchmark with parameters.

Components

- `core/config`
  - allow ignoring or force specific test patterns, it is opaque to bh and depends on framework
  - the command to run test, the command need to save the output in a place framework can find it (or stream to stdout/err)
- `core/storage/rdbms`
  - save test, saving all tests in a single table w/o partition may cause slow query in the end, but for now `id, project_id, test_name` should be good enough
- `cmd/bh`
  - `bh register test`
- `ui`
  - `host>owner>repo>project>bench`
  
### Register Benchmark

Same as [test](#register-test).

## Components

### Framework gotest

Description

Parse `go test` output

Used in

- [test](#register-test)

TODO

- [ ] there should be json output format for go test?
  - can try to use proto ...

### Framework gobench

Description

Parse `go test -bench` output

Used in

- [benchmark](#register-benchmark)

TODO

- [ ] parse go bench output? though it might be same as go test if we just need test name?

### Framework rusttest

Description

Get result from `cargo test`

Used in

- [test](#register-test)

TODO

- [ ] `cargo test` output format?