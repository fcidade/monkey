package object

type Error struct {
	Message string
}

var _ Object = &Error{}

func (e *Error) Inspect() string {
	return "Error: " + e.Message
}

func (e *Error) Type() ObjectType {
	return ERROR_OBJ
}
