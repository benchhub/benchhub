# Roadmap

## Upcoming

### v0.0.4

gobench without detail resource usage collected from system.
Main purpose is to have the run, submit result, generate report flow up and running. 

- [ ] initial support for go micro benchmark
  - flatten sub benchmark in storage? or somehow keep the 
  - skip label in sub benchmark `BenchmarkXXX/a=123,b=456/`
  - skip machine information as well ... thought it's actually very important ...
- [ ] in memory storage backend for easy testing
- [ ] have Travis CI working, need to fix the replace directive for go.ice
- [ ] relational database if we can create proper data model and test them properly in CI

### v0.0.5

- [ ] initial support for YCSB, maybe use go-ycsb instead of the java one
- [ ] save time series data

Dependencies

- [ ] libtsdb-go if we want to store process resource usage like mem, cpu etc.

### v0.0.6

- [ ] initial support for xephon-b

## Finished
