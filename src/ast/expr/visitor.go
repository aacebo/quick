package expr

import (
	"quick/src/error"
	"quick/src/reflect"
)

type Visitor interface {
	VisitAssignExpr(*Assign) (*reflect.Value, *error.Error)
	VisitBinaryExpr(*Binary) (*reflect.Value, *error.Error)
	VisitCallExpr(*Call) (*reflect.Value, *error.Error)
	VisitGetExpr(*Get) (*reflect.Value, *error.Error)
	VisitGroupingExpr(*Grouping) (*reflect.Value, *error.Error)
	VisitLiteralExpr(*Literal) (*reflect.Value, *error.Error)
	VisitLogicalExpr(*Logical) (*reflect.Value, *error.Error)
	VisitSelfExpr(*Self) (*reflect.Value, *error.Error)
	VisitSetExpr(*Set) (*reflect.Value, *error.Error)
	VisitUnaryExpr(*Unary) (*reflect.Value, *error.Error)
	VisitVarExpr(*Var) (*reflect.Value, *error.Error)
}
