package stmt

type Visitor interface {
	VisitStmt(Stmt)
}
