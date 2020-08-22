package ddlbuilder

import "github.com/benchhub/benchhub/lib/tqbuilder/sql/ast"

func Tables(tables ...ast.Table) []ast.Table {
	return tables
}

func Table(name string, cols ...ast.Column) ast.Table {
	return ast.Table{
		Database: "",
		Name:     name,
		Columns:  cols,
	}
}

// ----------------------------------------------------------------------------
// Columns of different data type

func Columns(cols ...ast.Column) []ast.Column {
	return cols
}

// TODO: allow setting constraints using variadic args
func Int(name string) *ast.IntColumn {
	s := ast.IntColumn{}
	s.SetName(name)
	return &s
}

const (
	// CharMax is max size of a char column, i.e. 255
	CharMax = 1<<8 - 1
	// TODO: figure out a reasonable default string length, might check web framework like ror, laravel etc.
	// StrSmall stores short string like name, email etc.
	StrSmall = CharMax
	// Medium stores longer string like file path, url etc.
	StrMedium = 1024
)

func String(name string, capacity int) *ast.StringColumn {
	s := ast.StringColumn{Cap: capacity}
	s.SetName(name)
	return &s
}

// ----------------------------------------------------------------------------
// Common Columns

func PrimaryKey(name string) *ast.IntColumn {
	s := Int(name)
	s.SetConstraint(ast.Constraint{
		PrimaryKey: true,
	})
	return s
}

// Id returns a PrimaryKey column with name `id`
func Id() *ast.IntColumn {
	return PrimaryKey("id")
}
