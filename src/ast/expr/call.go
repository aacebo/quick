package expr

import (
	"quick/src/error"
	"quick/src/reflect"
	"quick/src/token"
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

func (self *Call) GetType() (reflect.Type, *error.Error) {
	t, err := self.Callee.GetType()

	if err != nil {
		return nil, err
	}

	if callable, ok := t.(reflect.CallableType); ok {
		t = callable.ReturnType()
	}

	return t, nil
}

func (self *Call) Accept(v Visitor) (*reflect.Value, *error.Error) {
	return v.VisitCallExpr(self)
}
