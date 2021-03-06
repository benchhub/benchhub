# Directory

- [cmd](../cmd), command line tools, entry point for daemon (`bhubcentral`, `bubagent`) ctl (`bhubctl`) etc.
- [doc](.) documents, survey and internals, few is about usage
  - [legacy](legacy) out dated doc
  - [survey](survey) scheduler, CI systems
  - [third_party](third_party) notes about using third party libraries
  - [util](util) utility for generating graphs
- [example](../example) example benchhub job spec, mostly not working
- [lib](../lib) libraries that will be split out as standalone repositories
  - [benchmark](../lib/benchmark) wrapper on existing benchmark frameworks, empty
  - [monitor](../lib/monitor) collect host and container (only docker) metrics
  - [runner](../lib/runner) task executor, moved to [pkg/common/runner](../pkg/common/runner) due to dep on spec
  - [waitforit](../lib/waitforit) check if tcp port, http endpoint is ready, also contains a cli
- [pkg](../pkg) all the internal packages, putting them in one folder makes using tools like `gofmt`easier
  - [agent](../pkg/agent) daemon runs on worker node, run tasks, collect metrics
  - [bhpub](../pkg/bhpb) protobuf definition and generated go structs, custom error and YAML unmarshalers
  - [central](../pkg/central) the single point of failure, daemon runs on central node, api, scheduler, ui, job coordinator, in memory store 
  - [ctl](../pkg/ctl) command line tool logic, will be deprecated, just use `cmd` package directly
  - [util](../pkg/util) logutil, control logging of itself and dependencies using gommon/log, nodeutil
- [script](../script) provision, local dev setup, hacky tools
- [ui](../ui) SPA using Angular + Ant.Design + Echarts

## Agent

- [config](../pkg/agent/config) config struct for client and server, in standalone package to avoid dependency cycle
- [job](../pkg/agent/job) execute job dispatched from central, not actual implementation
- [server](../pkg/agent/server) http, grpc server, heart beat, state machine, all the long running stuff
- [transport](../pkg/agent/transport) grpc rpc definition, in standalone package to avoid dependency cycle

## Central

- [config](../pkg/central/config) same as agent
- [job](../pkg/central/job) planner, executor, estimator (not implemented)
- [scheduler](../pkg/central/scheduler) scheduler interface and built in implementation
- [server](../pkg/central/server) http, grpc server, job poller
- [store](../pkg/central/store) meta and time series store
- [transport](../pkg/central/transport) same as agent

## TODO

- [x] move all package in `pkg/common` to `pkg`
- [ ] store might become a common package? though agent does not need meta store much, just ts store is enough
  - [ ] a central meta store interface?
- [x] ? split `bhpb` into `bhspec` `bhstore` `bhjob` etc. or find a better way to use the giant file ... don't want to have too many packages ...
  - https://github.com/sensu/sensu-go/tree/master/types is an example
  - it makes import harder, so just bear with the one giant proto file