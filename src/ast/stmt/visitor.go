package stmt

import (
	"quick/src/error"
	"quick/src/value"
)

type Visitor interface {
	VisitBlockStmt(*Block) (value.Value, *error.Error)
	VisitExprStmt(*Expr) (value.Value, *error.Error)
	VisitForStmt(*For) (value.Value, *error.Error)
	VisitFnStmt(*Fn) (value.Value, *error.Error)
	VisitIfStmt(*If) (value.Value, *error.Error)
	VisitPrintStmt(*Print) (value.Value, *error.Error)
	VisitReturnStmt(*Return) (value.Value, *error.Error)
	VisitStructStmt(*Struct) (value.Value, *error.Error)
	VisitVarStmt(*Var) (value.Value, *error.Error)
	VisitUseStmt(*Use) (value.Value, *error.Error)
}
