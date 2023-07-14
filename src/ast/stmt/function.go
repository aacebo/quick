package stmt

import "quick/src/token"

type Function struct {
	Name   *token.Token
	Params []*token.Token
	Body   []*Stmt
}

func (self Function) Accept(v Visitor) {
	v.VisitStmt(self)
}
