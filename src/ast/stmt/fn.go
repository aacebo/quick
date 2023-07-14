package stmt

import (
	"quick/src/error"
	"quick/src/token"
	"quick/src/value"
)

type Fn struct {
	Name   *token.Token
	Params []*token.Token
	Body   []Stmt
}

func (self *Fn) Accept(v Visitor) (value.Value, *error.Error) {
	return v.VisitFnStmt(self)
}
