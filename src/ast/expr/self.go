package expr

import (
	"quick/src/error"
	"quick/src/token"
	"quick/src/value"
)

type Self struct {
	Type    *value.Definition
	Keyword *token.Token
}

func NewSelf(
	_type *value.Definition,
	keyword *token.Token,
) *Self {
	return &Self{
		Type:    _type,
		Keyword: keyword,
	}
}

func (self *Self) CheckType() (*value.Definition, *error.Error) {
	return self.Type, nil
}

func (self *Self) Accept(v Visitor) (value.Value, *error.Error) {
	return v.VisitSelfExpr(self)
}
