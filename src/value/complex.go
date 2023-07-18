package value

type ComplexKind int

const (
	NONE ComplexKind = iota
	SLICE
	STRING
	MAP
	FUNCTION
	STRUCT
	MODULE
)

type Complex interface {
	ComplexKind() ComplexKind
	Name() string
	String() string
}
