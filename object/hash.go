package object

import (
	"bytes"
	"fmt"
	"strings"
)

type Hash struct {
	Pairs map[HashKey]HashPair
}

var _ Object = &Hash{}

func (h *Hash) Inspect() string {
	var out bytes.Buffer
	pairs := []string{}
	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s",
			pair.Key.Inspect(), pair.Value.Inspect()))
	}
	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")
	return out.String()
}

func (e *Hash) Type() ObjectType {
	return HASH_OBJ
}
