package stmt

import "quick/src/value"

type Block struct {
	Stmts []Stmt
}

func (self *Block) Accept(v Visitor) value.Value {
	return v.VisitBlockStmt(self)
}
