package stmt

import (
	"quick/src/error"
	"quick/src/reflect"
)

type Stmt interface {
	Accept(Visitor) (*reflect.Value, *error.Error)
}
