package expr

import (
	"quick/src/token"
	"quick/src/value"
)

type Call struct {
	Callee Expr
	Paren  *token.Token
	Args   []Expr
}

func (self *Call) Accept(v Visitor) value.Value {
	return v.VisitCallExpr(self)
}
