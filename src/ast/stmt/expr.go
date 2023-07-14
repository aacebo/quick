package stmt

import "quick/src/ast/expr"

type Expr struct {
	Expr *expr.Expr
}

func (self Expr) Accept(v Visitor) {
	v.VisitStmt(self)
}
