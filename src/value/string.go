package value

type String Slice[Byte]

func (self String) Kind() Kind {
	return COMPLEX
}

func (self String) ComplexKind() ComplexKind {
	return STRING
}

func (self String) String() string {
	return string(self)
}
