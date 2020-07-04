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

- status: Under major rewrite. See [roadmap](doc/ROADMAP.md)

## Usage

NOTE: not in a usable state. See [_example](_example)

## Alternatives

- https://cloud.google.com/profiler
- https://www.aliyun.com/product/pts Aliyun Performance Testing Service
- https://github.com/conbench/conbench the repo is still README only, created by the creator of pandas

## License

MIT

[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fbenchhub%2Fbenchhub.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fbenchhub%2Fbenchhub?ref=badge_large)

## About

BenchHub is an implementation of [@at15](https://github.com/at15)'s master thesis in [UCSC](https://www.ucsc.edu/) 'BenchHub: Store database benchmark result in database'.
Its goal is defining a spec for running distributed database benchmark and provides a continuous integration service, so database developer can focus on developing database itself.
By saving benchmark result in databases, it allows people to query and compare results across different times and sources.