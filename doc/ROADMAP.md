# Roadmap

## v0.0.4

- [ ] initial support for go micro benchmark
  - [ ] may not support label in sub benchmark `BenchmarkXXX/a=123,b=456/`
  - [ ] might skip machine information as well ... thought it's actually very important ...
- [ ] in memory storage backend for easy testing
- [ ] relational database if we can create proper data model

Dependencies

- [ ] libtsdb-go if we want to store process resource usage like mem, cpu etc.

## v0.0.5

- [ ] initial support for YCSB, maybe use go-ycsb instead of the java one
- [ ] save time series data

## v0.0.6

- [ ] initial support for xephon-b