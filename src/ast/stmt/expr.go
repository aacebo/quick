package stmt

import (
	"quick/src/ast/expr"
	"quick/src/error"
	"quick/src/value"
)

type Expr struct {
	Expr expr.Expr
}

func NewExpr(expr expr.Expr) *Expr {
	return &Expr{
		Expr: expr,
	}
}

func (self *Expr) Accept(v Visitor) (value.Value, *error.Error) {
	return v.VisitExprStmt(self)
}
