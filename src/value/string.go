package value

type String string

func (self String) Kind() Kind {
	return COMPLEX
}

func (self String) ComplexKind() ComplexKind {
	return STRING
}

func (self String) Truthy() Bool {
	return self != ""
}

func (self String) Name() string {
	return "string"
}

func (self String) String() string {
	return string(self)
}

func (self String) TypeEq(other Value) bool {
	return other.Kind() == COMPLEX &&
		other.ComplexKind() == STRING
}

func (self String) Eq(other Comparable) Bool {
	return self == other.(String)
}

func (self String) Concat(other Concatable) Concatable {
	return self + other.(String)
}
