# 2020-01-22 Data Store

First let's consider the requirement 

- each test job need to stored
  - the original job request
  - tags, e.g. size=micro lang=go size=e2e framework=ycsb, though if we have limited set of tags, just using column is fine
  - summary time_cost, money_cost
  - individual runs inside a big run, mainly caused by the matrix, e.g. cloud_provider=gcp
- raw time series is collected during monitoring (useful for e2e, may not be that helpful for micro benchmark)
- want know how much a benchmark will take (time, resource, money etc.)

Tables

Spec: keep track of original spec, so it can be used for estimation

