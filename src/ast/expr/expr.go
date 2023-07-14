package expr

import "quick/src/value"

type Expr interface {
	Accept(Visitor) value.Value
}
