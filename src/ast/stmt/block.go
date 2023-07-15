package stmt

import (
	"quick/src/error"
	"quick/src/value"
)

type Block struct {
	Stmts []Stmt
}

func NewBlock(stmts []Stmt) *Block {
	return &Block{
		Stmts: stmts,
	}
}

func (self *Block) Accept(v Visitor) (value.Value, *error.Error) {
	return v.VisitBlockStmt(self)
}
