# 2020-05-28 Ali contest quick hack

## Background

Thanks to Qin, I attended alibaba cloud native bla bla performance contest bla.
Question 1 is write a stream processing service on log data, it is distributed traces but the requirements make it irrelevant.

The workflow after the code is finished is like following.

```
for myPerf < topPerf {
    totalProcessTime = RunTest()
    lookAtLogAndProfile
    adjustParameter(chunkSize, batchSize, numThreads)
}
```

It would be good if I can automate this and let the machine do the grid search.
An example job config file is like the following:

```
parameters:
    numChunk:
        inject:
            env: "NUM_CHUNK"
        range:
            min: 1
            max: 10
    batchPerReq:
        inject:
            env: "BATCH_PER_REQ"
        range:
            min: 1
            max: 10
    batchSize:
        inject:
            env: "BATCH_SIZE"
        range:
            min: 20000
            max: 40000
            step: 1000
    dataSize: # TODO: actually we should treat parameters for input differently w/ parameters of our system
        enum:
            - 300m
            - 4g
spec:
    containers:
        c1:
            image: bla
            cpu: 2
            ram: 4g
        c2:
            image: bla
            cpu: 2
            ram: 4g
        b:
            image: bla
            cpu: 1
            ram: 2g
        score:
            image: bla
            mount: ....  # express how to extract metrics out, might expose a http endpoint
    timeout: 20s
```

To make things simple, we can skip database and use file system (and soft link as index) ...
Man ... I think using a RDBMS is better ... (wish gegecece is ready ...)

```
result
   2020-05-28-132357
        c1.log
        c2.log
        b.log
        merged.log
        score.txt
params
    numChunk
        best -> soft link to the best numChunk ... (though we need to way to control other variables)
        1
        2
          2020-05-28-1 // link to actual result
          2020-05-28-3 
        10
```

The result in database (or fs) should be able to answer the following query

- what's the best parameter combination (actually pretty easy, select * from results order by score desc limit 1)
- will increase chunk size help? (a bit hard ... it may or may not help depends on combination of other parameters)

The result in profile data (pprof)

- what's the top CPU usage func
- what's the top RAM usage func
- how does the top usage changes when parameter changes

## Design

This should be an adhoc implementation, and should help the actual design and impl of benchhub itself.
It's actually more like the original goal of benchboard where everything is stored locally.

- config format, hard coded for a docker compose like format (docker-compose v3 don't support resource usage limit meh.)
- runner, run on a single machine (i.e. my linux workstation) in serial (the tests are fast 10s)
- database, create table based on config (not ideal? this could be the storage requirement for benchhub)

## Implementation

- write it under `exp/qaq16` because the team name for the alibaba contest is `QAQ15`
- use MySQL as storage, embed SQLite makes compile go code painful ...