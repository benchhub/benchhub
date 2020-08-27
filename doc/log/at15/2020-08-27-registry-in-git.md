# 2020-08-27 Registry in Git

## Background

While working on dyweb/pm, user saves document in a directory tracked by git.
Similar approach can be taken for benchhub registry.

This provides three benefits

- file system enforces name uniqueness
- we no longer need to dump database to a filesystem
- shell becomes a UI for exploring the registry

The registry looks like following

```
/users
  /at15
  /gaocegege
/organizations
  benchhub/
    benchhub
/git
  github.com/
    benchhub/
      benchhub
    dyweb/
      gommon
/projects
  benchhub/
    benchhub
```

It has the following drawbacks

- the git repo is almost not viewable on github, e.g. homebrew, rust crate
- the sync logic is more complex, the database schema builds from a specific snapshot (git commit)

There are also extra things we need to consider

- how can people test a project without adding themself to the registry
- efficient sync w/o looping through entire tree when system restarts/upgrade