package expr

import (
	"quick/src/error"
	"quick/src/reflect"
	"quick/src/token"
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

func (self *Set) GetType() (reflect.Type, *error.Error) {
	object, err := self.Object.GetType()

	if err != nil {
		return nil, err
	}

	value, err := self.Value.GetType()

	if err != nil {
		return nil, err
	}

	if !object.Equals(value) {
		return nil, error.New(
			self.Name.Path,
			self.Name.Ln,
			self.Name.Start,
			self.Name.End,
			"expected type '"+object.Name()+"', received '"+value.Name()+"'",
		)
	}

	return object, nil
}

func (self *Set) Accept(v Visitor) (*reflect.Value, *error.Error) {
	return v.VisitSetExpr(self)
}
