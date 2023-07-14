package ast

import (
	"fmt"
	"quick/src/ast/expr"
	"quick/src/ast/stmt"
	"quick/src/scope"
	"quick/src/token"
	"quick/src/value"
)

type AST struct {
	scope *scope.Scope
}

func New() *AST {
	return &AST{
		scope: scope.New(),
	}
}

func NewChild(parent *scope.Scope) *AST {
	return &AST{
		scope: scope.NewChild(parent),
	}
}

func (self *AST) Interpret(stmts []stmt.Stmt) value.Value {
	for _, stmt := range stmts {
		self.Exec(stmt)
	}

	return value.Nil{}
}

func (self *AST) Eval(e expr.Expr) value.Value {
	return e.Accept(self)
}

func (self *AST) Exec(s stmt.Stmt) value.Value {
	return s.Accept(self)
}

/*
 * Statements
 */

func (self *AST) VisitBlockStmt(s *stmt.Block) value.Value {
	return self.Interpret(s.Stmts)
}

func (self *AST) VisitExprStmt(s *stmt.Expr) value.Value {
	return self.Eval(s.Expr)
}

func (self *AST) VisitForStmt(s *stmt.For) value.Value {
	return value.Nil{}
}

func (self *AST) VisitFnStmt(s *stmt.Fn) value.Value {
	fn := NewFn(s)
	self.scope.Set(s.Name.String(), fn)
	return fn
}

func (self *AST) VisitIfStmt(s *stmt.If) value.Value {
	v := self.Eval(s.Cond)

	if v.Truthy() {
		self.Exec(s.Then)
	} else if s.Else != nil {
		self.Exec(s.Else)
	}

	return value.Nil{}
}

func (self *AST) VisitPrintStmt(s *stmt.Print) value.Value {
	fmt.Print(self.Eval(s.Expr).String())
	return value.Nil{}
}

func (self *AST) VisitReturnStmt(s *stmt.Return) value.Value {
	var value value.Value = value.Nil{}

	if s.Value != nil {
		value = self.Eval(s.Value)
	}

	return value
}

func (self *AST) VisitStructStmt(s *stmt.Struct) value.Value {
	return value.Nil{}
}

func (self *AST) VisitVarStmt(s *stmt.Var) value.Value {
	var value value.Value = value.Nil{}

	if s.Init != nil {
		value = self.Eval(s.Init)
	}

	self.scope.Set(s.Name.String(), value)
	return value
}

/*
 * Expressions
 */

func (self *AST) VisitAssignExpr(e *expr.Assign) value.Value {
	value := self.Eval(e.Value)
	self.scope.Set(e.Name.String(), value)
	return value
}

func (self *AST) VisitBinaryExpr(e *expr.Binary) value.Value {
	left := self.Eval(e.Left).(value.Comparable)
	right := self.Eval(e.Right).(value.Comparable)

	switch e.Op.Kind {
	case token.EQ_EQ:
		return left.Eq(right)
	case token.NOT_EQ:
		return !left.Eq(right)
	case token.GT:
		return (left.(value.Numeric)).Gt(right.(value.Numeric))
	case token.GT_EQ:
		return (left.(value.Numeric)).GtEq(right.(value.Numeric))
	case token.LT:
		return (left.(value.Numeric)).Lt(right.(value.Numeric))
	case token.LT_EQ:
		return (left.(value.Numeric)).LtEq(right.(value.Numeric))
	case token.PLUS:
		return (left.(value.Numeric)).Add(right.(value.Numeric))
	case token.MINUS:
		return (left.(value.Numeric)).Subtract(right.(value.Numeric))
	case token.STAR:
		return (left.(value.Numeric)).Multiply(right.(value.Numeric))
	case token.SLASH:
		return (left.(value.Numeric)).Divide(right.(value.Numeric))
	}

	return value.Nil{}
}

func (self *AST) VisitCallExpr(e *expr.Call) value.Value {
	callee := self.Eval(e.Callee)
	args := []value.Value{}

	for _, arg := range e.Args {
		args = append(args, self.Eval(arg))
	}

	fn, ok := callee.(*Fn)

	if !ok {
		// expected function
	}

	if len(args) != len(fn.stmt.Params) {
		// expected {x} arguments, received {y}
	}

	return fn.Call(self.scope, args)
}

func (self *AST) VisitGetExpr(e *expr.Get) value.Value {
	return value.Nil{}
}

func (self *AST) VisitGroupingExpr(e *expr.Grouping) value.Value {
	return self.Eval(e.Expr)
}

func (self *AST) VisitLiteralExpr(e *expr.Literal) value.Value {
	return e.Value
}

func (self *AST) VisitLogicalExpr(e *expr.Logical) value.Value {
	left := self.Eval(e.Left)

	if e.Op.Kind == token.OR {
		if left.Truthy() {
			return left
		}
	} else {
		if !left.Truthy() {
			return left
		}
	}

	return self.Eval(e.Right)
}

func (self *AST) VisitSelfExpr(e *expr.Self) value.Value {
	return *self.scope.GetLocal("self")
}

func (self *AST) VisitSetExpr(e *expr.Set) value.Value {
	return value.Nil{}
}

func (self *AST) VisitUnaryExpr(e *expr.Unary) value.Value {
	right := self.Eval(e.Right)

	switch e.Op.Kind {
	case token.NOT:
		return !right.Truthy()
	case token.MINUS:
		v := right.(value.Numeric)
		v.Dec()
		return v
	}

	return value.Nil{}
}

func (self *AST) VisitVarExpr(e *expr.Var) value.Value {
	return *self.scope.Get(e.Name.String())
}
