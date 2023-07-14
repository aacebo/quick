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
	String() string
}
