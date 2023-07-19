package expr

import (
	"quick/src/error"
	"quick/src/value"
)

type Literal struct {
	Value value.Value
}

func NewLiteral(value value.Value) *Literal {
	return &Literal{
		Value: value,
	}
}

func (self *Literal) CheckValue() (value.Value, *error.Error) {
	return self.Value, nil
}

func (self *Literal) Accept(v Visitor) (value.Value, *error.Error) {
	return v.VisitLiteralExpr(self)
}
