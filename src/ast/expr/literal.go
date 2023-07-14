package expr

import "quick/src/value"

type Literal struct {
	Value value.Value
}

func (self Literal) Accept(v Visitor) value.Value {
	return v.VisitExpr(self)
}
