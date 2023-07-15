package expr

import (
	"quick/src/error"
	"quick/src/token"
	"quick/src/value"
)

type Logical struct {
	Left  Expr
	Op    *token.Token
	Right Expr
}

func NewLogical(left Expr, op *token.Token, right Expr) *Logical {
	return &Logical{
		Left:  left,
		Op:    op,
		Right: right,
	}
}

func (self *Logical) Accept(v Visitor) (value.Value, *error.Error) {
	return v.VisitLogicalExpr(self)
}
