package expr

import (
	"quick/src/error"
	"quick/src/token"
	"quick/src/value"
)

type Get struct {
	Object Expr
	Name   *token.Token
}

func (self *Get) Accept(v Visitor) (value.Value, *error.Error) {
	return v.VisitGetExpr(self)
}
