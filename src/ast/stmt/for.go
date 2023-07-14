package stmt

import "quick/src/ast/expr"

type For struct {
	Init *expr.Expr
	Cond *expr.Expr
	Inc  *expr.Expr
	Body *Stmt
}

func (self For) Accept(v Visitor) {
	v.VisitStmt(self)
}
