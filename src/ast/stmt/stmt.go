package stmt

import (
	"quick/src/error"
	"quick/src/value"
)

type Stmt interface {
	Accept(Visitor) (value.Value, *error.Error)
}
