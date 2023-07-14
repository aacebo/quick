package expr

import "quick/src/value"

type Visitor interface {
	VisitExpr(Expr) value.Value
}
