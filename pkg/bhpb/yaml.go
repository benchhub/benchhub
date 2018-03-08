package bhpb

import (
	"encoding/json"
)

// custom YAML unmarshaler to deal with enum type in proto, we use string in config, but unmarshal to int to match the enum

func (x *OwnerType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "user":
		*x = OwnerType_USER
	case "org":
		*x = OwnerType_ORG
	}
	return nil
}

//func (m *Owner) UnmarshalYAML(unmarshal func(interface{}) error) error {
//	log.Printf("unmarshaler called")
//	var aux struct {
//		Id   string
//		Name string
//		Type string
//	}
//	if err := unmarshal(&aux); err != nil {
//		return err
//	}
//	log.Printf("type %s", aux.Type)
//	m.Id = aux.Id
//	m.Name = aux.Name
//	switch aux.Type {
//	case "user":
//		m.Type = OwnerType_USER
//	case "org":
//		m.Type = OwnerType_ORG
//	}
//	return nil
//}

//func (m *Owner) UnmarshalJSON(b []byte) error {
//	var aux struct {
//		Id   string
//		Name string
//		Type string
//	}
//	if err := json.Unmarshal(b, &aux); err != nil {
//		return err
//	}
//	log.Printf("type %s", aux.Type)
//	m.Id = aux.Id
//	m.Name = aux.Name
//	switch aux.Type {
//	case "user":
//		m.Type = OwnerType_USER
//	case "org":
//		m.Type = OwnerType_ORG
//	}
//	return nil
//}

// Deprecated
// used by github.com/ghodss/yaml because it it convert yaml to json
func (x *OwnerType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "user":
		*x = OwnerType_USER
	case "org":
		*x = OwnerType_ORG
	}
	return nil
}

func (x *Role) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "any":
		*x = Role_ANY
	case "central":
		*x = Role_CENTRAL
	case "loader":
		*x = Role_LOADER
	case "database":
		*x = Role_DATABASE
	}
	return nil
}

func (x *TaskDriver) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "stopper":
		*x = TaskDriver_STOPPER
	case "shell":
		*x = TaskDriver_SHELL
	case "exec":
		*x = TaskDriver_EXEC
	case "docker":
		*x = TaskDriver_DOCKER
	}
	return nil
}
