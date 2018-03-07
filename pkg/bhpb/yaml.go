package bhpb

import (
	"log"
	"encoding/json"
)

//func (x *OwnerType) UnmarshalYAML(unmarshal func(interface{}) error) error {
//	log.Printf("unmarshaler called")
//	var s string
//	if err := unmarshal(&s); err != nil {
//		return err
//	}
//	switch s {
//	case "user":
//		*x = OwnerType_USER
//	case "org":
//		*x = OwnerType_ORG
//	}
//	return nil
//}

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

func (x *OwnerType) UnmarshalJSON(b []byte) error {
	log.Printf("unmarshaler called")
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
