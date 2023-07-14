package expr

import (
	"quick/src/error"
	"quick/src/value"
)

type Literal struct {
	Value value.Value
}

func (self *Literal) Accept(v Visitor) (value.Value, *error.Error) {
	return v.VisitLiteralExpr(self)
}
