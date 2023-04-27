package object

type BuiltinFunction func(args ...Object) Object

type Builtin struct {
	Fn BuiltinFunction
}

var _ Object = &Builtin{}

func (i *Builtin) Inspect() string {
	return "builtin function"
}

func (i *Builtin) Type() ObjectType {
	return BUILTIN_OBJ
}
