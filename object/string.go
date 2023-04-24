package object

type String struct {
	Value string
}

var _ Object = &String{}

func (s *String) Inspect() string {
	return s.Value
}

func (s *String) Type() ObjectType {
	return STRING_OBJ
}
