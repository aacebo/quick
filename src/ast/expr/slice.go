package expr

import (
	"quick/src/error"
	"quick/src/reflect"
)

type Slice struct {
	Type  reflect.Type
	Items []Expr
}

func NewSlice(_type reflect.Type, items []Expr) *Slice {
	return &Slice{
		Type:  reflect.NewSliceType(_type, -1),
		Items: items,
	}
}

func (self *Slice) GetType() (reflect.Type, *error.Error) {
	return self.Type, nil
}

func (self *Slice) Accept(v Visitor) (*reflect.Value, *error.Error) {
	return v.VisitSliceExpr(self)
}
