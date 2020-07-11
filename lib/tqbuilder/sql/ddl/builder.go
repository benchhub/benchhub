package ddl

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

func VarChar(name string, length int) ColumnDef {
	return ColumnDef{
		Name: name,
		Type: DataTypeDef{
			Type:   TypeVarchar,
			Length: length,
		},
	}
}
