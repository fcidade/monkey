package ast

import (
	"bytes"
	"strings"

	"github.com/fcidade/monkey-lang/token"
)

type FunctionLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

var _ Expression = &FunctionLiteral{}

func (fl *FunctionLiteral) expressionNode()      {}
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }

func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	out.WriteString(fl.TokenLiteral())
	out.WriteString(" ")

	params := []string{}
	for _, param := range fl.Parameters {
		params = append(params, param.String())
	}
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")

	out.WriteString("{ ")
	out.WriteString(fl.Body.String())
	out.WriteString(" }")

	return out.String()
}
