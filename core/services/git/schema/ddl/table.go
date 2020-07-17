package ddl

import "github.com/benchhub/benchhub/lib/tqbuilder/sql/ddl"

func Tables() []ddl.TableDef {
	ownerCols := []ddl.ColumnDef{
		ddl.PrimaryKey("id"),
		ddl.Int("type"), // TODO: enum or a smaller int, and how to convert to proto
		ddl.Int("host_id"),
		ddl.VarChar("name", ddl.StrSmall),
	}
	owner := ddl.Table("git_owners", ownerCols)

	repoCols := []ddl.ColumnDef{
		ddl.PrimaryKey("id"),
		ddl.Int("type"), // TODO: enum or a smaller int, and how to convert to proto
		ddl.Int("owner_id"),
		ddl.VarChar("owner", ddl.StrSmall),
		ddl.VarChar("name", ddl.StrSmall),
		ddl.VarChar("goimport", ddl.StrMedium),
	}
	repo := ddl.Table("git_repos", repoCols)
	return ddl.Tables(owner, repo)
}
