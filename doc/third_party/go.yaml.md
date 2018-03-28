# go.yaml

https://godoc.org/gopkg.in/yaml.v2#Unmarshaler

- don't call unmarshal function on itself, it will loop forever
  - use an alias struct, or another struct ...
- [ ] TODO: an example for custom unmarshaler

````go
type Unmarshaler interface {
    UnmarshalYAML(unmarshal func(interface{}) error) error
}
````

Example

````go
type OwnerType int32

const (
	OwnerType_UNKNOWN_OWNER OwnerType = 0
	OwnerType_USER          OwnerType = 1
	OwnerType_ORG           OwnerType = 2
)

func (x *OwnerType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}
	s = strings.ToLower(s)
	switch s {
	case "user":
		*x = OwnerType_USER
	case "org":
		*x = OwnerType_ORG
	}
	return nil
}
````

Implementation

````go
func (d *decoder) prepare(n *node, out reflect.Value) (newout reflect.Value, unmarshaled, good bool) {
	if n.tag == yaml_NULL_TAG || n.kind == scalarNode && n.tag == "" && (n.value == "null" || n.value == "~" || n.value == "" && n.implicit) {
		return out, false, false
	}
	again := true
	for again {
		again = false
		if out.Kind() == reflect.Ptr {
			if out.IsNil() {
				out.Set(reflect.New(out.Type().Elem()))
			}
			out = out.Elem()
			again = true
		}
		if out.CanAddr() {
			if u, ok := out.Addr().Interface().(Unmarshaler); ok {
				good = d.callUnmarshaler(n, u)
				return out, true, good
			}
		}
	}
	return out, false, false
}

func (d *decoder) callUnmarshaler(n *node, u Unmarshaler) (good bool) {
	terrlen := len(d.terrors)
	err := u.UnmarshalYAML(func(v interface{}) (err error) {
		defer handleErr(&err)
		d.unmarshal(n, reflect.ValueOf(v))
		if len(d.terrors) > terrlen {
			issues := d.terrors[terrlen:]
			d.terrors = d.terrors[:terrlen]
			return &TypeError{issues}
		}
		return nil
	})
	if e, ok := err.(*TypeError); ok {
		d.terrors = append(d.terrors, e.Errors...)
		return false
	}
	if err != nil {
		fail(err)
	}
	return true
}
````

