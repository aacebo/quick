package expr

import (
	"quick/src/error"
	"quick/src/token"
	"quick/src/value"
)

type Logical struct {
	Left  Expr
	Op    *token.Token
	Right Expr
}

func NewLogical(left Expr, op *token.Token, right Expr) *Logical {
	return &Logical{
		Left:  left,
		Op:    op,
		Right: right,
	}
}

func (self *Logical) CheckValue() (value.Value, *error.Error) {
	left, err := self.Left.CheckValue()

	if err != nil {
		return nil, err
	}

	right, err := self.Right.CheckValue()

	if err != nil {
		return nil, err
	}

	if !left.TypeEq(right) {
		return nil, error.New(
			self.Op.Path,
			self.Op.Ln,
			self.Op.Start,
			self.Op.End,
			"type '"+left.Name()+"' is not '"+right.Name()+"'",
		)
	}

	return left, nil
}

func (self *Logical) Accept(v Visitor) (value.Value, *error.Error) {
	return v.VisitLogicalExpr(self)
}
