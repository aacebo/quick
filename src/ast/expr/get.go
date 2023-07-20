package expr

import (
	"quick/src/error"
	"quick/src/reflect"
	"quick/src/token"
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

func (self *Get) GetType() (reflect.Type, *error.Error) {
	return self.Object.GetType()
}

func (self *Get) Accept(v Visitor) (*reflect.Value, *error.Error) {
	return v.VisitGetExpr(self)
}
