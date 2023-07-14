package expr

import (
	"quick/src/token"
	"quick/src/value"
)

type Logical struct {
	Left  Expr
	Op    *token.Token
	Right Expr
}

func (self *Logical) Accept(v Visitor) value.Value {
	return v.VisitLogicalExpr(self)
}
