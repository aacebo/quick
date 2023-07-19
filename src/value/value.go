package value

type Kind int

const (
	BYTE Kind = iota
	BOOL
	INT
	FLOAT
	NIL
	COMPLEX
)

type Value interface {
	Kind() Kind
	ComplexKind() ComplexKind
	Truthy() Bool
	Name() string
	String() string
	TypeEq(Value) bool
}
