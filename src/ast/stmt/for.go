package stmt

import (
	"quick/src/ast/expr"
	"quick/src/error"
	"quick/src/value"
)

type For struct {
	Init Stmt
	Cond expr.Expr
	Inc  expr.Expr
	Body Stmt
}

func NewFor(init Stmt, cond expr.Expr, inc expr.Expr, body Stmt) *For {
	return &For{
		Init: init,
		Cond: cond,
		Inc:  inc,
		Body: body,
	}
}

func (self *For) Accept(v Visitor) (value.Value, *error.Error) {
	return v.VisitForStmt(self)
}
