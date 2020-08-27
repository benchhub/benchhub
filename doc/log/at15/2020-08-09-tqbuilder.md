# 2020-08-08 tqbuilder

## TODO

- [ ] stub for building query from schema should take 1-2 days
- [ ] write query in go code should take 2-3 days

## Background

See [lib/tqbuilder](../../../lib/tqbuilder). It is a SQL query builder that writes schema and query in go code.

## Updates

## 2020-08-27

- Honestly, we can write database logic w/o tqbuilder using existing tools, we can switch to tqbuilder later ...

### 2020-08-20

- flush the design into document, current doc is too outdated.
- adjust the ast struct, might need to consider changing column definition to interface in ddl, or in dml.
- consider dialect, even ddl has dialect when it comes to type mapping.
- might split package

```
// before
sql
  ddl
    ast.go
    builder.go
  dml
    ast.go
    builder.go
// after
sql
  ast
    ddl.go
    dml.go
  builder // split into two packages for dot import
    ddl
    dml
```

### 2020-08-12

I should have finished tqbuilder if the TODO list was correct.
However, like most TODO list, `cp -r this-year next-year`.

Some extra time is spent on designing the runtime and job interface.
Those work can happen in parallel and there should be not that much work in register user, project etc.

Though I might still want to think more on how the registry work (for merging multiple benchhub instances etc.)
I think the linux VFS like approach could be a good idea.

e.g. for benchhub instance A, it assumes everything is under its root.

```
/users
    at15
    gaocegege
/projects
    benchhub
        benchhub
```

for benchhub instance B, it has something similar

```
/users
    boar
    arrowrowe
/projects
    tongqu
        tongqu7
```

And we can have a benchhub instance C that mounts data from A and B into one.
There can be confilict, and it depends on conflict resolution (kind of like overlayfs)

```
/users
    at15
    arrowrowe
    boar
    gaocegege
/projects
    benchhub
        benchhub
    tongqu
        tongqu7
```

Or it we want to have a longer URL, mount each under different paths

```
/users
    /gh
        at15
    /dy
        boar
/projects
    /gh
        benchhub
            benchhub
    /dy
        tongqu
            tongqu7        
```