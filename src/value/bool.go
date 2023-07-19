package value

type Bool bool

func (self Bool) Kind() Kind {
	return BOOL
}

func (self Bool) ComplexKind() ComplexKind {
	return NONE
}

func (self Bool) Truthy() Bool {
	return self == true
}

func (self Bool) Name() string {
	return "bool"
}

func (self Bool) String() string {
	if self == true {
		return "true"
	}

	return "false"
}

func (self Bool) TypeEq(other Value) bool {
	return other.Kind() == BOOL
}

func (self Bool) Eq(other Bool) Bool {
	return self == other
}
