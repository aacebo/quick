package expr

import (
	"quick/src/error"
	"quick/src/token"
	"quick/src/value"
)

type Binary struct {
	Left  Expr
	Op    *token.Token
	Right Expr
}

func NewBinary(left Expr, op *token.Token, right Expr) *Binary {
	return &Binary{
		Left:  left,
		Op:    op,
		Right: right,
	}
}

func (self *Binary) Accept(v Visitor) (value.Value, *error.Error) {
	return v.VisitBinaryExpr(self)
}
