package expr

import (
	"quick/src/error"
	"quick/src/token"
	"quick/src/value"
)

type Assign struct {
	Name  *token.Token
	Value Expr
}

func NewAssign(name *token.Token, value Expr) *Assign {
	return &Assign{
		Name:  name,
		Value: value,
	}
}

func (self *Assign) CheckType() (*value.Definition, *error.Error) {
	return self.Value.CheckType()
}

func (self *Assign) Accept(v Visitor) (value.Value, *error.Error) {
	return v.VisitAssignExpr(self)
}
