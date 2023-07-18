package expr

import (
	"quick/src/error"
	"quick/src/value"
)

type Expr interface {
	CheckType() (*value.Definition, *error.Error)
	Accept(Visitor) (value.Value, *error.Error)
}
