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

func (self String) String() string {
	return string(self)
}

func (self String) Eq(other String) Bool {
	return self == other
}
