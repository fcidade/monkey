package object

type ReturnValue struct {
	Value Object
}

var _ Object = &ReturnValue{}

func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}

func (rv *ReturnValue) Type() ObjectType {
	return RETURN_VALUE_OBJ
}
