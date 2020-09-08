# 2020-08-30 Featherweight BenchHub

## Background

The name of the log comes from TAPL (honestly I forgot everything except lambada calculus and featherweight java etc....)
Also, there is featherweight go ...

The main motivation is I don't have much time left to work on side projects (the gap year is coming to and end).
So I need to trim down benchhub so at least I can use it when I need it.

### The very minimal

- register user, project, repo automatically based on path
- run the test/benchmark directly w/o submit a job, i.e. `bh run <target-name>` runs in current shell instead of remote
- collect result and send it via rpc
- save the result in database
- have user write sql directly or use existing sql visualization tools e.g. metabase to render graph

### A bit advanced

- submit a job to local runtime
- runtime does what user did
  - run tests
  - send result via rpc
- plus
  - get a copy of code

### More advanced

- submit a job to k8s/vm runtime

## Design

Framework runs in 3 places

- client cli, e.g. decode config, generate proper payload, all local mode
- server, e.g. convert rpc and database (proto mover)
- runtime, e.g. process in local, sidecar in k8s

## Implementation

- bh core serve
  - check storage backend
  - load registry
  - start core services, users, projects etc.
  - start frameworks (that run along w/ core)
  - accept user request
- bh run <target-name>
  - locate and decode config file to find the target
  - all local mode
    - trigger framework to run the test
    - convert the output (result, log, metrics) and save it locally
    - report to remove server via rpc
  - local runtime
    - generates a job request and submit
    - runtime make a copy of the context (i.e. code) (or simply build a docker image)
    - the rest is same as all local mode
  - k8s runtime
    - build and push an image to registry to contain required code
    - generates a job request and submit
    - crd etc. if necessary, though simply a pod w/ sidecar should work as well
- gobench
  - go tool test json for test output
  - benchmark parse
  - resource usage might rely on core to monitor the specific go test binary
  
Dependencies

- format for defining spec and how to handle different frameworks (maybe refer blaze)
- database schema for frameworks, and a common schema (maybe fill it async in background)