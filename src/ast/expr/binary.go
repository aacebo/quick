package expr

import (
	"quick/src/error"
	"quick/src/token"
	"quick/src/value"
)

type Binary struct {
	Left  Expr
	Op    *token.Token
	Right Expr
}

func NewBinary(left Expr, op *token.Token, right Expr) *Binary {
	return &Binary{
		Left:  left,
		Op:    op,
		Right: right,
	}
}

func (self *Binary) CheckType() (*value.Definition, *error.Error) {
	left, err := self.Left.CheckType()

	if err != nil {
		return nil, err
	}

	right, err := self.Right.CheckType()

	if err != nil {
		return nil, err
	}

	if !left.Equals(right) {
		return nil, error.New(
			self.Op.Path,
			self.Op.Ln,
			self.Op.Start,
			self.Op.End,
			"type '"+left.Name+"' is not '"+right.Name+"'",
		)
	}

	return left, nil
}

func (self *Binary) Accept(v Visitor) (value.Value, *error.Error) {
	return v.VisitBinaryExpr(self)
}
