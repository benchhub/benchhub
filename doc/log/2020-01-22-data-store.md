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

Job: Every time we run a spec we produce a job? oe multiple jobs? 

A simple work flow, take pingcap's go training plan as example (because I just finished it recently ... if 2 weeks ago is recent)

Under $GOPATH/at15/pingcap-talent-plan/tidb/mergesort

```text
$ bh gobench
```

- look for a benchhub.yml to locate target servers
  - could look up like go mod or git but let's make things easier for now
- register the job to get a job id
  - the spec will be created if it does not exist yet
  - if there are historical jobs, and estimation is given back
- run the benchmark (using local executor)
- collect the result
- send back to server using that id

a simple config would be like

```yaml
# you can select multiple servers for running benchmark ... the would run them in parallel on same/different environment
servers:
    - name: local # unique name on client side, not send to server
      addr: localhost:1124
commands:
    - type: shell 
      command: make bench # the output is saved
      output: bench.txt
report:
    - format: gobench
      input: bench.txt
```
