package expr

import (
	"quick/src/token"
	"quick/src/value"
)

type Unary struct {
	Op    *token.Token
	Right Expr
}

func (self *Unary) Accept(v Visitor) value.Value {
	return v.VisitUnaryExpr(self)
}
