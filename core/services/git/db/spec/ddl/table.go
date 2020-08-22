package ddl

import (
	"github.com/benchhub/benchhub/lib/tqbuilder/sql/ast"
	. "github.com/benchhub/benchhub/lib/tqbuilder/sql/ddlbuilder"
)

func AllTables() []ast.Table {
	owner := Table("owners",
		Id(),
		Int("type"), // TODO: smaller int, and how to convert to/from proto automatically
		Int("host_id"),
		String("name", StrSmall), // TODO: unique? btw: unique constraint also becomes a secondary index?
	)
	repo := Table("git_repos",
		Id(),
		Int("type"), // TODO: maybe have a Enum method ... and let the dialect handle it
		String("owner", StrSmall),
		String("name", StrSmall),
		String("goimport", StrMedium),
	)
	return Tables(owner, repo)
}
