package stmt

import (
	"quick/src/ast/expr"
	"quick/src/error"
	"quick/src/value"
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

func (self *If) Accept(v Visitor) (value.Value, *error.Error) {
	return v.VisitIfStmt(self)
}
