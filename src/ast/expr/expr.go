package expr

import (
	"quick/src/error"
	"quick/src/value"
)

type Expr interface {
	CheckValue() (value.Value, *error.Error)
	Accept(Visitor) (value.Value, *error.Error)
}
