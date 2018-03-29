# Design

NOTE: this is a draft, and the eventual implementation might be much simpler than it

- bhubctl  command line tools for admin and debug
- bhubd  a mega service for web backend and scheduling (should split in the future, which is ... never)
  - user auth
  - github integration
  - explore benchmark results
  - ci master
- bhubagent node agent for running and monitor
  - communicate with master or other service (nomad? etc?)
  - run container/command
  - collect hardware/container metric
  - allow debug on specific machine

## directory layout

cmd simple`main.go` import package w/ side effect and cobra.Command, no real command logic 
  - bhubctl
  - ~~bhubd~~ bhubcentral
  - bhubagent
lib should be standalone package and can be used by other project, should not import internal packages, real common package should go to gommon
  - monitor will be reused by benchboard
  - ~~ice ? temporary home for go.ice ?~~ better not, just sync-local
  - tsdb should be libtsdb-go w/ just client implementation
pkg internal package
  - ~~bhubctl~~ ctl
    - cmd cobra command files
    - client? client wrapper for remote API (bhubd) or should both (bhubd and bhubagent) have client?
  - ~~bhubd~~ central
    - auth
    - github integration w/ GitHub
    - scheduler communicate w/ agent to decide when and were to run the job
    - db db operations ? .. TODO: tsdb related logic might also be needed here
    - server
      - apidoc.go
      - http.go
      - grpc.go
    - client 
      - TODO: go.ice should have good way for get most code for client library done when server logic is finished
        - go-kit is an example
        - error handling is inverse in server and client, server give error message when it got error, client return error by parsing server error message
  - ~~bhubagent~~ agent
    - server
    - client
    - runner
      - container
      - exec
    - monitor (collect host/container metrics)
    - debug (allow user to ssh into machine or even container ?...)
  - config config for many different packages to avoid cycle import in config
  - util
    - logutil