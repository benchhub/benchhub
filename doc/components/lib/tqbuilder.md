# tqbuilder

- [Code](../../../lib/tqbuilder)

## TODO

- [ ] sync core part from milestone

## Overview

tqbuilder is a type safe SQL query builder and code generator. All the queries are known at compile time.
The overall process is the following:

- define table schema in go code
- scan for schema and generates the generator for ddl
- run the ddl generator to generates the following
  - go struct
  - dml query builder stub
  - sql
  - markdown
  - debug console/UI
- define query in go code using generated code
- scan for query and generates the generator for dml
- run the dml generator to generates the following
  - go functions that execute sql
  - sql
  - markdown
  - debug console/UI

## Milestones

- [v0.0.4 Register User](../../milestones/v0.0.4-register-user/tqbuilder.md) Init

## Layout

tqbuilder relies on a fixed directory layout for generating code and proper import.
Following is an example using BenchHub.

- specs are in `prefix/<service-name>/db/spec/ddl|dml`
- generated code are in `prefix/<service-name>/db/generated/<service-name>suffix`

```text
build
    generated
        tqbuilder
            ddl
                main.go // generated ddl generator for all the services
            dml
                main.go // generated dml generator for all the services
cmd
    pmgen // imports tqbuilder/generator to generate generators, i.e. bootstrap the code generation
core
    services
        user
            db
                spec
                    ddl
                        table.go // define table schema by importing tqbuilder/sql/ddlbuilder
                    dml
                        query.go // define table schema by importing tqbuilder/sql/dmlbuilder and generated/userschema
                generated // use multiple packages because they are generated in different stages
                    usermodel
                       pkg.go // go struct from table definition, TODO: might split into one file per table
                    userschema
                       pkg.go // query builder stub
                    userquery
                       pkg.go // high level function that calls compiled SQL query
```

## Workflow

- user write a driver to trigger the first generator, e.g. [bhgen/schema.go](../../../cmd/bhgen/schema.go)

## DDL

### Define DDL

DDL is defined using go code within a service, we use dot import on tqbuilder/sql/ddlbuilder

```go
import (
    . "github.com/benchhub/benchhub/lib/tqbuilder/sql/ddlbuilder"
)

// AllTables instead of Tables to avoid conflict with dot import
func AllTables() []ast.Table {
    cols := Columns(
        PrimaryKey("id"),
        String("name", StrSmall),
        String("email", StrMedium),
    )
    user := Table("users", cols)
    return Tables(user) // shorthand for []TableDef{user, bla}
}
```