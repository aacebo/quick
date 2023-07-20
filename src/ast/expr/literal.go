package expr

import (
	"quick/src/error"
	"quick/src/reflect"
)

type Literal struct {
	Value *reflect.Value
}

func NewLiteral(value *reflect.Value) *Literal {
	return &Literal{
		Value: value,
	}
}

func (self *Literal) GetType() (reflect.Type, *error.Error) {
	return self.Value.Type(), nil
}

func (self *Literal) Accept(v Visitor) (*reflect.Value, *error.Error) {
	return v.VisitLiteralExpr(self)
}
