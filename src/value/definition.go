package value

var TypeDefinitions = map[string]*Definition{
	"bool": NewDefinition(
		BOOL,
		NONE,
		"bool",
		false,
		false,
	),
	"byte": NewDefinition(
		BYTE,
		NONE,
		"byte",
		false,
		false,
	),
	"int": NewDefinition(
		INT,
		NONE,
		"int",
		false,
		false,
	),
	"float": NewDefinition(
		FLOAT,
		NONE,
		"float",
		false,
		false,
	),
	"string": NewDefinition(
		COMPLEX,
		STRING,
		"string",
		false,
		false,
	),
	"nil": NewDefinition(
		NIL,
		NONE,
		"nil",
		false,
		false,
	),
}

type Definition struct {
	Kind        Kind
	ComplexKind ComplexKind
	Name        string
	IsConst     bool
	IsNilable   bool
}

func NewDefinition(
	kind Kind,
	ckind ComplexKind,
	name string,
	isConst bool,
	isNilable bool,
) *Definition {
	return &Definition{
		Kind:        kind,
		ComplexKind: ckind,
		Name:        name,
		IsConst:     isConst,
		IsNilable:   isNilable,
	}
}

func NewFnDefinition(name string) *Definition {
	return &Definition{
		Kind:        COMPLEX,
		ComplexKind: FUNCTION,
		Name:        name,
		IsConst:     true,
		IsNilable:   false,
	}
}

func NewStructDefinition(name string) *Definition {
	return &Definition{
		Kind:        COMPLEX,
		ComplexKind: STRUCT,
		Name:        name,
		IsConst:     true,
		IsNilable:   false,
	}
}

func NewModDefinition(name string) *Definition {
	return &Definition{
		Kind:        COMPLEX,
		ComplexKind: MODULE,
		Name:        name,
		IsConst:     true,
		IsNilable:   false,
	}
}

func (self Definition) Equals(other *Definition) bool {
	return self.Kind == other.Kind &&
		self.ComplexKind == other.ComplexKind &&
		self.Name == other.Name
}
