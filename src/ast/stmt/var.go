package stmt

import (
	"quick/src/ast/expr"
	"quick/src/token"
)

type Var struct {
	Keyword *token.Token
	Name    *token.Token
	Init    *expr.Expr
}

func (self Var) Accept(v Visitor) {
	v.VisitStmt(self)
}
