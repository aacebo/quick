package expr

import (
	"quick/src/error"
	"quick/src/reflect"
	"quick/src/token"
)

type Assign struct {
	Name  *token.Token
	Value Expr
}

func NewAssign(name *token.Token, value Expr) *Assign {
	return &Assign{
		Name:  name,
		Value: value,
	}
}

func (self *Assign) GetType() (reflect.Type, *error.Error) {
	return self.Value.GetType()
}

func (self *Assign) Accept(v Visitor) (*reflect.Value, *error.Error) {
	return v.VisitAssignExpr(self)
}
