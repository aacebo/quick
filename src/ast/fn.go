package ast

import (
	"quick/src/ast/stmt"
	"quick/src/scope"
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

func (self Fn) String() string {
	return "<fn " + self.stmt.Name.String() + ">"
}

func (self *Fn) Call(parent *scope.Scope, args []value.Value) value.Value {
	self.ast = NewChild(parent)

	for i, t := range self.stmt.Params {
		self.ast.scope.Set(t.String(), args[i])
	}

	self.ast.Interpret(self.stmt.Body)
	self.ast = nil
	return value.Nil{}
}
