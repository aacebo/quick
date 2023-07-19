package expr

import (
	"quick/src/error"
	"quick/src/token"
	"quick/src/value"
)

type Self struct {
	Keyword *token.Token
	Value   value.Value
}

func NewSelf(keyword *token.Token, value value.Value) *Self {
	return &Self{
		Keyword: keyword,
		Value:   value,
	}
}

func (self *Self) CheckValue() (value.Value, *error.Error) {
	return self.Value, nil
}

func (self *Self) Accept(v Visitor) (value.Value, *error.Error) {
	return v.VisitSelfExpr(self)
}
