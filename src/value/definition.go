package value

type Definition struct {
	Kind        Kind
	ComplexKind ComplexKind
	IsConst     bool
	IsNilable   bool
}

func NewDefinition(
	kind Kind,
	ckind ComplexKind,
	isConst bool,
	isNilable bool,
) *Definition {
	return &Definition{
		Kind:        kind,
		ComplexKind: ckind,
		IsConst:     isConst,
		IsNilable:   isNilable,
	}
}

func NewFnDefinition() *Definition {
	return &Definition{
		Kind:        COMPLEX,
		ComplexKind: FUNCTION,
		IsConst:     true,
		IsNilable:   false,
	}
}

func NewStructDefinition() *Definition {
	return &Definition{
		Kind:        COMPLEX,
		ComplexKind: STRUCT,
		IsConst:     true,
		IsNilable:   false,
	}
}

func NewModDefinition() *Definition {
	return &Definition{
		Kind:        COMPLEX,
		ComplexKind: MODULE,
		IsConst:     true,
		IsNilable:   false,
	}
}
