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
