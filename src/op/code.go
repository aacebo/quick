package op

type Code uint8

const (
	Push Code = iota
	P

	Get
	Set
	Define

	Eq
	Gt
	Lt

	Add
	Subtract
	Multiply
	Divide

	Not
	Negate

	Jump
	JumpIfFalse
	Lo
	Call
	Fn
	Return
	Struct
	Mod
)
