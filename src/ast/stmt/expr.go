package stmt

import (
	"quick/src/ast/expr"
	"quick/src/value"
)

type Expr struct {
	Expr expr.Expr
}

func (self *Expr) Accept(v Visitor) value.Value {
	return v.VisitExprStmt(self)
}
