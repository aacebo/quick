package stmt

import (
	"quick/src/ast/expr"
	"quick/src/error"
	"quick/src/reflect"
)

type If struct {
	Cond expr.Expr
	Then Stmt
	Else Stmt
}

func NewIf(cond expr.Expr, then Stmt, _else Stmt) *If {
	return &If{
		Cond: cond,
		Then: then,
		Else: _else,
	}
}

func (self *If) Accept(v Visitor) (*reflect.Value, *error.Error) {
	return v.VisitIfStmt(self)
}
