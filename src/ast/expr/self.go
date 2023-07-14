package expr

import (
	"quick/src/error"
	"quick/src/token"
	"quick/src/value"
)

type Self struct {
	Keyword *token.Token
}

func (self *Self) Accept(v Visitor) (value.Value, *error.Error) {
	return v.VisitSelfExpr(self)
}
