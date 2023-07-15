package stmt

import (
	"quick/src/ast/expr"
	"quick/src/error"
	"quick/src/token"
	"quick/src/value"
)

type Return struct {
	Keyword *token.Token
	Value   expr.Expr
}

func NewReturn(keyword *token.Token, value expr.Expr) *Return {
	return &Return{
		Keyword: keyword,
		Value:   value,
	}
}

func (self *Return) Accept(v Visitor) (value.Value, *error.Error) {
	return v.VisitReturnStmt(self)
}
