# Overview of BenchHub

BenchHub is a service for running database benchmark, result of benchmark is stored in database instead of plain text,
further more, BenchHub service is written as if it is a database.

## Terminology

Some naming in BenchHub is not very straight forward due to historical reason

- Central: master node
- Agent: worker node
- Spec: benchmark job specification, consider it as `.travis.yml`, it's not RFC
- Artifact: files like binary, dataset, config, report, git repo etc.

## Components

- Central
  - api server: talks with client and agent
  - scheduler: assign node base on job spec
  - planner: generate execution plan base on job spec
  - executor: dispatch plan to agents, monitor responses from agents
  - estimator: give user cost estimation before run the actual job so they can cancel it 
- Agent
  - executor: run plan got from central, start new process/container/vm
  - monitor: collect metrics, and tag/create new series running plan
- UI
  - a single tenant UI, though it does have route for login ...
- Cli
  - tool for submit job 

## Packages

NOTE: it may not reflect current package layout 

- lib
  - benchmark, wrapper/implementation for benchmark suite
  - monitor, use `/proc` to monitor host, use docker client to monitor container (might just use cgroup)
  - waitforit, wait until a tcp port or http endpoint is ready
- pkg/bhpb types defined using protobuf
- pkg/central central only logic
- pkg/agent agent only logic
- pkg/scheduler scheduler
- pkg/estimator estimator
- pkg/planner expand job spec to plan
- pkg/executor execute plan locally or in remote
- pkg/runner spawn process and container, maybe vm in the future

## Limitations

- no fail over for master
- scheduler is a simple FIFO for single tenant
  - need to change when becomes multi tenant
  - even for single tenant, it may not be optimal to schedule job in the order they came
  - also requires better implementation in estimator to aid scheduling in multi tenant mode