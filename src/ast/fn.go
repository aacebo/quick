package ast

import (
	"quick/src/ast/stmt"
	"quick/src/error"
	"quick/src/value"
)

type Fn struct {
	stmt *stmt.Fn
	ast  *AST
}

func NewFn(stmt *stmt.Fn) *Fn {
	return &Fn{
		stmt: stmt,
		ast:  nil,
	}
}

func (self Fn) Kind() value.Kind {
	return value.COMPLEX
}

func (self Fn) ComplexKind() value.ComplexKind {
	return value.FUNCTION
}

func (self Fn) Truthy() value.Bool {
	return false
}

func (self Fn) Name() string {
	return self.stmt.Name.String()
}

func (self Fn) String() string {
	return "<fn " + self.stmt.Name.String() + ">"
}

func (self *Fn) Call(parent *Scope, args []value.Value) (value.Value, *error.Error) {
	self.ast = NewChild(parent)

	defer func() {
		self.ast = nil
	}()

	for i, param := range self.stmt.Params {
		self.ast.scope.Set(param.Name.String(), args[i])
	}

	return self.ast.InterpretChild(self.stmt.Body)
}
