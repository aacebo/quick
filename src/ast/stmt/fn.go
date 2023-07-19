package stmt

import (
	"quick/src/error"
	"quick/src/token"
	"quick/src/value"
)

type Fn struct {
	Name       *token.Token
	Params     []*Var
	ReturnType value.Value
	Body       []Stmt
}

func NewFn(
	name *token.Token,
	params []*Var,
	returnType value.Value,
	body []Stmt,
) *Fn {
	return &Fn{
		Name:       name,
		Params:     params,
		ReturnType: returnType,
		Body:       body,
	}
}

func (self *Fn) Accept(v Visitor) (value.Value, *error.Error) {
	return v.VisitFnStmt(self)
}
