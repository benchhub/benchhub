package generator

// go.go defines types for go struct and function definition

type structDef struct {
	Table  string
	Name   string
	Fields []fieldDef
}

type fieldDef struct {
	Column string
	Name   string
	Type   string
}
