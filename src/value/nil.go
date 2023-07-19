package value

type Nil struct{}

func (self Nil) Kind() Kind {
	return NIL
}

func (self Nil) ComplexKind() ComplexKind {
	return NONE
}

func (self Nil) Truthy() Bool {
	return false
}

func (self Nil) Name() string {
	return "nil"
}

func (self Nil) String() string {
	return "<nil>"
}

func (self Nil) TypeEq(other Value) bool {
	return other.Kind() == NIL
}
