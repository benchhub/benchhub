# v0.0.4: tqbuilder

## TODO

- [x] list goal and non goals
- [ ] [schema generated code](#schema-generated-code)
  - [x] remove code in generated folder
  - [x] go struct in `xxxmodel`
  - [ ] struct definition for query builder `xxxschema`
  - [ ] markdown table
  - [ ] sql

## Code

- [bhgen/schema.go](../../../cmd/bhgen/schema.go) for `bhgen schema`
- [lib/tqbuilder](../../../lib/tqbuilder)

## Command

```bash
bhgen schema generate
bhgen schema clean

# or
make gen-schema
```

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

- the generated code has a package level output for `table`, `view`, `model`
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

### Schema Definition Design

```go
// core/services/user/schema/ddl/table.go
package ddl

import "github.com/benchhub/benchhub/lib/tqbuilder/sql/ddl"

func Tables() []ddl.TableDef {
	cols := []ddl.ColumnDef{
		ddl.PrimaryKey("id"),
		ddl.VarChar("name", ddl.StrSmall),
		ddl.VarChar("full_name", ddl.StrSmall),
		ddl.VarChar("email", ddl.StrSmall),
	}
	user := ddl.Table("users", cols)
	return ddl.Tables(user)
}
```

### Schema Layout

```text
benchhub
  build
    generated
       tqbuilder
         ddl
            main.go // run it to generate usermodel, userschema, userquery
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
                  usermodel
                    pkg.go
                  userschema
                  userquery
```

### Trigger Generator

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

### Schema Generated Code

- go struct that maps to table, e.g. `core/services/user/schema/generated/usermodel/user.go`
- markdown table snippet or simply `README.md`
- SQL
- table definition for query builder, similar to [jet](#jet) and cqlc
- ? orm like query e.g. getAll, getById/Name etc. could leave it to query

```go
package usermodel

type User struct {
    ID int64
    Name string
    FullName string
    Email string          
}
```

```markdown
## User

- Defined in [core/services/users/](...)
- Generated in [...](...)

TODO: forgot the markdown table syntax ...
| name | type | length | comment |
| email | varchar | 128 | |
```

```sql
CREATE TABLE users (
    id INT PRIMARY KEY,
    name VARCHAR(128),
    email VARCHAR(128)
);
```

```go
// package userschema

var User = newUserTable()

type UserTable struct {
    ID ddl.IntegerDef
    FirstName ddl.StringDef
}

func newUserTable() UserTable {
    return UserTable {
        ID: ddl.IntegerDef{Name: "id"}, // basically copy from what user wrote in their ddl/table.go
    }
}
```

### Query Definition Design

- dot import schema generated in previous phase
- dot import generic (well untyped) sql keywords

```go
// package dml

import (
    "lib/tqbuilder/sql/ddl"
    "lib/tqbuilder/sql/dml/crud"

    . "core/services/user/schema/generated/userschema"
    . "lib/tqbuilder/sql/dml"
)

func Queries() []Query {
    // write query manually
    s1 := SELECT(User.All()).FROM(User)
    q1 := NewQuery("getAllUsers", s1)

    // common crud in one go, can switch to functional argument
    crudOption := crud.Option{
        FindBy: []ddl.ColumnDef{
            User.Email, User.Twitter
        }
    }
    cruds := crud.New(User, crudOptions)
    return []Query{q1}
}
```