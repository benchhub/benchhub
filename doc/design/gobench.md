# Gobench

Go bench is the adapter for running go microbenchmark and send result to benchhub.

- [proto](../../bhpb/gobench.proto)

## Workflow

Client

- allocate id
  - register spec and job
- run command
- parse go benchmark output
- send result back

Server

- allocate id
  - hash job spec so job with same spec can be compared
  - insert new job spec if it does not exist
  - insert new job
  - return id
- do nothing when client is running benchmark
- update database when result is sent back

## Spec



## Schema

