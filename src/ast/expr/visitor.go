package expr

import (
	"quick/src/error"
	"quick/src/value"
)

type Visitor interface {
	VisitAssignExpr(*Assign) (value.Value, *error.Error)
	VisitBinaryExpr(*Binary) (value.Value, *error.Error)
	VisitCallExpr(*Call) (value.Value, *error.Error)
	VisitGetExpr(*Get) (value.Value, *error.Error)
	VisitGroupingExpr(*Grouping) (value.Value, *error.Error)
	VisitLiteralExpr(*Literal) (value.Value, *error.Error)
	VisitLogicalExpr(*Logical) (value.Value, *error.Error)
	VisitSelfExpr(*Self) (value.Value, *error.Error)
	VisitSetExpr(*Set) (value.Value, *error.Error)
	VisitUnaryExpr(*Unary) (value.Value, *error.Error)
	VisitVarExpr(*Var) (value.Value, *error.Error)
}
