package expr

import (
	"quick/src/token"
	"quick/src/value"
)

type Binary struct {
	Left  *Expr
	Op    *token.Token
	Right *Expr
}

func (self Binary) Accept(v Visitor) value.Value {
	return v.VisitExpr(self)
}
