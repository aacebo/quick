package parser

import (
	"quick/src/ast/expr"
	"quick/src/ast/stmt"
	"quick/src/error"
	"quick/src/scanner"
	"quick/src/token"
	"quick/src/value"
)

type Parser struct {
	path    string
	curr    *token.Token
	prev    *token.Token
	errs    []*error.Error
	scanner *scanner.Scanner
}

func New(path string, src []byte) *Parser {
	return &Parser{
		path:    path,
		curr:    nil,
		prev:    nil,
		errs:    []*error.Error{},
		scanner: scanner.New(path, src),
	}
}

func (self *Parser) Parse() ([]stmt.Stmt, []*error.Error) {
	self.next()
	stmts := []stmt.Stmt{}

	for {
		if self.curr.Kind == token.EOF {
			break
		}

		stmt, err := self.declaration()

		if err != nil {
			self.errs = append(self.errs, err)
			self.sync()
			continue
		}

		stmts = append(stmts, stmt)
	}

	return stmts, self.errs
}

/*
 * Statements
 */

func (self *Parser) statement() (stmt.Stmt, *error.Error) {
	if self.match(token.FOR) {
		return self._for()
	} else if self.match(token.IF) {
		return self._if()
	} else if self.match(token.PRINT) {
		return self._print()
	} else if self.match(token.RETURN) {
		return self._return()
	} else if self.match(token.LEFT_BRACE) {
		v, e := self.block()
		return stmt.NewBlock(v), e
	}

	return self.expr()
}

func (self *Parser) declaration() (stmt.Stmt, *error.Error) {
	if self.match(token.STRUCT) {
		return self._struct()
	} else if self.match(token.FN) {
		return self.fn()
	} else if self.match(token.LET) || self.match(token.CONST) {
		return self._var()
	}

	return self.statement()
}

func (self *Parser) _if() (stmt.Stmt, *error.Error) {
	_, err := self.consume(token.LEFT_PAREN, "expected '('")

	if err != nil {
		return nil, err
	}

	cond, err := self.expression()

	if err != nil {
		return nil, err
	}

	_, err = self.consume(token.RIGHT_PAREN, "expected ')'")

	if err != nil {
		return nil, err
	}

	then, err := self.statement()

	if err != nil {
		return nil, err
	}

	var _else stmt.Stmt = nil

	if self.match(token.ELSE) {
		_else, err = self.statement()
	}

	return stmt.NewIf(cond, then, _else), nil
}

func (self *Parser) _print() (stmt.Stmt, *error.Error) {
	value, err := self.expression()

	if err != nil {
		return nil, err
	}

	_, err = self.consume(token.SEMI_COLON, "expected ';'")

	if err != nil {
		return nil, err
	}

	return stmt.NewPrint(value), nil
}

func (self *Parser) _return() (stmt.Stmt, *error.Error) {
	var value expr.Expr
	keyword := self.prev

	if self.curr.Kind != token.SEMI_COLON {
		v, err := self.expression()

		if err != nil {
			return nil, err
		}

		value = v
	}

	return stmt.NewReturn(keyword, value), nil
}

func (self *Parser) _var() (stmt.Stmt, *error.Error) {
	var init expr.Expr = nil
	keyword := self.prev
	name, err := self.consume(token.IDENTIFIER, "expected variable name")

	if err != nil {
		return nil, err
	}

	if self.match(token.EQ) {
		_init, err := self.expression()

		if err != nil {
			return nil, err
		}

		init = _init
	}

	_, err = self.consume(token.SEMI_COLON, "expected ';'")

	if err != nil {
		return nil, err
	}

	return stmt.NewVar(keyword, name, init), nil
}

func (self *Parser) _struct() (stmt.Stmt, *error.Error) {
	name, err := self.consume(token.IDENTIFIER, "expected struct name")

	if err != nil {
		return nil, err
	}

	_, err = self.consume(token.LEFT_BRACE, "expected '{'")

	if err != nil {
		return nil, err
	}

	methods := []*stmt.Fn{}

	for self.curr.Kind != token.RIGHT_BRACE && self.curr.Kind != token.EOF {
		fn, err := self.fn()

		if err != nil {
			return nil, err
		}

		methods = append(methods, fn.(*stmt.Fn))
	}

	_, err = self.consume(token.RIGHT_BRACE, "expected '}'")

	if err != nil {
		return nil, err
	}

	return stmt.NewStruct(name, methods), nil
}

func (self *Parser) _for() (stmt.Stmt, *error.Error) {
	var init stmt.Stmt = nil
	var cond expr.Expr = nil
	var inc expr.Expr = nil

	_, err := self.consume(token.LEFT_PAREN, "expected '('")

	if err != nil {
		return nil, err
	}

	if self.match(token.SEMI_COLON) {
		init = nil
	} else if self.match(token.LET) || self.match(token.CONST) {
		s, err := self._var()

		if err != nil {
			return nil, err
		}

		init = s
	} else {
		s, err := self.expr()

		if err != nil {
			return nil, err
		}

		init = s
	}

	if self.curr.Kind != token.SEMI_COLON {
		e, err := self.expression()

		if err != nil {
			return nil, err
		}

		cond = e
	}

	_, err = self.consume(token.SEMI_COLON, "expected ';'")

	if err != nil {
		return nil, err
	}

	if self.curr.Kind != token.RIGHT_PAREN {
		e, err := self.expression()

		if err != nil {
			return nil, err
		}

		inc = e
	}

	_, err = self.consume(token.RIGHT_PAREN, "expected ')'")

	if err != nil {
		return nil, err
	}

	body, err := self.statement()

	if err != nil {
		return nil, err
	}

	if cond == nil {
		cond = expr.NewLiteral(value.Bool(true))
	}

	return stmt.NewFor(init, cond, inc, body), nil
}

func (self *Parser) expr() (stmt.Stmt, *error.Error) {
	expr, err := self.expression()

	if err != nil {
		return nil, err
	}

	_, err = self.consume(token.SEMI_COLON, "expected ';'")

	if err != nil {
		return nil, err
	}

	return stmt.NewExpr(expr), nil
}

func (self *Parser) fn() (stmt.Stmt, *error.Error) {
	params := []*token.Token{}
	name, err := self.consume(token.IDENTIFIER, "expected function name")

	if err != nil {
		return nil, err
	}

	_, err = self.consume(token.LEFT_PAREN, "expected '('")

	if err != nil {
		return nil, err
	}

	if self.curr.Kind != token.RIGHT_PAREN {
		for {
			param, err := self.consume(token.IDENTIFIER, "expected parameter name")

			if err != nil {
				return nil, err
			}

			params = append(params, param)

			if !self.match(token.COMMA) {
				break
			}
		}
	}

	_, err = self.consume(token.RIGHT_PAREN, "expected ')'")

	if err != nil {
		return nil, err
	}

	_, err = self.consume(token.LEFT_BRACE, "expected '{'")

	if err != nil {
		return nil, err
	}

	body, err := self.block()

	if err != nil {
		return nil, err
	}

	return stmt.NewFn(name, params, body), nil
}

func (self *Parser) block() ([]stmt.Stmt, *error.Error) {
	stmts := []stmt.Stmt{}

	for self.curr.Kind != token.RIGHT_BRACE && self.curr.Kind != token.EOF {
		stmt, err := self.declaration()

		if err != nil {
			return nil, err
		}

		stmts = append(stmts, stmt)
	}

	_, err := self.consume(token.RIGHT_BRACE, "expected '}'")

	if err != nil {
		return nil, err
	}

	return stmts, nil
}

/*
 * Expressions
 */

func (self *Parser) expression() (expr.Expr, *error.Error) {
	return self.assignment()
}

func (self *Parser) assignment() (expr.Expr, *error.Error) {
	e, err := self.or()

	if err != nil {
		return nil, err
	}

	if self.match(token.EQ) {
		value, err := self.assignment()

		if err != nil {
			return nil, err
		}

		switch e.(type) {
		case *expr.Var:
			_var := e.(*expr.Var)
			return expr.NewAssign(_var.Name, value), nil
		case *expr.Get:
			get := e.(*expr.Get)
			return expr.NewGet(get.Object, get.Name), nil
		}

		return nil, self.error("invalid assignment target")
	}

	return e, nil
}

func (self *Parser) or() (expr.Expr, *error.Error) {
	e, err := self.and()

	if err != nil {
		return nil, err
	}

	for self.match(token.OR) {
		op := self.prev
		right, err := self.and()

		if err != nil {
			return nil, err
		}

		e = expr.NewLogical(e, op, right)
	}

	return e, nil
}

func (self *Parser) and() (expr.Expr, *error.Error) {
	e, err := self.equality()

	if err != nil {
		return nil, err
	}

	for self.match(token.AND) {
		op := self.prev
		right, err := self.equality()

		if err != nil {
			return nil, err
		}

		e = expr.NewLogical(e, op, right)
	}

	return e, nil
}

func (self *Parser) equality() (expr.Expr, *error.Error) {
	e, err := self.comparison()

	if err != nil {
		return nil, err
	}

	for self.match(token.EQ_EQ) || self.match(token.NOT_EQ) {
		op := self.prev
		right, err := self.comparison()

		if err != nil {
			return nil, err
		}

		e = expr.NewBinary(e, op, right)
	}

	return e, nil
}

func (self *Parser) comparison() (expr.Expr, *error.Error) {
	e, err := self.term()

	if err != nil {
		return nil, err
	}

	for self.match(token.GT) || self.match(token.GT_EQ) || self.match(token.LT) || self.match(token.LT_EQ) {
		op := self.prev
		right, err := self.term()

		if err != nil {
			return nil, err
		}

		e = expr.NewBinary(e, op, right)
	}

	return e, nil
}

func (self *Parser) term() (expr.Expr, *error.Error) {
	e, err := self.factor()

	if err != nil {
		return nil, err
	}

	for self.match(token.PLUS) || self.match(token.MINUS) {
		op := self.prev
		right, err := self.factor()

		if err != nil {
			return nil, err
		}

		e = expr.NewBinary(e, op, right)
	}

	return e, nil
}

func (self *Parser) factor() (expr.Expr, *error.Error) {
	e, err := self.unary()

	if err != nil {
		return nil, err
	}

	for self.match(token.STAR) || self.match(token.SLASH) {
		op := self.prev
		right, err := self.unary()

		if err != nil {
			return nil, err
		}

		e = expr.NewBinary(e, op, right)
	}

	return e, nil
}

func (self *Parser) unary() (expr.Expr, *error.Error) {
	if self.match(token.NOT) || self.match(token.MINUS) {
		op := self.prev
		right, err := self.unary()

		if err != nil {
			return nil, err
		}

		return expr.NewUnary(op, right), nil
	}

	return self.call()
}

func (self *Parser) call() (expr.Expr, *error.Error) {
	e, err := self.primary()

	if err != nil {
		return nil, err
	}

	for {
		if self.match(token.LEFT_PAREN) {
			args := []expr.Expr{}

			if self.curr.Kind != token.RIGHT_PAREN {
				for {
					e, err := self.expression()

					if err != nil {
						return nil, err
					}

					args = append(args, e)

					if !self.match(token.COMMA) {
						break
					}
				}
			}

			paren, err := self.consume(token.RIGHT_PAREN, "expected ')'")

			if err != nil {
				return nil, err
			}

			e = expr.NewCall(e, paren, args)
		} else if self.match(token.DOT) {
			name, err := self.consume(token.IDENTIFIER, "expected property name")

			if err != nil {
				return nil, err
			}

			e = expr.NewGet(e, name)
		} else {
			break
		}
	}

	return e, nil
}

func (self *Parser) primary() (expr.Expr, *error.Error) {
	if self.match(token.TRUE) || self.match(token.FALSE) {
		v, err := self.prev.Bool()

		if err != nil {
			return nil, self.error(err.Error())
		}

		return expr.NewLiteral(value.Bool(v)), nil
	} else if self.match(token.LINT) {
		v, err := self.prev.Int()

		if err != nil {
			return nil, self.error(err.Error())
		}

		return expr.NewLiteral(value.Int(v)), nil
	} else if self.match(token.LFLOAT) {
		v, err := self.prev.Float()

		if err != nil {
			return nil, self.error(err.Error())
		}

		return expr.NewLiteral(value.Float(v)), nil
	} else if self.match(token.LSTRING) {
		return expr.NewLiteral(value.String(self.prev.String())), nil
	} else if self.match(token.NIL) {
		return expr.NewLiteral(value.Nil{}), nil
	} else if self.match(token.SELF) {
		return expr.NewSelf(self.prev), nil
	} else if self.match(token.IDENTIFIER) {
		return expr.NewVar(self.prev), nil
	} else if self.match(token.LEFT_PAREN) {
		e, err := self.expression()

		if err != nil {
			return nil, err
		}

		self.consume(token.RIGHT_PAREN, "expected ')'")
		return expr.NewGrouping(e), nil
	}

	return nil, self.error("expected expression")
}

func (self *Parser) next() bool {
	self.prev = self.curr
	t, err := self.scanner.Next()

	if err != nil {
		self.errs = append(self.errs, err)
		return self.next()
	}

	self.curr = t

	if t.Kind == token.EOF {
		return false
	}

	return true
}

func (self *Parser) match(kind token.Kind) bool {
	if self.curr.Kind != kind {
		return false
	}

	self.next()
	return true
}

func (self *Parser) consume(kind token.Kind, message string) (*token.Token, *error.Error) {
	if self.curr.Kind == kind {
		self.next()
		return self.prev, nil
	}

	return nil, error.New(
		self.path,
		self.curr.Ln,
		self.curr.Start,
		self.curr.End,
		message,
	)
}

func (self *Parser) sync() {
	self.next()

	for self.curr.Kind != token.EOF {
		if self.prev.Kind == token.SEMI_COLON {
			return
		}

		switch self.curr.Kind {
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

		self.next()
	}
}

func (self Parser) error(message string) *error.Error {
	return error.New(
		self.path,
		self.prev.Ln,
		self.prev.Start,
		self.prev.End,
		message,
	)
}
