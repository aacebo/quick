package expr

import (
	"quick/src/error"
	"quick/src/reflect"
	"quick/src/token"
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

func (self *Binary) GetType() (reflect.Type, *error.Error) {
	left, err := self.Left.GetType()

	if err != nil {
		return nil, err
	}

	right, err := self.Right.GetType()

	if err != nil {
		return nil, err
	}

	if !left.Equals(right) {
		return nil, error.New(
			self.Op.Path,
			self.Op.Ln,
			self.Op.Start,
			self.Op.End,
			"expected type '"+left.Name()+"', received '"+right.Name()+"'",
		)
	}

	return left, nil
}

func (self *Binary) Accept(v Visitor) (*reflect.Value, *error.Error) {
	return v.VisitBinaryExpr(self)
}
