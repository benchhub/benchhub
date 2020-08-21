package dml

import "github.com/benchhub/benchhub/lib/tqbuilder/sql/ddl"

// TODO(gce4-go): try to keep the plan in sync w/ tqbuilder
// TODO: unexport so we can do dot import on dml package? or use internal, or split package?

type selectStmt struct {
	Columns ddl.Column
	From    dataSource
}

type dataSource interface {
	IsDataSource()
}

type tableSource struct {
	Def ddl.TableDef
}

func (t *tableSource) IsDataSource() {
}
