package expr

import (
	"quick/src/error"
	"quick/src/value"
)

type Literal struct {
	Type  *value.Definition
	Value value.Value
}

func NewLiteral(
	_type *value.Definition,
	value value.Value,
) *Literal {
	return &Literal{
		Type:  _type,
		Value: value,
	}
}

func (self *Literal) CheckType() (*value.Definition, *error.Error) {
	return self.Type, nil
}

func (self *Literal) Accept(v Visitor) (value.Value, *error.Error) {
	return v.VisitLiteralExpr(self)
}
