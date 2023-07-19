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

func NewGet(object Expr, name *token.Token) *Get {
	return &Get{
		Object: object,
		Name:   name,
	}
}

func (self *Get) CheckValue() (value.Value, *error.Error) {
	return self.Object.CheckValue()
}

func (self *Get) Accept(v Visitor) (value.Value, *error.Error) {
	return v.VisitGetExpr(self)
}
