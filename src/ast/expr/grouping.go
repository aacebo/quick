package expr

import (
	"quick/src/error"
	"quick/src/reflect"
)

type Grouping struct {
	Expr Expr
}

func NewGrouping(expr Expr) *Grouping {
	return &Grouping{
		Expr: expr,
	}
}

func (self *Grouping) GetType() (reflect.Type, *error.Error) {
	return self.Expr.GetType()
}

func (self *Grouping) Accept(v Visitor) (*reflect.Value, *error.Error) {
	return v.VisitGroupingExpr(self)
}
