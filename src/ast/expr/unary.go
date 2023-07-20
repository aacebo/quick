package expr

import (
	"quick/src/error"
	"quick/src/reflect"
	"quick/src/token"
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

func (self *Unary) GetType() (reflect.Type, *error.Error) {
	return self.Right.GetType()
}

func (self *Unary) Accept(v Visitor) (*reflect.Value, *error.Error) {
	return v.VisitUnaryExpr(self)
}
