package stmt

type Block struct {
	Stmts []*Stmt
}

func (self Block) Accept(v Visitor) {
	v.VisitStmt(self)
}
