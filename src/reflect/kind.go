package reflect

type Kind int

const (
	Invalid Kind = iota
	Byte
	Bool
	Int
	Float
	Nil
	Map
	String
	Slice
	Fn
	Struct
	Mod
)

func (self Kind) String() string {
	switch self {
	case Byte:
		return "byte"
	case Bool:
		return "bool"
	case Int:
		return "int"
	case Float:
		return "float"
	case Nil:
		return "nil"
	case Map:
		return "map"
	case String:
		return "string"
	case Slice:
		return "slice"
	case Fn:
		return "fn"
	case Struct:
		return "struct"
	case Mod:
		return "mod"
	default:
		return "invalid"
	}
}
