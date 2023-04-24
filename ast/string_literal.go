package ast

import "github.com/fcidade/monkey-lang/token"

type StringLiteral struct {
	Token token.Token
	Value string
}

var _ Expression = &StringLiteral{}

func (s *StringLiteral) expressionNode() {}

func (s *StringLiteral) TokenLiteral() string { return s.Token.Literal }
func (s *StringLiteral) String() string       { return s.Value }
