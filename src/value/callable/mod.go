package callable

import (
	"quick/src/ast/stmt"
	"quick/src/value"
)

type Mod struct {
	Stmt   *stmt.Use
	Values map[string]value.Value
}

func NewMod(stmt *stmt.Use) *Mod {
	return &Mod{
		Stmt:   stmt,
		Values: map[string]value.Value{},
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

func (self Mod) Name() string {
	name := ""

	for i, p := range self.Stmt.Path {
		name += p.String()

		if i < len(self.Stmt.Path)-1 {
			name += "::"
		}
	}
	return name
}

func (self Mod) String() string {
	return "<mod " + self.Name() + ">"
}

func (self Mod) TypeEq(other value.Value) bool {
	return other.Kind() == value.COMPLEX &&
		other.ComplexKind() == value.MODULE
}
