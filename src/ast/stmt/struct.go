package stmt

import (
	"quick/src/error"
	"quick/src/token"
	"quick/src/value"
)

type Struct struct {
	Name    *token.Token
	Methods []*Fn
}

func (self *Struct) Accept(v Visitor) (value.Value, *error.Error) {
	return v.VisitStructStmt(self)
}
