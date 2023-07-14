package expr

import "quick/src/value"

type Visitor interface {
	VisitAssignExpr(*Assign) value.Value
	VisitBinaryExpr(*Binary) value.Value
	VisitCallExpr(*Call) value.Value
	VisitGetExpr(*Get) value.Value
	VisitGroupingExpr(*Grouping) value.Value
	VisitLiteralExpr(*Literal) value.Value
	VisitLogicalExpr(*Logical) value.Value
	VisitSelfExpr(*Self) value.Value
	VisitSetExpr(*Set) value.Value
	VisitUnaryExpr(*Unary) value.Value
	VisitVarExpr(*Var) value.Value
}
