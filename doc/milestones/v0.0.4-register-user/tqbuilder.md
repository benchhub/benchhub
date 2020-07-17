# v0.0.4: tqbuilder

## TODO

- [ ] list goal and non goals

## Motivation

We need a sql database library with the following features:

- type safe
  - no `sql.Query("bla")` `row.Scan(&s.a, &s.b)` like using `database/sql`
  - no struct tag and comment based generator (like k8s)
  - we always know they exact schema of query parameter and result
- go code centric
  - no more dsl, SQL is a DSL already
  - avoid parsing SQL to generate go code, write SQL AST in go code instead (query builder)
  - define schema in go code and apply to database, not the opposite
- support complex query
  - e.g. join, aggregation, group by
- support column with object value
  - e.g. json and protobuf
- protobuf integration
  - convert between proto generated struct and internal database struct
- battery included
  - built in schema migration
  - no third party dependency except sql driver
- fast
  - not orm, everything is explicit in the query, no hook or join under the hood
  - don't use reflection, use generated code when decode into struct

It is not:

- ORM
- multi dialect support (we are in the MySQL family, for now)

Obviously, implementing all of them is not practical, the plan is to only implement those blocking benchhub.

## Comparison

### sqlc

- [sqlc](https://github.com/kyleconroy/sqlc)

Highlight

- complex implementation, includes a sql parser, and has its own annotation language in comment
- support generating code from SQL query i.e. not just basic CRUD like ORM
- type safe, everything got a struct, including read queries

Why not

- don't want to include a sql parser, though using pingcap/parser might be fast
- don't want a tiny dsl for annotation

### kallax

- [kallax](https://github.com/src-d/go-kallax)

Why not

- it's more orm

### squirrel

- [squirrel](https://github.com/Masterminds/squirrel)

Why not

- plan text based query builder that is no aware of target database schema

### jet

- [jet](https://github.com/go-jet/jet)

Highlight

- the generated code has a package level out for `table`, `view`, `model`
- use dot import and looks like SQL when writing

Why not

- go code from existing schema
- feels a lot like [cqlc](https://github.com/pingginp/cqlc)

### entgo

Highlight

- define schema in go code
- write query using generated go code i.e. type safe query builder

Why not

- it is for graph query on rdbms 
  - in the future bechhub will need graph database for knowledge base and more 'connected' data

## Design

The definition and code generation contains multiple stages and each one depends on the previous one.

- schema definition
- query definition
- proto conversion

### Schema definition design

```go
// TODO: this is outdated ...
// TODO: maybe need a schema package because query has col as well
func t1() tqbuilder.Table {
    return tqbuilder.Table{
        Name: "users",
        Columns: []tqbuilder.Col{
            {
                Name: "id",
                Type: tqbuilder.Int,
                Primary: true,
            },
            {
                Name: "name",
                Type: tqbuilder.String,
                Unique: true,
            },
        }
    }
}
```

### Schema user layout

```text
benchhub
  cmd
    pmgen // imports userddl and run it
  lib
    tqbuilder
      generator
        ddl.go // generates go code from sql/ddl/ast
      sql // sql ast and builder
        ddl
          ast.go
          builder.go
        dml
  core
    services
        user
            schema
               ddl // ddl and dml in different package because dml definition relies on code generated from ddl
                  table.go
               dml
                  query.go
               generated
                  userschema
                  userquery
```

### Trigger generator

Requirements

- avoid user manually register schema for each table
- if one package is broken, code generation should still work for other packages

Scan under specific path e.g. `core/services` for pattern like `${prefix}/schema/ddl/table.go`

```text
scanFolder {
    ddls
    if hasSuffix(dir, schema/ddl/table.go) {
        addPath(ddls, dir)
    }
}

// the generated generator go file in build/generated/tqbuilder/ddl/main.go
import (
    userddl "core/services/user/schema/ddl"
    gitddl "core/services/git/schema/ddl"
)

func main() {
    ddls := []bundle{
        {
            path: "core/services/user/schema/ddl",
            tables: userddl.Tables(),
        },
        {
            path: "core/services/git/schema/ddl",
            tables: gitddl.Tables(),
        }
    }
    for _, ddl := range ddls {
       generator.GenDDL(ddl.tales, ddl.path) 
    }
}
```