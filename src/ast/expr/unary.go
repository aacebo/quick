package expr

import (
	"quick/src/error"
	"quick/src/token"
	"quick/src/value"
)

type Unary struct {
	Op    *token.Token
	Right Expr
}

func (self *Unary) Accept(v Visitor) (value.Value, *error.Error) {
	return v.VisitUnaryExpr(self)
}
