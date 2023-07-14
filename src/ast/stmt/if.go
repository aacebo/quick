package stmt

import "quick/src/ast/expr"

type If struct {
	Cond *expr.Expr
	Then *Stmt
	Else *Stmt
}

func (self If) Accept(v Visitor) {
	v.VisitStmt(self)
}
