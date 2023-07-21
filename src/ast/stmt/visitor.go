package stmt

import (
	"quick/src/error"
	"quick/src/reflect"
)

type Visitor interface {
	VisitBlockStmt(*Block) (*reflect.Value, *error.Error)
	VisitExprStmt(*Expr) (*reflect.Value, *error.Error)
	VisitForStmt(*For) (*reflect.Value, *error.Error)
	VisitFnStmt(*Fn) (*reflect.Value, *error.Error)
	VisitIfStmt(*If) (*reflect.Value, *error.Error)
	VisitReturnStmt(*Return) (*reflect.Value, *error.Error)
	VisitStructStmt(*Struct) (*reflect.Value, *error.Error)
	VisitVarStmt(*Var) (*reflect.Value, *error.Error)
	VisitUseStmt(*Use) (*reflect.Value, *error.Error)
}
