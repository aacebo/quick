package expr

import (
	"quick/src/error"
	"quick/src/value"
)

type Grouping struct {
	Expr Expr
}

func (self *Grouping) Accept(v Visitor) (value.Value, *error.Error) {
	return v.VisitGroupingExpr(self)
}
