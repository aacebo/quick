package stmt

type Print struct {
	Expr *Expr
}

func (self Print) Accept(v Visitor) {
	v.VisitStmt(self)
}
