package ddl

func Tables(tables ...TableDef) []TableDef {
	return tables
}

func Table(name string, cols []ColumnDef) TableDef {
	return TableDef{
		Name:    name,
		Columns: cols,
	}
}

func PrimaryKey(name string) ColumnDef {
	return ColumnDef{
		Name: name,
		Type: DataTypeDef{
			Type: TypeInt,
		},
		Constraint: ColumnConstraintDef{
			PrimaryKey: true,
		},
	}
}

func Int(name string) ColumnDef {
	return ColumnDef{
		Name: name,
		Type: DataTypeDef{
			Type: TypeInt,
		},
	}
}

func VarChar(name string, length int) ColumnDef {
	return ColumnDef{
		Name: name,
		Type: DataTypeDef{
			Type:   TypeVarchar,
			Length: length,
		},
	}
}
