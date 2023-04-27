package ast

import (
	"bytes"

	"github.com/fcidade/monkey-lang/token"
)

type IndexExpression struct {
	Token token.Token
	Left  Expression
	Index Expression
}

var _ Expression = &IndexExpression{}

func (ie *IndexExpression) expressionNode()      {}
func (ie *IndexExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IndexExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("])")

	return out.String()
}
