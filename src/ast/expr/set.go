package expr

import (
	"quick/src/token"
	"quick/src/value"
)

type Set struct {
	Object *Expr
	Name   *token.Token
	Value  *Expr
}

func (self Set) Accept(v Visitor) value.Value {
	return v.VisitExpr(self)
}
