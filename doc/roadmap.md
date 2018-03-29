# Roadmap

## 0.1

- use push style without retry, the whole job won't work if their is a single call goes wrong
  - central call agent directly instead of agent pull from central or watch store
  - agent call central directly after a task is finished
- use state machine to keep an updated view in memory for central and agent
- use pub/sub to record updates into store

## 0.2

- use push + pull, push is to notify event right away, pull (during heartbeat) will tell the state eventually

## 0.3

- works with other scheduler, i.e. Mesos, Kubernetes so it can run along side existing infrastructure

## 0.4

- github integration

## 0.5

- multi tenant
- cost estimator

## 0.6

- benchboard, non database benchmark, i.e. go's benchmark, jmh etc.