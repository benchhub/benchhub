package bhpb

// util.go defines handy factory funcs

func NewId(id int) *IdOrName {
	return &IdOrName{
		IsId: true,
		Id:   uint64(id),
	}
}

func NewName(name string) *IdOrName {
	return &IdOrName{
		IsId: false,
		Name: name,
	}
}
