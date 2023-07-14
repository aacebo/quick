package expr

import (
	"quick/src/error"
	"quick/src/token"
	"quick/src/value"
)

type Var struct {
	Name *token.Token
}

func (self *Var) Accept(v Visitor) (value.Value, *error.Error) {
	return v.VisitVarExpr(self)
}
