package expr

import "quick/src/value"

type Grouping struct {
	Expr *Expr
}

func (self Grouping) Accept(v Visitor) value.Value {
	return v.VisitExpr(self)
}
