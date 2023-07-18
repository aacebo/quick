package expr

import (
	"quick/src/error"
	"quick/src/token"
	"quick/src/value"
)

type Var struct {
	Type *value.Definition
	Name *token.Token
}

func NewVar(
	_type *value.Definition,
	name *token.Token,
) *Var {
	return &Var{
		Type: _type,
		Name: name,
	}
}

func (self *Var) CheckType() (*value.Definition, *error.Error) {
	return self.Type, nil
}

func (self *Var) Accept(v Visitor) (value.Value, *error.Error) {
	return v.VisitVarExpr(self)
}
