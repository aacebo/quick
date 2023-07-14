package stmt

import (
	"quick/src/ast/expr"
	"quick/src/value"
)

type For struct {
	Init expr.Expr
	Cond expr.Expr
	Inc  expr.Expr
	Body Stmt
}

func (self *For) Accept(v Visitor) value.Value {
	return v.VisitForStmt(self)
}
