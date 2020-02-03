# Central

- bootstrap
  - node util, all nodes does this, regardless of agent or central
  - check data store, meta, ts
- run job
  - scheduler, schedule nodes
  - planner, expand the config using assigned node
  - dispatch job to agents
  - go to next stage when all nodes of one stage is finished

Long running

- http server
  - ui
- grpc server
  - accept call from agent
- job poller
  - pull queued job from store

Short running

- job executor