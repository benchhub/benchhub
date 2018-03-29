# Roadmap

## 0.1

- use push style without retry, the whole job won't work if their is a single call goes wrong
  - central call agent directly instead of agent pull from central or watch store
  - agent call central directly after a task is finished
- use state machine to keep an updated view in memory for central and agent
- use pub/sub to record updates into store

## 0.2

- use push + pull, push is to notify event right away, pull (during heartbeat) will tell the state eventually