package ast

import "github.com/fcidade/monkey-lang/token"

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

var _ Statement = &ExpressionStatement{}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
