package stmt

import (
	"quick/src/error"
	"quick/src/reflect"
	"quick/src/token"
)

type Fn struct {
	Name       *token.Token
	Params     []*Var
	ReturnType reflect.Type
	Body       []Stmt
}

func NewFn(
	name *token.Token,
	params []*Var,
	returnType reflect.Type,
	body []Stmt,
) *Fn {
	return &Fn{
		Name:       name,
		Params:     params,
		ReturnType: returnType,
		Body:       body,
	}
}

func (self *Fn) Accept(v Visitor) (*reflect.Value, *error.Error) {
	return v.VisitFnStmt(self)
}
