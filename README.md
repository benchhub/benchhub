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

## Usage

````bash
make install
bhubcentral serve
bhubagent serve
# on another machine, change bhubagent.yml to point to the right central
bhubagent serve
bhubctl c sumbit pingpong.yml
````


## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fbenchhub%2Fbenchhub.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fbenchhub%2Fbenchhub?ref=badge_large)