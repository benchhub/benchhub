# 2020-04-07 Framework Init

## Background

As stated in [previous log](2020-04-01-gobench-continued.md), I want to split core and language/database specific logic out.

## Design

- core handles framework agnostic logic, provides storage interface
- framework handles integration with actual benchmark framework and save framework specific data

```text
cmd
    bh
bhpb
    lang
        gostd
            spec.proto // framework specific proto
            rpc.prot // framework specific rpc, client ONLY calls its framework's endpoint, glue response from core and framework should be handled in server.
core
    server // core server and common logic shared by application servers
    store // data store
    scheduler // when benchhub provides managed runtime
framework
    lang // <lang><framework>, didn't use <lang>-<framework> due to go package naming
        cppgoogle
        gostd
            cmd.go // decode framework specific proto config 
            server.go // grpc server that validates & save go benchmark data NOTE: you can add multiple grpc service in one actual server
    rdbms
        gotpcc
        oltpbench
    nosql
        ycsb
        goycsb
    tsdb
        xephonb
```

## Implementation

- first split existing gobench code out, or remove them entirely to make the code compile ...
- second list core schema
  - might sync schema, markdown, proto manually (until I can extract code generator logic out from pm)
- third list go bench specific schema
- fourth add cpp/rust (if I can have cmake stuff running ....)
