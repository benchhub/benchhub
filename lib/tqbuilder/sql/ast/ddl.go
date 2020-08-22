package ast

type Table struct {
	Database string
	Name     string
	Columns  []Column
}

func (t *Table) isDataSource() {
}

// TODO: index?
type Column interface {
	Name() string
	GoType() string // TODO: rename to ValueGoType()? we also need go type of column implementation in generator
	Constraint() Constraint
	isColumn()
}

// TODO(gce4-go): other table constraint?
type Constraint struct {
	PrimaryKey bool
	Unique     bool
	Null       bool // Null instead of NotNull because most time we want NOT NULL on all columns.
}

type baseColumn struct {
	name       string
	constraint Constraint
}

func (b *baseColumn) isColumn() {
}

func (b *baseColumn) Name() string {
	return b.name
}

func (b *baseColumn) SetName(s string) {
	b.name = s
}

func (b *baseColumn) Constraint() Constraint {
	return b.constraint
}

func (b *baseColumn) SetConstraint(c Constraint) {
	b.constraint = c
}

type BoolColumn struct {
	baseColumn
}

func (*BoolColumn) GoType() string {
	return "bool"
}

type IntColumn struct {
	baseColumn
}

// TODO: maybe int64?
func (*IntColumn) GoType() string {
	return "int"
}

type StringColumn struct {
	baseColumn
	Cap int // Cap is max length in VARCHAR TODO: are we going to support fixed length string type like CHAR
}

func (*StringColumn) GoType() string {
	return "string"
}
