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

func NewSet(object Expr, name *token.Token, value Expr) *Set {
	return &Set{
		Object: object,
		Name:   name,
		Value:  value,
	}
}

func (self *Set) Accept(v Visitor) (value.Value, *error.Error) {
	return v.VisitSetExpr(self)
}
