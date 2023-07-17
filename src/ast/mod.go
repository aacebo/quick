package ast

import (
	"quick/src/ast/stmt"
	"quick/src/error"
	"quick/src/value"
)

type Mod struct {
	stmt *stmt.Use
	ast  *AST
}

func NewMod(stmt *stmt.Use) *Mod {
	return &Mod{
		stmt: stmt,
		ast:  New(),
	}
}

func (self Mod) Kind() value.Kind {
	return value.COMPLEX
}

func (self Mod) ComplexKind() value.ComplexKind {
	return value.MODULE
}

func (self Mod) Truthy() value.Bool {
	return false
}

func (self Mod) String() string {
	return "<mod " + self.stmt.Name.String() + ">"
}

func (self *Mod) Call() (value.Value, *error.Error) {
	return self.ast.Interpret(self.stmt.Stmts)
}
