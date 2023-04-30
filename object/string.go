package object

import "hash/fnv"

type String struct {
	Value string
}

var _ Object = &String{}
var _ Hashable = &String{}

func (s *String) Inspect() string {
	return s.Value
}

func (s *String) Type() ObjectType {
	return STRING_OBJ
}

func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))
	return HashKey{
		Type:  s.Type(),
		Value: h.Sum64(),
	}
}
