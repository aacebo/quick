package expr

import (
	"quick/src/error"
	"quick/src/reflect"
)

type Expr interface {
	GetType() (reflect.Type, *error.Error)
	Accept(Visitor) (*reflect.Value, *error.Error)
}
