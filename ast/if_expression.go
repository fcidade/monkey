package ast

import (
	"bytes"

	"github.com/fcidade/monkey-lang/token"
)

type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

var _ Expression = &IfExpression{}

func (ie *IfExpression) expressionNode() {}

func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString(ie.TokenLiteral())
	out.WriteString(" ")
	out.WriteString(ie.Condition.String())
	out.WriteString(" { ")
	out.WriteString(ie.Consequence.String())
	out.WriteString(" }")

	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}
