package stmt

import (
	"quick/src/ast/expr"
	"quick/src/token"
	"quick/src/value"
)

type Var struct {
	Keyword *token.Token
	Name    *token.Token
	Init    expr.Expr
}

func (self *Var) Accept(v Visitor) value.Value {
	return v.VisitVarStmt(self)
}
