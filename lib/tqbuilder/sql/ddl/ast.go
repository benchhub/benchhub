package ddl

type DatabaseDef struct {
	Name string
}

type TableDef struct {
	Database string
	Name     string
	Columns  []ColumnDef
}

type Column interface {
	Name() string
	Type() DatabaseDef
	Constraint() ColumnConstraintDef
}

type ColumnDef struct {
	Name       string
	Type       DataTypeDef
	Constraint ColumnConstraintDef
}

type DataType string

const (
	// TODO: split int and int64? though I'd assume the code always run under env w/ int == int64
	TypeInt     DataType = "INT"
	TypeChar    DataType = "CHAR"
	TypeVarchar DataType = "VARCHAR"
)

const (
	CharMax = 1<<8 - 1
)

const (
	// TODO: figure out a reasonable default string length, might check web framework like ror, laravel etc.
	// StrSmall is for name, email etc.
	StrSmall = CharMax
	// Medium is for path, url etc.
	StrMedium = 1024
)

type DataTypeDef struct {
	Type   DataType
	Length int
}

func (t *DataTypeDef) GoType() string {
	switch t.Type {
	case TypeInt:
		return "int"
	case TypeChar, TypeVarchar:
		return "string"
	default:
		return "unknown" + string(t.Type)
	}
}

type ColumnConstraintDef struct {
	PrimaryKey bool
	Unique     bool
	// TODO: NotNull?
}
