package stmt

import (
	"quick/src/ast/expr"
	"quick/src/error"
	"quick/src/value"
)

type Print struct {
	Expr expr.Expr
}

func NewPrint(expr expr.Expr) *Print {
	return &Print{
		Expr: expr,
	}
}

func (self *Print) Accept(v Visitor) (value.Value, *error.Error) {
	return v.VisitPrintStmt(self)
}
