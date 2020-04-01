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
