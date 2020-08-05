# 2020-06-16 Components

Just flush notes I wrote on the scratch paper while downloading games.

- provide runtime environment
  - scheduler
  - adapter for k8s and public cloud provider
- run the benchmark and collect results
  - integration w/ test & benchmark framework
  - metrics from monitoring operation system and/or container
  - profile like pprof, perf
  - distributed traces
  - (optional) failure injection
  - (optional) parameter tuning
- save results
  - storage
    - data retention
    - index
  - query interface/language (maybe just SQL)
- compute
  - estimate costs based on historical data i.e. pre compute
  - ML for tuning scheduler and future benchmarks