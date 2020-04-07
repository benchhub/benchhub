# Roadmap

## Upcoming

### v0.0.4

gobench without detail resource usage collected from system.
Main purpose is to have the run, submit result, generate report flow up and running. 

- [x] initial support for go micro benchmark
  - flatten sub benchmark in storage? or somehow keep the 
  - skip label in sub benchmark `BenchmarkXXX/a=123,b=456/`
  - skip machine information as well ... thought it's actually very important ...
- [ ] support C++ and rust and extract the logic into framework packages
- [ ] (high) relational database if we can create proper data model and test them properly in CI
- [ ] (low) in memory storage backend for easy testing
   - the reason in memory implementation is low priority is we can use container to run actual database in unit test
- [ ] have Travis CI working, need to fix the replace directive for go.ice
   - if gommon/dcli is working in 0.0.14, we can remove go.ice entirely

### v0.0.5

- [ ] initial support for xephon-b
   - xephon-b didn't got much effort for a long time ... otherwise I would like to do it before microbenchmark

### v0.0.6

- [ ] save machine spec
- [ ] save time series data from runtime (e.g. Hardware, OS metrics)
  - might need to use other tsdb if xephon-k is not ready ...

### v0.0.7

- [ ] support managed runtime using single vm on AWS

### v0.0.8

- [ ] support managed runtime on k8s

### v0.0.9

- [ ] initial support for YCSB, maybe use go-ycsb instead of the java one

### v0.0.10

- [ ] initial support for tpc-c

## Finished
