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

func NewUnary(op *token.Token, right Expr) *Unary {
	return &Unary{
		Op:    op,
		Right: right,
	}
}

func (self *Unary) CheckType() (*value.Definition, *error.Error) {
	return self.Right.CheckType()
}

func (self *Unary) Accept(v Visitor) (value.Value, *error.Error) {
	return v.VisitUnaryExpr(self)
}
