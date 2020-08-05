package generator

// go.go defines types for go struct and function definition
// TODO(dyweb/gommon/generator): put struct definition there

type structDef struct {
	Name   string
	Fields []fieldDef
}

type fieldDef struct {
	Name string
	Type string
}
