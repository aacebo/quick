package value

type ComplexKind int

const (
	SLICE ComplexKind = iota
	STRING
	MAP
	FUNCTION
)

type Complex interface {
	ComplexKind() ComplexKind
	String() string
}
