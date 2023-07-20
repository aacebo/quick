package stmt

import (
	"quick/src/ast/expr"
	"quick/src/error"
	"quick/src/reflect"
	"quick/src/token"
)

type Var struct {
	Keyword *token.Token
	Name    *token.Token
	Type    reflect.Type
	Nilable *token.Token
	Init    expr.Expr
}

func NewVar(
	keyword *token.Token,
	name *token.Token,
	_type reflect.Type,
	nilable *token.Token,
	init expr.Expr,
) *Var {
	return &Var{
		Keyword: keyword,
		Name:    name,
		Type:    _type,
		Nilable: nilable,
		Init:    init,
	}
}

func (self *Var) Accept(v Visitor) (*reflect.Value, *error.Error) {
	return v.VisitVarStmt(self)
}
