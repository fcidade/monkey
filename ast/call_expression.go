package ast

import (
	"bytes"
	"strings"

	"github.com/fcidade/monkey-lang/token"
)

type CallExpression struct {
	Token     token.Token
	Function  Expression
	Arguments []Expression
}

var _ Expression = &CallExpression{}

func (ce *CallExpression) expressionNode() {}

func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CallExpression) String() string {
	var out bytes.Buffer

	arguments := []string{}
	for _, param := range ce.Arguments {
		arguments = append(arguments, param.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(arguments, ", "))
	out.WriteString(")")

	return out.String()
}
