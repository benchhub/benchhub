package generator

// go.go defines types for go struct and function definition

type structDef struct {
	Name   string
	Fields []fieldDef
}

type fieldDef struct {
	Name string
	Type string
}
