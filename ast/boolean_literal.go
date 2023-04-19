package ast

import "github.com/fcidade/monkey-lang/token"

type Boolean struct {
	Token token.Token
	Value bool
}

var _ Expression = &Boolean{}

func (b *Boolean) expressionNode() {}

func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string       { return b.Token.Literal }
