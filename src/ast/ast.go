package ast

import (
	"fmt"
	"quick/src/ast/expr"
	"quick/src/ast/stmt"
	"quick/src/error"
	"quick/src/token"
	"quick/src/value"
)

type AST struct {
	scope *Scope
}

func New() *AST {
	return &AST{
		scope: NewScope(),
	}
}

func NewChild(parent *Scope) *AST {
	return &AST{
		scope: NewChildScope(parent),
	}
}

func (self *AST) Interpret(stmts []stmt.Stmt) (value.Value, *error.Error) {
	var value value.Value = nil

	for _, stmt := range stmts {
		v, err := self.Exec(stmt)

		if err != nil {
			return nil, err
		}

		if v != nil {
			value = v
		}
	}

	return value, nil
}

func (self *AST) InterpretChild(stmts []stmt.Stmt) (value.Value, *error.Error) {
	var value value.Value = nil
	parent := self.scope
	self.scope = NewChildScope(parent)

	defer func() {
		self.scope = parent
	}()

	for _, stmt := range stmts {
		v, err := self.Exec(stmt)

		if err != nil {
			return nil, err
		}

		if v != nil {
			value = v
		}
	}

	return value, nil
}

func (self *AST) Eval(e expr.Expr) (value.Value, *error.Error) {
	return e.Accept(self)
}

func (self *AST) Exec(s stmt.Stmt) (value.Value, *error.Error) {
	return s.Accept(self)
}

/*
 * Statements
 */

func (self *AST) VisitBlockStmt(s *stmt.Block) (value.Value, *error.Error) {
	return self.InterpretChild(s.Stmts)
}

func (self *AST) VisitExprStmt(s *stmt.Expr) (value.Value, *error.Error) {
	return self.Eval(s.Expr)
}

func (self *AST) VisitForStmt(s *stmt.For) (value.Value, *error.Error) {
	parent := self.scope
	self.scope = NewChildScope(parent)

	defer func() {
		self.scope = parent
	}()

	if s.Init != nil {
		_, err := self.Exec(s.Init)

		if err != nil {
			return nil, err
		}
	}

	for {
		cond, err := self.Eval(s.Cond)

		if err != nil {
			return nil, err
		}

		if cond == nil || !cond.Truthy() {
			break
		}

		_, err = self.Exec(s.Body)

		if err != nil {
			return nil, err
		}

		_, err = self.Eval(s.Inc)

		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}

func (self *AST) VisitFnStmt(s *stmt.Fn) (value.Value, *error.Error) {
	fn := NewFn(s)
	self.scope.Set(s.Name.String(), fn)
	return nil, nil
}

func (self *AST) VisitIfStmt(s *stmt.If) (value.Value, *error.Error) {
	v, err := self.Eval(s.Cond)

	if err != nil {
		return nil, err
	}

	if v.Truthy() {
		self.Exec(s.Then)
	} else if s.Else != nil {
		self.Exec(s.Else)
	}

	return nil, nil
}

func (self *AST) VisitPrintStmt(s *stmt.Print) (value.Value, *error.Error) {
	v, err := self.Eval(s.Expr)

	if err != nil {
		return nil, err
	}

	fmt.Print(v.String())
	return nil, nil
}

func (self *AST) VisitReturnStmt(s *stmt.Return) (value.Value, *error.Error) {
	var value value.Value = nil

	if s.Value != nil {
		v, err := self.Eval(s.Value)

		if err != nil {
			return nil, err
		}

		value = v
	}

	return value, nil
}

func (self *AST) VisitStructStmt(s *stmt.Struct) (value.Value, *error.Error) {
	return nil, nil
}

func (self *AST) VisitVarStmt(s *stmt.Var) (value.Value, *error.Error) {
	var value value.Value = value.Nil{}

	if s.Init != nil {
		v, err := self.Eval(s.Init)

		if err != nil {
			return nil, err
		}

		value = v
	}

	self.scope.Set(s.Name.String(), value)
	return nil, nil
}

func (self *AST) VisitUseStmt(s *stmt.Use) (value.Value, *error.Error) {
	mod := NewMod(s)
	_, err := mod.Call()

	if err != nil {
		return nil, err
	}

	self.scope.Set(s.Name.String(), mod)
	return nil, nil
}

/*
 * Expressions
 */

func (self *AST) VisitAssignExpr(e *expr.Assign) (value.Value, *error.Error) {
	value, err := self.Eval(e.Value)

	if err != nil {
		return nil, err
	}

	self.scope.Set(e.Name.String(), value)
	return value, nil
}

func (self *AST) VisitBinaryExpr(e *expr.Binary) (value.Value, *error.Error) {
	left, err := self.Eval(e.Left)

	if err != nil {
		return nil, err
	}

	right, err := self.Eval(e.Right)

	if err != nil {
		return nil, err
	}

	switch e.Op.Kind {
	case token.EQ_EQ:
		return left.(value.Comparable).Eq(right.(value.Comparable)), nil
	case token.NOT_EQ:
		return !left.(value.Comparable).Eq(right.(value.Comparable)), nil
	case token.GT:
		return left.(value.Numeric).Gt(right.(value.Numeric)), nil
	case token.GT_EQ:
		return left.(value.Numeric).GtEq(right.(value.Numeric)), nil
	case token.LT:
		return left.(value.Numeric).Lt(right.(value.Numeric)), nil
	case token.LT_EQ:
		return left.(value.Numeric).LtEq(right.(value.Numeric)), nil
	case token.PLUS:
		switch left.(type) {
		case value.Numeric:
			return left.(value.Numeric).Add(right.(value.Numeric)), nil
		case value.Concatable:
			return left.(value.Concatable).Concat(right.(value.Concatable)), nil
		}

		return nil, error.New(
			e.Op.Path,
			e.Op.Ln,
			e.Op.Start,
			e.Op.End,
			"invalid operands",
		)
	case token.MINUS:
		return left.(value.Numeric).Subtract(right.(value.Numeric)), nil
	case token.STAR:
		return left.(value.Numeric).Multiply(right.(value.Numeric)), nil
	case token.SLASH:
		return left.(value.Numeric).Divide(right.(value.Numeric)), nil
	}

	return nil, nil
}

func (self *AST) VisitCallExpr(e *expr.Call) (value.Value, *error.Error) {
	callee, err := self.Eval(e.Callee)

	if err != nil {
		return nil, err
	}

	args := []value.Value{}

	for _, arg := range e.Args {
		v, err := self.Eval(arg)

		if err != nil {
			return nil, err
		}

		args = append(args, v)
	}

	fn, ok := callee.(*Fn)

	if !ok {
		return nil, error.New(
			e.Paren.Path,
			e.Paren.Ln,
			e.Paren.Start,
			e.Paren.End,
			"expected function",
		)
	}

	if len(args) != len(fn.stmt.Params) {
		return nil, error.New(
			e.Paren.Path,
			e.Paren.Ln,
			e.Paren.Start,
			e.Paren.End,
			fmt.Sprintf(
				"expected %d arguments, received %d",
				len(fn.stmt.Params),
				len(args),
			),
		)
	}

	return fn.Call(self.scope, args)
}

func (self *AST) VisitGetExpr(e *expr.Get) (value.Value, *error.Error) {
	value, err := self.Eval(e.Object)

	if err != nil {
		return nil, err
	}

	switch value.(type) {
	case *Mod:
		return value.(*Mod).ast.scope.GetLocal(e.Name.String()), nil
	}

	return nil, error.New(
		e.Name.Path,
		e.Name.Ln,
		e.Name.Start,
		e.Name.End,
		"expected object",
	)
}

func (self *AST) VisitGroupingExpr(e *expr.Grouping) (value.Value, *error.Error) {
	return self.Eval(e.Expr)
}

func (self *AST) VisitLiteralExpr(e *expr.Literal) (value.Value, *error.Error) {
	return e.Value, nil
}

func (self *AST) VisitLogicalExpr(e *expr.Logical) (value.Value, *error.Error) {
	left, err := self.Eval(e.Left)

	if err != nil {
		return nil, err
	}

	if e.Op.Kind == token.OR {
		if left.Truthy() {
			return left, nil
		}
	} else {
		if !left.Truthy() {
			return left, nil
		}
	}

	return self.Eval(e.Right)
}

func (self *AST) VisitSelfExpr(e *expr.Self) (value.Value, *error.Error) {
	return self.scope.GetLocal("self"), nil
}

func (self *AST) VisitSetExpr(e *expr.Set) (value.Value, *error.Error) {
	return nil, nil
}

func (self *AST) VisitUnaryExpr(e *expr.Unary) (value.Value, *error.Error) {
	right, err := self.Eval(e.Right)

	if err != nil {
		return nil, err
	}

	switch e.Op.Kind {
	case token.NOT:
		return !right.Truthy(), nil
	case token.MINUS:
		v := right.(value.Numeric)
		v.Dec()
		return v, nil
	}

	return nil, nil
}

func (self *AST) VisitVarExpr(e *expr.Var) (value.Value, *error.Error) {
	if !self.scope.Has(e.Name.String()) {
		return nil, error.New(
			e.Name.Path,
			e.Name.Ln,
			e.Name.Start,
			e.Name.End,
			"undefined identifier",
		)
	}

	return self.scope.Get(e.Name.String()), nil
}
