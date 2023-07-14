package expr

import (
	"quick/src/error"
	"quick/src/value"
)

type Expr interface {
	Accept(Visitor) (value.Value, *error.Error)
}
