package stmt

import "quick/src/value"

type Stmt interface {
	Accept(Visitor) value.Value
}
