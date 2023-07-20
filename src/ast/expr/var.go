package expr

import (
	"quick/src/error"
	"quick/src/reflect"
	"quick/src/token"
)

type Var struct {
	Name *token.Token
	Type reflect.Type
}

func NewVar(name *token.Token, _type reflect.Type) *Var {
	return &Var{
		Name: name,
		Type: _type,
	}
}

func (self *Var) GetType() (reflect.Type, *error.Error) {
	return self.Type, nil
}

func (self *Var) Accept(v Visitor) (*reflect.Value, *error.Error) {
	return v.VisitVarExpr(self)
}
