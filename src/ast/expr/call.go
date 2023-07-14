package expr

import (
	"quick/src/error"
	"quick/src/token"
	"quick/src/value"
)

type Call struct {
	Callee Expr
	Paren  *token.Token
	Args   []Expr
}

func (self *Call) Accept(v Visitor) (value.Value, *error.Error) {
	return v.VisitCallExpr(self)
}
