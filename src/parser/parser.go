package parser

import (
	"quick/src/error"
	"quick/src/scanner"
	"quick/src/token"
)

type Parser struct {
	Curr   *token.Token
	Prev   *token.Token
	Errors []*error.Error

	scanner *scanner.Scanner
}

func New(src []byte) *Parser {
	return &Parser{
		Curr:    nil,
		Prev:    nil,
		Errors:  []*error.Error{},
		scanner: scanner.New(src),
	}
}

func (self *Parser) Next() bool {
	self.Prev = self.Curr
	t, err := self.scanner.Next()

	if err != nil {
		self.Errors = append(self.Errors, err)
		return self.Next()
	}

	self.Curr = t

	if t.Kind == token.EOF {
		return false
	}

	return true
}

func (self *Parser) Match(kind token.Kind) bool {
	if self.Curr.Kind != kind {
		return false
	}

	self.Next()
	return true
}

func (self *Parser) Consume(kind token.Kind, message string) {
	if self.Curr.Kind == kind {
		self.Next()
		return
	}

	self.Errors = append(self.Errors, error.New(
		self.Curr.Ln,
		self.Curr.Start,
		self.Curr.End,
		message,
	))
}

func (self *Parser) Sync() {
	for self.Curr.Kind != token.EOF {
		if self.Prev.Kind == token.SEMI_COLON {
			return
		}

		switch self.Curr.Kind {
		case token.STRUCT:
		case token.FN:
		case token.LET:
		case token.CONST:
		case token.FOR:
		case token.IF:
		case token.PRINT:
		case token.RETURN:
			return
		default:
		}

		self.Next()
	}
}
