# 2020-03-03 Gobench continued

It's been a month since last log, dig commits a bit to find where I was last time

- the schema is tracked is in [doc/design/gobench.md](../../../design/gobench.md), it also includes some general tables
- parse is improved to return proto directly and has a new `Converter` interface for transforming benchmark results
- nothing is done for relational database yet, depends on container test and need to recap how to use sql driver in go

## TODO

- [ ] use RDBMS to store go bench result, adjust the in memory store implementation and store interface accordingly
- [ ] fix travis
- [ ] allow estimate benchmark time before it starts