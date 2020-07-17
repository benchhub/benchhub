package ddl

import "github.com/benchhub/benchhub/lib/tqbuilder/sql/ddl"

const (
	// TODO: not sure what should be the length of those string columns
	strLen = ddl.CharMax
)

func Tables() []ddl.TableDef {
	cols := []ddl.ColumnDef{
		ddl.PrimaryKey("id"),
		ddl.VarChar("name", strLen),
		ddl.VarChar("full_name", strLen),
		ddl.VarChar("email", strLen),
	}
	user := ddl.Table("users", cols)
	return []ddl.TableDef{user}
}