package expr

import (
	"quick/src/token"
	"quick/src/value"
)

type Assign struct {
	Name  *token.Token
	Value *Expr
}

func (self Assign) Accept(v Visitor) value.Value {
	return v.VisitExpr(self)
}
