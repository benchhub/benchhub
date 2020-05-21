# 2020-05-21 Mergesort microbenchmark

## Background

This a follow up on [framework and agent](2020-05-08-framework-and-agent.md).
In order to get the common part of different micro benchmark framework in different programming languages.
We first need to write same benchmark in different programming languages which will give us the following:

- the report format of different micro benchmark frameworks
- a scenario to compare performance of different programming languages (and their standard library) on simple tasks i.e. merge sort
- tune performance by changing parameters (and do this exploration automatically)
  - ref https://github.com/at15/papers-i-read/issues/82 Google Vizier paper
  
## Design

- write mergesort in different programming languages, I already got a go implementation when doing pingcap talent plan.
  - need to learn rust, cpp, java.
- document the export format of those frameworks under `doc/frameworks` and figure out the common part
- convert the common and framework data models to proto and RDBMS schema
- implement a basic backend using gRPC and a single node RDBMS
- deploy it using a small instance e.g. AWS cloud sail and authenticate using github (save allowed user in github as text proto, lol such google, so piper)
  - runtime environment can use github action (not accurate but it works)
- after that we move on to the real work, database benchmark in shared distributed environment (survey on benchmark frameworks can happen in parallel)

## Implementation

Go

- [x] copy old code
- [ ] write doc about go bench format, also maybe consider pprof

C++

- [x] have cmake up and running

Rust

- [ ] port go code
- [ ] compare w/ std

Java

- [ ] init JMH (never used it before)