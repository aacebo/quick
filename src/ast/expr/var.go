package expr

import (
	"quick/src/token"
	"quick/src/value"
)

type Var struct {
	Name *token.Token
}

func (self *Var) Accept(v Visitor) value.Value {
	return v.VisitVarExpr(self)
}
