package stmt

import "quick/src/token"

type Struct struct {
	Name    *token.Token
	Methods []*Function
}

func (self Struct) Accept(v Visitor) {
	v.VisitStmt(self)
}
