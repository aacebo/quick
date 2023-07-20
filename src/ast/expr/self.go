package expr

import (
	"quick/src/error"
	"quick/src/reflect"
	"quick/src/token"
)

type Self struct {
	Keyword *token.Token
	Type    reflect.Type
}

func NewSelf(keyword *token.Token, _type reflect.Type) *Self {
	return &Self{
		Keyword: keyword,
		Type:    _type,
	}
}

func (self *Self) GetType() (reflect.Type, *error.Error) {
	return self.Type, nil
}

func (self *Self) Accept(v Visitor) (*reflect.Value, *error.Error) {
	return v.VisitSelfExpr(self)
}
