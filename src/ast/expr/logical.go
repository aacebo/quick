package expr

import (
	"quick/src/error"
	"quick/src/token"
	"quick/src/value"
)

type Logical struct {
	Left  Expr
	Op    *token.Token
	Right Expr
}

func (self *Logical) Accept(v Visitor) (value.Value, *error.Error) {
	return v.VisitLogicalExpr(self)
}
