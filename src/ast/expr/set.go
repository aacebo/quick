package expr

import (
	"quick/src/error"
	"quick/src/token"
	"quick/src/value"
)

type Set struct {
	Object Expr
	Name   *token.Token
	Value  Expr
}

func NewSet(object Expr, name *token.Token, value Expr) *Set {
	return &Set{
		Object: object,
		Name:   name,
		Value:  value,
	}
}

func (self *Set) CheckValue() (value.Value, *error.Error) {
	object, err := self.Object.CheckValue()

	if err != nil {
		return nil, err
	}

	value, err := self.Value.CheckValue()

	if err != nil {
		return nil, err
	}

	if !object.TypeEq(value) {
		return nil, error.New(
			self.Name.Path,
			self.Name.Ln,
			self.Name.Start,
			self.Name.End,
			"type '"+object.Name()+"' is not '"+value.Name()+"'",
		)
	}

	return object, nil
}

func (self *Set) Accept(v Visitor) (value.Value, *error.Error) {
	return v.VisitSetExpr(self)
}
