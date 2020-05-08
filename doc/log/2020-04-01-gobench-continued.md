# 2020-04-01 

Again ... it's been about 2 months since last work on gobench ...

## Previous work

- `_example/sort` contains merge sort code taken from pingcap talent plan to generate benchmark result
- parse text output of gobenchmark and convert to proto

## Major problems

- define the spec for client submission and server response
  - some fields are only populated on server side, e.g. id of new resource
- database schema
  - it would be good if proto can be saved directly in database, however even cloud spanner does not support it, so fields that requires filter etc. need to be duplicated as database columns
  - the schema is in a flux, doc and code can go out of sync easily
    - one way to solve it is define schema in go code and generates the markdown, it can be lifted to go.ice later
- data store implementation
  - only in memory store is running and the interface may not suit for RDBMS, and we may consider use DynamoDB to reduce storage cost in the future

## Current flow

Client 

- register
- run
- report

Server

- register
- report
- [ ] query data  

## TODO

- [ ] change spec and client to support other benchmark, e.g. rust, cpp 
- [ ] support using relational database, find a way to sync and check document, proto, database. Use docker to do unit test
- [ ] allow query result in cli and have a simple web UI

## Log

- want to clean up the package layout a bit, e.g. put `pkg/gobench` under `pkg/frameworks/gobench`
  - or maybe `framework/language/gobench` `framework/rdbms/tpcc` or I can call it `integration` ...
  - then initialize multiple grpc server? There could be two types of endpoints, core and integration
  - in that case the logic in store might become a bit complex ...
  - for simplicity, don't allow query across different frameworks for now, it is a valid use case for sure

common functionality

- register job, validate and return an id
  - core does common validation, auth, quota etc.
  - each framework should have its own validation on both client and server side
    - i.e. client send serialized user spec (with sensitive information removed?) to server
    - decoding and validation on client side is for better user experience
    - validation on server side is a must anyway ...
    - [ ] can we use Any?
- provide estimation for a job
  - core can provide a very rough estimation because it does not understand the parameters
  - framework can do estimation better our tell core how it can understand the parameters (give weights etc.)
- save actual result after client finish running the benchmark
  - [ ] assume we are not scheduling the resource for running the benchmark, it is a essential feature
- query historical data