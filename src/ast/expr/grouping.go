package expr

import (
	"quick/src/error"
	"quick/src/value"
)

type Grouping struct {
	Expr Expr
}

func NewGrouping(expr Expr) *Grouping {
	return &Grouping{
		Expr: expr,
	}
}

func (self *Grouping) CheckValue() (value.Value, *error.Error) {
	return self.Expr.CheckValue()
}

func (self *Grouping) Accept(v Visitor) (value.Value, *error.Error) {
	return v.VisitGroupingExpr(self)
}
