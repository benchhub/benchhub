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
	Constraint() Constraint
	isColumn()
}

// TODO(gce4-go): other table constraint, not null?
type Constraint struct {
	PrimaryKey bool
	Unique     bool
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

type IntColumn struct {
	baseColumn
}

type BoolColumn struct {
	baseColumn
}

type StringColumn struct {
	baseColumn
	Cap int // Cap is max length in VARCHAR TODO: are we going to support fixed length string like CHAR
}
