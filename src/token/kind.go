package token

type Kind uint8

const (
	EOF Kind = iota

	// singles
	COMMA
	DOT
	COLON
	SEMI_COLON
	LEFT_PAREN
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	LEFT_BRACKET
	RIGHT_BRACKET
	QUESTION_MARK

	// doubles
	DOUBLE_COLON
	RETURN_TYPE

	// arithmetic
	PLUS
	PLUS_EQ
	MINUS
	MINUS_EQ
	STAR
	STAR_EQ
	SLASH
	SLASH_EQ

	// logical
	NOT
	NOT_EQ
	EQ
	EQ_EQ
	GT
	GT_EQ
	LT
	LT_EQ
	AND
	OR

	// literals
	IDENTIFIER
	LSTRING
	LBYTE
	LINT
	LFLOAT
	NIL

	// keywords
	IF
	ELSE
	FOR
	LET
	CONST
	FN
	RETURN
	STRUCT
	SELF
	PUB
	USE
	TRUE
	FALSE

	// types
	TYPE
	STRING
	BYTE
	INT
	FLOAT
	BOOL
	MAP
)

var Keywords = map[string]Kind{
	"if":     IF,
	"else":   ELSE,
	"for":    FOR,
	"let":    LET,
	"const":  CONST,
	"fn":     FN,
	"return": RETURN,
	"struct": STRUCT,
	"self":   SELF,
	"pub":    PUB,
	"use":    USE,
	"true":   TRUE,
	"false":  FALSE,
	"string": TYPE,
	"byte":   TYPE,
	"int":    TYPE,
	"float":  TYPE,
	"bool":   TYPE,
	"map":    TYPE,
}

var Types = map[string]Kind{
	"string": STRING,
	"byte":   BYTE,
	"int":    INT,
	"float":  FLOAT,
	"bool":   BOOL,
	"map":    MAP,
}
