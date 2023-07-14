package expr

import (
	"quick/src/error"
	"quick/src/token"
	"quick/src/value"
)

type Set struct {
	Object Expr
	Name   *token.Token
	Value  Expr
}

func (self *Set) Accept(v Visitor) (value.Value, *error.Error) {
	return v.VisitSetExpr(self)
}
