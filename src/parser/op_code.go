package parser

type OpCode int

const (
	CONST OpCode = iota

	POP
	DEFINE
	RESOLVE
	ASSIGN

	EQ
	GT
	LT

	ADD
	SUBTRACT
	MULTIPLY
	DIVIDE

	NOT
	NEGATE
	PRINT
	JUMP
	JUMP_IF_FALSE
	LOOP
	INVOKE
	RETURN
	FN
	STRUCT
)
