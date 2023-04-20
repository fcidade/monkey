package object

import "fmt"

type Boolean struct {
	Value bool
}

var _ Object = &Boolean{}

func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

func (b *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}
