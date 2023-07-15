package expr

import (
	"quick/src/error"
	"quick/src/token"
	"quick/src/value"
)

type Call struct {
	Callee Expr
	Paren  *token.Token
	Args   []Expr
}

func NewCall(callee Expr, paren *token.Token, args []Expr) *Call {
	return &Call{
		Callee: callee,
		Paren:  paren,
		Args:   args,
	}
}

func (self *Call) Accept(v Visitor) (value.Value, *error.Error) {
	return v.VisitCallExpr(self)
}
