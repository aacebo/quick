package stmt

import (
	"quick/src/ast/expr"
	"quick/src/error"
	"quick/src/reflect"
)

type Print struct {
	Expr expr.Expr
}

func NewPrint(expr expr.Expr) *Print {
	return &Print{
		Expr: expr,
	}
}

func (self *Print) Accept(v Visitor) (*reflect.Value, *error.Error) {
	return v.VisitPrintStmt(self)
}
