package stmt

import "quick/src/value"

type Visitor interface {
	VisitBlockStmt(*Block) value.Value
	VisitExprStmt(*Expr) value.Value
	VisitForStmt(*For) value.Value
	VisitFnStmt(*Fn) value.Value
	VisitIfStmt(*If) value.Value
	VisitPrintStmt(*Print) value.Value
	VisitReturnStmt(*Return) value.Value
	VisitStructStmt(*Struct) value.Value
	VisitVarStmt(*Var) value.Value
}
