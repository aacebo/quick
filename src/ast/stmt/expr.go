package stmt

import (
	"quick/src/ast/expr"
	"quick/src/error"
	"quick/src/value"
)

type Expr struct {
	Expr expr.Expr
}

func (self *Expr) Accept(v Visitor) (value.Value, *error.Error) {
	return v.VisitExprStmt(self)
}
