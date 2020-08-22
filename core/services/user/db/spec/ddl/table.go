package ddl

import (
	"github.com/benchhub/benchhub/lib/tqbuilder/sql/ast"
	. "github.com/benchhub/benchhub/lib/tqbuilder/sql/ddlbuilder"
)

func AllTables() []ast.Table {
	user := Table("users",
		Id(),
		String("name", StrSmall),
		String("email", StrSmall),
		String("description", StrMedium),
	)
	return Tables(user)
}
