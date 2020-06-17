# qaq16 A parameterized docker runner

## Motivation

See [log/2020-05-28-ali-contest-quick-hack](../../doc/log/2020-05-28-ali-contest-quick-hack.md) for full background

- need a handy tool for [alibaba tianchi contest](https://tianchi.aliyun.com/competition/entrance/231790/introduction)
- `docker-compose` v3 does not allow setting resource constraint for single node
- need a simple prototype that could help actual benchhub implementation

## Requirement

For the given problem I have implementation in different programming languages (for the sake of learning rust & cpp, maybe java).
For same language I have different parameters, e.g. enable/disable gc, batch size, buffer size etc.

- each run should give a final score in float number (makes sorting and pick the best one much easier)
- log of each run should be kept on disk, so I can look at debug information
  - ideally we should able to extract metrics from log but that is too much work for now
  
```yaml
# NOTE: this doc may be out of date, see testdata/qaq15.yml for latest doc

data: # log and mysql volume
  db: mysql:8.0
  dir: .
contexts:
# NOTE: build and tag image externally, this is just a runner, does not build like docker-compose
    - name: go
      image: blabla:gov1
# TODO: should allow different default parameters for different context 
    - name: rust
      image: blabla:rsv1
    - name: cpp
      image: blabla:cppv1
score: # run in host directly, it should run in container as well, but I don't want bind mount
    capture: "final score is (\d+)" # regex
    timeout: 20s # shutdown if it is too slow
parameters:
  - name: batchSize
    default: 20_000
  - name: numChunk
    default: 1
containers:
    - name: f1
      image: context.image # might just skip that if all containers are using same image
      resource:
        cpu: 2
        ram: 4g
      envs:
        port: 8081
        batchSize: batchSize
    - name: f2
      image: context.image # TODO: avoid the dup, now just copy and paste
      envs:
        port: 8082
    - name: b
      image: context.image
```

## Design

```
main {
    contexts, params = readConfig
    
    createTableIfNotExists(contexts, params) // TODO: what about adding new fields when there are new parameters
    // id, context, score, batchSize, numChunk

    run(defaultParamters) // TODO: implement grid search instead of using default

    saveLogAndDatabase
}
```