# v0.0.4: bhpb

## TODO

- [ ] should we use different folders? i.e. split proto into its own framework
  - can be a good idea, unless framework need to share something (we can extract that out to a common part whe needed)
  - e.g. `core/config/configpb/config.proto`, btw: can run `protoc` using bind mount
- [ ] other definitions following roadmap
- [ ] relationship w/ database

Done

- [x] have protoc and the new proto go runtime up and running
  - [x] define a echo service, `hello` -> `benchhub: hello`. Always return at15 in `GetUser`
