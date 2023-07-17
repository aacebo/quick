package stmt

import (
	"quick/src/error"
	"quick/src/token"
	"quick/src/value"
)

type Use struct {
	Path  []*token.Token
	Stmts []Stmt
}

func NewUse(path []*token.Token, stmts []Stmt) *Use {
	return &Use{
		Path:  path,
		Stmts: stmts,
	}
}

func (self *Use) Accept(v Visitor) (value.Value, *error.Error) {
	return v.VisitUseStmt(self)
}
