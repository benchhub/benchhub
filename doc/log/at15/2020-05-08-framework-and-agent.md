# 2020-05-08 Framework and Agent

- 2020-05-08 created initial draft when working on merge sort example

## Background

In current roadmap, supporting go micro benchmark is the major goal.
However, I find it's better to support more than one language so we can have a framework inside of something tied to go specific.

## Design

### Framework

- a framework in bhub matches to a benchmark framework, e.g. go, criterion(rust), xephon-b, tpc-c etc.
- each framework has its own data model
- there is a common part that can be extracted out all the frameworks
- there is a common part that can be extracted from frameworks of same ~~kind~~ trait
  - a framework may have more than one kind (or should I call it type ... maybe trait?)

### Agent

- agent is code embed into the application/system being benhmarked so it can export metrics/traces to benchhub
- we should provide integration for standard tools like perf, prometheus, opentracing/opencensus (... talking about community)
- low profile instrument for date intensive application like database (most existing tools may not be good enough)

### Core

BenchHub itself is the core and provides the following

- cli, `bh ctl` run benchmark locally, submit and query result.
- server, `bh srv` save and share benchmark result.
- UI (if I can wrap my head around ant design pro)
- runtime, `bh sched` run benchmark on cloud and spin up/down infra structure as needed. It can give concrete cost estimation and better scheduling based on historical benchmark data.
- tune, `bh tune` tune parameters and find most cost efficient solution.

### Relationship between core and framework

- core defines common traits that is implemented by different frameworks and can be used in saving data, schedule runtime and tune.
- common data is materialized, i.e. don't query framework

## Implementation

- agent is skipped because it's hard ...

```text
bhpb // contains proto for both core and framework
cmd
  bh // a single binary for everything
  bhgen // code generator, a separate binary so it compiles even if bh fails to do so ...
core
  framework // interface and framework registry
  server // grpc server and mounts all the frameworks
  store // storage engine interface used by all frameworks
frameworks
  lang // language specific micro benchmark
    gobench
    criterion
    googlebenchmark
  rdbms // relation database
    gotpc
    oltpbench
    sysbench
  kv
    ycsb  
  tsdb
    xephonb
test // integration, end to end test etc.
```

For the core data model and traits ... I think we can figure it along the way after supporting more than one frameworks.
There should be very big change after adding database benchmark, but I think we should be good for now ...