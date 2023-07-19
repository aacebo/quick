package callable

import (
	"quick/src/ast/stmt"
	"quick/src/value"
)

type Fn struct {
	Stmt *stmt.Fn
}

func NewFn(stmt *stmt.Fn) *Fn {
	return &Fn{
		Stmt: stmt,
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
	return self.String()
}

func (self Fn) String() string {
	return "<fn " + self.Stmt.Name.String() + ">"
}

func (self Fn) TypeEq(other value.Value) bool {
	return other.Kind() == value.COMPLEX &&
		other.ComplexKind() == value.FUNCTION
}
