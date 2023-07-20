package expr

import (
	"quick/src/error"
	"quick/src/reflect"
	"quick/src/token"
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

func (self *Logical) GetType() (reflect.Type, *error.Error) {
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

func (self *Logical) Accept(v Visitor) (*reflect.Value, *error.Error) {
	return v.VisitLogicalExpr(self)
}
