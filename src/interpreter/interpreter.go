package interpreter

import (
	"fmt"
	"quick/src/ast/expr"
	"quick/src/ast/stmt"
	"quick/src/error"
	"quick/src/reflect"
	"quick/src/token"
	"quick/src/utils"
)

type Interpreter struct {
	scope *Scope
}

func New() *Interpreter {
	return &Interpreter{
		scope: NewScope(),
	}
}

func NewChild(parent *Scope) *Interpreter {
	return &Interpreter{
		scope: NewChildScope(parent),
	}
}

func (self *Interpreter) Interpret(stmts []stmt.Stmt) (*reflect.Value, *error.Error) {
	var value *reflect.Value = nil

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

func (self *Interpreter) InterpretChild(stmts []stmt.Stmt) (*reflect.Value, *error.Error) {
	var value *reflect.Value = nil
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

func (self *Interpreter) Eval(e expr.Expr) (*reflect.Value, *error.Error) {
	return e.Accept(self)
}

func (self *Interpreter) Exec(s stmt.Stmt) (*reflect.Value, *error.Error) {
	return s.Accept(self)
}

/*
 * Statements
 */

func (self *Interpreter) VisitBlockStmt(s *stmt.Block) (*reflect.Value, *error.Error) {
	return self.InterpretChild(s.Stmts)
}

func (self *Interpreter) VisitExprStmt(s *stmt.Expr) (*reflect.Value, *error.Error) {
	return self.Eval(s.Expr)
}

func (self *Interpreter) VisitForStmt(s *stmt.For) (*reflect.Value, *error.Error) {
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

func (self *Interpreter) VisitFnStmt(s *stmt.Fn) (*reflect.Value, *error.Error) {
	fn := reflect.NewFn(
		s.Name.String(),
		utils.MapSlice(s.Params, func(v *stmt.Var) reflect.Param {
			return reflect.Param{
				Name: v.Name.String(),
				Type: v.Type,
			}
		}),
		s.ReturnType,
		s,
	)

	self.scope.Set(s.Name.String(), fn)
	return nil, nil
}

func (self *Interpreter) VisitIfStmt(s *stmt.If) (*reflect.Value, *error.Error) {
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

func (self *Interpreter) VisitPrintStmt(s *stmt.Print) (*reflect.Value, *error.Error) {
	v, err := self.Eval(s.Expr)

	if err != nil {
		return nil, err
	}

	fmt.Print(v.ToString())
	return nil, nil
}

func (self *Interpreter) VisitReturnStmt(s *stmt.Return) (*reflect.Value, *error.Error) {
	var value *reflect.Value = nil

	if s.Value != nil {
		v, err := self.Eval(s.Value)

		if err != nil {
			return nil, err
		}

		value = v
	}

	return value, nil
}

func (self *Interpreter) VisitStructStmt(s *stmt.Struct) (*reflect.Value, *error.Error) {
	return nil, nil
}

func (self *Interpreter) VisitVarStmt(s *stmt.Var) (*reflect.Value, *error.Error) {
	value := reflect.NewNil()

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

func (self *Interpreter) VisitUseStmt(s *stmt.Use) (*reflect.Value, *error.Error) {
	mod := reflect.NewMod()
	sibling := New()
	_, err := sibling.Interpret(s.Stmts)

	if err != nil {
		return nil, err
	}

	if s.Path[len(s.Path)-1].Kind == token.STAR {
		for key, value := range sibling.scope.values {
			self.scope.Set(key, value)
		}
	} else {
		for key, value := range sibling.scope.values {
			mod.SetExport(key, value)
		}

		self.scope.Set(s.Path[len(s.Path)-1].String(), mod)
	}

	return nil, nil
}

/*
 * Expressions
 */

func (self *Interpreter) VisitAssignExpr(e *expr.Assign) (*reflect.Value, *error.Error) {
	value, err := self.Eval(e.Value)

	if err != nil {
		return nil, err
	}

	self.scope.Set(e.Name.String(), value)
	return value, nil
}

func (self *Interpreter) VisitBinaryExpr(e *expr.Binary) (*reflect.Value, *error.Error) {
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
		return reflect.NewBool(left.Eq(right)), nil
	case token.NOT_EQ:
		return reflect.NewBool(!left.Eq(right)), nil
	case token.GT:
		return reflect.NewBool(left.Gt(right)), nil
	case token.GT_EQ:
		return reflect.NewBool(left.GtEq(right)), nil
	case token.LT:
		return reflect.NewBool(left.Lt(right)), nil
	case token.LT_EQ:
		return reflect.NewBool(left.LtEq(right)), nil
	case token.PLUS:
		if left.Numeric() {
			return left.Add(right), nil
		}

		return reflect.NewString(left.String() + right.String()), nil
	case token.MINUS:
		return left.Subtract(right), nil
	case token.STAR:
		return left.Multiply(right), nil
	case token.SLASH:
		return left.Divide(right), nil
	}

	return nil, nil
}

func (self *Interpreter) VisitCallExpr(e *expr.Call) (*reflect.Value, *error.Error) {
	callee, err := self.Eval(e.Callee)

	if err != nil {
		return nil, err
	}

	args := []*reflect.Value{}

	for _, arg := range e.Args {
		v, err := self.Eval(arg)

		if err != nil {
			return nil, err
		}

		args = append(args, v)
	}

	if !callee.IsFn() && !callee.IsNativeFn() {
		return nil, error.New(
			e.Paren.Path,
			e.Paren.Ln,
			e.Paren.Start,
			e.Paren.End,
			"expected function",
		)
	}

	if callee.IsNativeFn() {
		return callee.NativeFn()(args), nil
	}

	if len(args) != len(callee.FnType().Params()) {
		return nil, error.New(
			e.Paren.Path,
			e.Paren.Ln,
			e.Paren.Start,
			e.Paren.End,
			fmt.Sprintf(
				"expected %d arguments, received %d",
				len(callee.FnType().Params()),
				len(args),
			),
		)
	}

	interp := NewChild(self.scope)

	for i, arg := range args {
		interp.scope.Set(callee.FnType().Params()[i].Name, arg)
	}

	return interp.InterpretChild(callee.Fn().(*stmt.Fn).Body)
}

func (self *Interpreter) VisitGetExpr(e *expr.Get) (*reflect.Value, *error.Error) {
	v, err := self.Eval(e.Object)

	if err != nil {
		return nil, err
	}

	if !v.HasMember(e.Name.String()) {
		return nil, error.New(
			e.Name.Path,
			e.Name.Ln,
			e.Name.Start,
			e.Name.End,
			"module export '"+e.Name.String()+"' not found",
		)
	}

	return v.GetMember(e.Name.String()), nil
}

func (self *Interpreter) VisitGroupingExpr(e *expr.Grouping) (*reflect.Value, *error.Error) {
	return self.Eval(e.Expr)
}

func (self *Interpreter) VisitLiteralExpr(e *expr.Literal) (*reflect.Value, *error.Error) {
	return e.Value, nil
}

func (self *Interpreter) VisitLogicalExpr(e *expr.Logical) (*reflect.Value, *error.Error) {
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

func (self *Interpreter) VisitSelfExpr(e *expr.Self) (*reflect.Value, *error.Error) {
	return self.scope.GetLocal("self"), nil
}

func (self *Interpreter) VisitSetExpr(e *expr.Set) (*reflect.Value, *error.Error) {
	return nil, nil
}

func (self *Interpreter) VisitSliceExpr(e *expr.Slice) (*reflect.Value, *error.Error) {
	slice := reflect.NewSlice(e.Type, []*reflect.Value{})

	for _, exp := range e.Items {
		value, err := self.Eval(exp)

		if err != nil {
			return nil, err
		}

		slice.Push(value)
	}

	return slice, nil
}

func (self *Interpreter) VisitUnaryExpr(e *expr.Unary) (*reflect.Value, *error.Error) {
	right, err := self.Eval(e.Right)

	if err != nil {
		return nil, err
	}

	switch e.Op.Kind {
	case token.NOT:
		return reflect.NewBool(!right.Truthy()), nil
	case token.MINUS:
		right.Decrement()
		return right, nil
	}

	return nil, nil
}

func (self *Interpreter) VisitVarExpr(e *expr.Var) (*reflect.Value, *error.Error) {
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
