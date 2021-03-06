# BenchHub

[![Build Status](https://travis-ci.org/benchhub/benchhub.svg?branch=master)](https://travis-ci.org/benchhub/benchhub)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fbenchhub%2Fbenchhub.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fbenchhub%2Fbenchhub?ref=badge_shield)

<h1 align="center">
	<br>
	<img width="200" src="https://avatars3.githubusercontent.com/u/32344687" alt="benchhub">
	<br>
	<br>
	<br>
</h1>

BenchHub is a service for running (database) benchmark and save result in databases.

- status: In active development. See [roadmap](doc/README.md#roadmap)

## Usage

Not usable at all, it won't even build due to using latest gommon with `replace` ...

## Alternatives

- https://openbenchmarking.org/
- https://cloud.google.com/profiler
- [Aliyun Performance Testing Service](https://www.aliyun.com/product/pts)
- https://github.com/conbench/conbench the repo is still README only, created by the creator of pandas

## License

MIT

[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fbenchhub%2Fbenchhub.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fbenchhub%2Fbenchhub?ref=badge_large)

## About

BenchHub is an implementation of [@at15](https://github.com/at15)'s master thesis in [UCSC](https://www.ucsc.edu/) 'BenchHub: Store database benchmark result in database'.
Its goal is defining a spec for running distributed database benchmark and providing a continuous integration service with managed runtime. 
BenchHub help database developer to focus on developing database itself instead of tools.
By saving benchmark result in databases, people can query and compare results across different times and sources.
