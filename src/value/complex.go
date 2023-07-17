package value

type ComplexKind int

const (
	SLICE ComplexKind = iota
	STRING
	MAP
	FUNCTION
	STRUCT
	MODULE
)

type Complex interface {
	ComplexKind() ComplexKind
	String() string
}
