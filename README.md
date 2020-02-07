# BenchHub

[![Build Status](https://travis-ci.org/benchhub/benchhub.svg?branch=master)](https://travis-ci.org/benchhub/benchhub)
[![Go Report Card](https://goreportcard.com/badge/github.com/benchhub/benchhub)](https://goreportcard.com/report/github.com/benchhub/benchhub)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fbenchhub%2Fbenchhub.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fbenchhub%2Fbenchhub?ref=badge_shield)

<h1 align="center">
	<br>
	<img width="200" src="https://avatars3.githubusercontent.com/u/32344687" alt="benchhub">
	<br>
	<br>
	<br>
</h1>

BenchHub is a service for running and storing benchmark result of databases in database.

- status: Under major rewrite. See [roadmap](doc/ROADMAP.md)

## Usage

NOTE: not in a usable state. See [_example](_example)

## License

MIT

[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fbenchhub%2Fbenchhub.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fbenchhub%2Fbenchhub?ref=badge_large)

## About

BenchHub is an implementation [@at15](https://github.com/at15)'s master thesis in [UCSC](https://www.ucsc.edu/) 'BenchHub: Store database benchmark result in database'.
It aims to define a spec for running distributed database benchmark and provides a continuous integration service.
Benchmark result is stored in database, it is structured and can be compared across sources.