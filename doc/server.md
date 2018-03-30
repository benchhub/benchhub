# Server

## Shared

common endpoints

- http
  - ping
  - node info + node status
- grpc
  - ping
  - node info + node status

logic for starting server

- bootstrap
  - read config
  - check environment
- start background go routines
- start server
  - grpc
  - http

## Central

- meta store
- job poller

## Agent

- heartbeat