package expr

import (
	"quick/src/error"
	"quick/src/token"
	"quick/src/value"
)

type Var struct {
	Name  *token.Token
	Value value.Value
}

func NewVar(name *token.Token, value value.Value) *Var {
	return &Var{
		Name:  name,
		Value: value,
	}
}

func (self *Var) CheckValue() (value.Value, *error.Error) {
	return self.Value, nil
}

func (self *Var) Accept(v Visitor) (value.Value, *error.Error) {
	return v.VisitVarExpr(self)
}
