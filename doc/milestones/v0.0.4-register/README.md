# v0.0.4 Register

## TODO

- [ ] forgot example
- [ ] still need user even if we don't have authentication
- [ ] http api? we can mount the logic in both http and grpc server, this should make ui life easier
  - or we could use grpc gateway etc.

Done

- [x] rdbms framework name `tqbuilder`, type safe query builder

## Motivation

Split from [v0.1.0](../v0.1.0-micro). v0.0.4 covers registering benchmark(s) from vcs to benchhub's database.

## Specs

Allow

- register a git repository e.g. `benchhub/benchhub`
- register a folder within a git repository, e.g. `benchhub/benchhub/_example/mergesort`
- move folder and keeps old data, e.g. `benchhub/benchhub/_example/mergesort` -> `benchhub/example/mergesort`

Skipped

- authentication, assuming a `bh` client can create record for anything w/o user credential

## Features

Order by implementation order.

- [Register User](#register-user)
- Git i.e. VCS
  - [Register Git Host](#register-git-host)
  - [Register Owner](#register-owner)
  - [Register Repository](#register-repository)
- [Register Project](#register-project)
- [Register Test](#register-test)
- [Register Benchmark](#register-benchmark)

### Register User

Description

Although we tide user, repo logic with vcs. It's common for a user to have account on multiple vcs or multiple accounts in same vcs.
Further more, the user running test/benchmark may not be the owner of the repo.

Components

- `bhpb`
  - define `User`
  - rpc `GetUser`, `ListUser`, `RegisterUser`
- `lib/tqbuilder`
  - define schema in go code
  - generate basic query like normal orm
- `core/storage/rdbms`
  - save user in rdbms
- `core/server`
  - init grpc server
- `cmd/bh`
  - `bh register user at15`
- `ui`
  - `users` show registered users
  - NO plan for implementing login, oauth etc. (for now)

### Register Git Host

Description

We can use `github.com` but there is gitlab, coding.net, private github/gitlab instances.
And all the cloud service providers have their own VCS solutions.

To make things easier and consistent, we hard code popular vcs providers in code and create them on startup.

- `0` reserved
- `1-7` common public providers (that I use)
  - `1` github.com
  - `2` gitlab.com
  - `3` bitbucket.org
  - `4` dyweb
  - `5` coding.net
  - `6` gitcafe.com (sold to coding.net)
- `8+` other providers

Providers can have the following type, this is hard coded in proto and can be used when we offer CI/CD integration.

- github (enterprise)
- gitlab (self deployed instances)
- other

Components

- `bhpb`
  - proto definition for `GitHost`
  - rpc `GetGitHost`, `ListGitHost`
- `core/storage/rdbms`
  - save git host in database, maybe memory as well?
- `core/server`
  - implement rpc
- `ui`
  - `githost` default to github
  
### Register Owner

Description

Register a owner (user/organization). It does not need to be associated with a benchhub user.
e.g. even if based on github API current user is does not own specific organization, they can still create that git owner.

Components

- `bhpb`
  - proto definition `GitOwner`
- `cmd/bh`
  - `bh register owner at15` or `bh register owner benchhub`
  - `bh list owner` or `bh owner list --host=gitlab` default host is github

### Register Repository

Description

Register a git repo `<git-host>/<owner>/<project>` as a project in benchhub.
Requires registering `<git-host>` and `<owner>` before registering project.

Components

- `bhpb`
  - proto definition for `GitRepo`, `RepoConfig`
  - rpc `GetRepo`, `ListRepo`, `UpdateRepo`, `DeleteRepo`
- `core/config`
  - specify the repo in config
  - use yaml if rcl is not ready
- `core/storage/rdbms`
  - save git host, owner, repo in three different tables
- `core/server`
  - return registered repos
- `cmd/bh`
  - `bh register repo gommon` or `bh repo register gommon`
  - `bh list repo` or `bh repo list`
- `ui`
  - `git-host>owner>repo` allow click down the menu
- `test`
  - framework to run the entire stack locally (using docker for database?)

Example

```rcl
// github.com/benchhub/benchhub/bh.rcl

type: repo,

repo: {
    host: github.com,
    owner: benchhub,
    project: benchhub,
    owners: [
        at15,
        gaocegege,
        arrowrowe,
    ],
    // TODO: path to projects?
}
```

### Register Project

Description

A repository can contain multiple projects. There are two types of projects:

- meta projects, group multiple projects together e.g. mergesort in go, cpp, rust etc.
- executable project, run the actual benchmark using a specific framework e.g. go, tpc-c

Components

- `core/config`
  - give a name for the project e.g. `mergesort` so it can be referenced using `benchhub/benchhub/mergesort`
    - default can be folder name, but rename is useful for things like `mergesort/c` to `mergesort-c`
  - title and description etc.
- `core/storage/rdbms`
  - save project
- `cmd/bh`
  - `bh register project`
- `ui`
  - `git-host>owner>repo>projects`

Example

```rcl
// github.com/benchhub/benchhub/_example/mergesort/bh.rcl

type: project,

project: {
    name: mergesort,
    meta: true,
    memembers: [ // path to subprojects
        "c",
        "go",
        "cpp",
    ]
}

// github.com/benchhub/benchhub/_example/mergesort/cpp/bh.rcl

type: project,

project: {
    name: mergesort-cpp,
    lang: cpp,
    test: lang/cpp/catch2,
    bench: lang/googlebenchmark,
}

// github.com/benchhub/benchhub/_example/mergesort/go/bh.rcl

type: project,

project: {
    name: mergesort-go,
    lang: go,
    test: lang/go/gotest,
    bench: lang/go/gobench,
}
```

### Register Test

Description

It's impractical to have people manually enter all the test/benchmark names.
The easiest solution is to run the test and parse the output.
The name of the test should be unique inside the project and relies on the framework.

Known issues

- conditional test won't work i.e. `go test -short`, this is a special case of parameter. We will consider that when we start actually running test/benchmark with parameters.

Components

- `core/config`
  - allow ignoring or force specific test patterns, it is opaque to bh and depends on framework
  - the command to run test, the command need to save the output in a place framework can find it (or stream to stdout/err)
- `core/storage/rdbms`
  - save test, saving all tests in a single table w/o partition may cause slow query in the end, but for now `id, project_id, test_name` should be good enough
- `cmd/bh`
  - `bh register test`
- `ui`
  - `host>owner>repo>project>bench`
  
### Register Benchmark

Same as [test](#register-test).

## Components

### bh

Description

The client binary. Also include server in `bh srv`.

### Core Config

Description

Lookup and parser config. TODO(bhpb): define config struct in proto as well?

Used in

- [repo](#register-repository)
- [project](#register-project)
- [test](#register-test)
- [benchmark](#register-benchmark)

Internal Dependencies

- NULL

External Dependencies

- `gommon`
  - `dcli`
- `reikalang`
  - `rcl/rcl-go`

TODO

- [ ] struct definition
- [ ] traverse dir for finding parent config, e.g. get repo for a project

### Core Storage RDBMS

Description

Connector for RDMBS (MySQL for now).

Used in

- [repo](#register-repository)
- [project](#register-project)
- [test](#register-test)
- [benchmark](#register-benchmark)

Internal Dependencies

- a query builder like `entgo`, avoid using raw sql, could consider using ORM for early prototype

TODO

- [ ] database migration, at least do it automatically when running tests
- [ ] CRUD for repo, project, test, benchmark

### Core Server

Description

gRPC server and connects to a RDBMS (for now).

Used in

- [repo](#register-repository)
- [project](#register-project)
- [test](#register-test)
- [benchmark](#register-benchmark)

Internal Dependencies

- `core/storage`

TODO

- [ ] service definition
   - can use a `Result<Bla>` in response and `Request<Bla>` for common error message and request meta.
- [ ] talks to RDBMS

### Framework gotest

Description

Parse `go test` output

Used in

- [test](#register-test)

TODO

- [ ] there should be json output format for go test?
  - can try to use proto ...

### Framework gobench

Description

Parse `go test -bench` output

Used in

- [benchmark](#register-benchmark)

TODO

- [ ] parse go bench output? though it might be same as go test if we just need test name?

### Framework rusttest

Description

Get result from `cargo test`

Used in

- [test](#register-test)

TODO

- [ ] `cargo test` output format?