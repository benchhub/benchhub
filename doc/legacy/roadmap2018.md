# Roadmap

Previous deprecated roadmap in [UCSC Winter 2018](legacy/roadmap-winter18.md)

## 0.0.3

Structure BenchHub itself like a database
 
- state of the cluster can be considered as schema stored in catalog
- checking and populating benchmark spec is like validating and expanding SQL statement
- scheduling nodes is like data placement and cost estimation
- executing benchmark spec is like execute SQL statement (except most benchmark runs much longer)
- the dependencies between stages among multiple nodes is like 2PC
- benchhub can also be seen as a proxy for meta and time series store

## 0.0.2

- old client server style design, most components are finished, but are not well linked together
- the schema for store is still quite a mess
