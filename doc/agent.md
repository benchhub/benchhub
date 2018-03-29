# Agent

- bootstrap
  - node util, get initial config, capacity
  - check if required software, i.e. docker is installed
- register
  - find central based on address in config
  - update state machine 
  - switch to heartbeat
- heartbeat
  - should send state of machine, running job etc.
  - switch to register if heartbeat fail
- run job
  - handler that accept calls from central
  - change state
  - call central after job (stage) is finished 
    - one node only run one job stage at a time, because coordination from central is needed
  - report metrics to ts store during job, attach job, task etc. in job
  
  
Long running

- http server
  - for ui
- grpc server
  - accept call from central
- heartbeat
  - register
  - keep alive

Short running

- job stage executor
- task executor