package ast

import "github.com/fcidade/monkey-lang/token"

type Identifier struct {
	Token token.Token
	Value string
}

var _ Expression = &Identifier{}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }
