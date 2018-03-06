# BenchHub

[![Build Status](https://travis-ci.org/benchhub/benchhub.svg?branch=master)](https://travis-ci.org/benchhub/benchhub)
[![Go Report Card](https://goreportcard.com/badge/github.com/benchhub/benchhub)](https://goreportcard.com/report/github.com/benchhub/benchhub)

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
