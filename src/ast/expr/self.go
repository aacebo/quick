package expr

import (
	"quick/src/token"
	"quick/src/value"
)

type Self struct {
	Keyword *token.Token
}

func (self Self) Accept(v Visitor) value.Value {
	return v.VisitExpr(self)
}
