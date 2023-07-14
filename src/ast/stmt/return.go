package stmt

import (
	"quick/src/ast/expr"
	"quick/src/token"
)

type Return struct {
	Keyword *token.Token
	Value   *expr.Expr
}

func (self Return) Accept(v Visitor) {
	v.VisitStmt(self)
}
