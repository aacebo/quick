package parser

import "quick/src/token"

type Precedence int

const (
	NONE = iota
	ASSIGNMENT
	OR
	AND
	EQUALITY
	COMPARISON
	TERM
	FACTOR
	UNARY
	CALL
	PRIMARY
)

func GetTokenPrecedence(kind token.Kind) Precedence {
	switch kind {
	case token.LEFT_PAREN:
	case token.DOT:
		return CALL
	case token.PLUS:
	case token.MINUS:
		return TERM
	case token.STAR:
	case token.SLASH:
		return FACTOR
	case token.EQ_EQ:
	case token.NOT_EQ:
		return EQUALITY
	case token.GT:
	case token.GT_EQ:
	case token.LT:
	case token.LT_EQ:
		return COMPARISON
	case token.AND:
		return AND
	case token.OR:
		return OR
	default:
	}

	return NONE
}
