package stmt

import (
	"quick/src/error"
	"quick/src/token"
	"quick/src/value"
)

type Struct struct {
	Name    *token.Token
	Methods []*Fn
}

func NewStruct(name *token.Token, methods []*Fn) *Struct {
	return &Struct{
		Name:    name,
		Methods: methods,
	}
}

func (self *Struct) Accept(v Visitor) (value.Value, *error.Error) {
	return v.VisitStructStmt(self)
}
