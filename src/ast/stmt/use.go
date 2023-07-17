package stmt

import (
	"quick/src/error"
	"quick/src/token"
	"quick/src/value"
)

type Use struct {
	Name  *token.Token
	Stmts []Stmt
}

func NewUse(name *token.Token, stmts []Stmt) *Use {
	return &Use{
		Name:  name,
		Stmts: stmts,
	}
}

func (self *Use) Accept(v Visitor) (value.Value, *error.Error) {
	return v.VisitUseStmt(self)
}
