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
	t, err := self.Object.GetType()

	if err != nil {
		return nil, err
	}

	return t.GetMember(self.Name.String()), nil
}

func (self *Get) Accept(v Visitor) (*reflect.Value, *error.Error) {
	return v.VisitGetExpr(self)
}
