package parser

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"quick/src/ast/expr"
	"quick/src/ast/stmt"
	"quick/src/error"
	"quick/src/reflect"
	"quick/src/scanner"
	"quick/src/token"
	"quick/src/utils"
)

var cache = map[string]*Parser{}

type Parser struct {
	path    string
	curr    *token.Token
	prev    *token.Token
	stmts   []stmt.Stmt
	errs    []*error.Error
	scanner *scanner.Scanner
	scope   *Scope
}

func New(path string) *Parser {
	src, err := os.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	return &Parser{
		path:    path,
		curr:    nil,
		prev:    nil,
		stmts:   []stmt.Stmt{},
		errs:    []*error.Error{},
		scanner: scanner.New(path, src),
		scope:   NewScope(),
	}
}

func (self *Parser) Parse() ([]stmt.Stmt, []*error.Error) {
	self.next()

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

		self.stmts = append(self.stmts, stmt)
	}

	return self.stmts, self.errs
}

/*
 * Statements
 */

func (self *Parser) statement() (stmt.Stmt, *error.Error) {
	if self.match(token.USE) {
		return self.use()
	} else if self.match(token.FOR) {
		return self._for()
	} else if self.match(token.IF) {
		return self._if()
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

func (self *Parser) _return() (stmt.Stmt, *error.Error) {
	var value expr.Expr = expr.NewLiteral(reflect.NewNil())
	keyword := self.prev

	if self.curr.Kind != token.SEMI_COLON {
		v, err := self.expression()

		if err != nil {
			return nil, err
		}

		value = v
	}

	_, err := self.consume(token.SEMI_COLON, "expected ';'")

	if err != nil {
		return nil, err
	}

	return stmt.NewReturn(keyword, value), nil
}

func (self *Parser) _var() (stmt.Stmt, *error.Error) {
	var _type reflect.Type = nil
	var nilable *token.Token = nil
	var init expr.Expr = nil

	isSlice := false
	keyword := self.prev
	name, err := self.consume(token.IDENTIFIER, "expected variable name")

	if err != nil {
		return nil, err
	}

	if self.scope.HasLocal(name.String()) {
		return nil, self.error("duplicate name")
	}

	if self.match(token.LEFT_BRACKET) {
		_, err = self.consume(token.RIGHT_BRACKET, "expected ']'")

		if err != nil {
			return nil, err
		}

		isSlice = true
	}

	if self.match(token.TYPE) || self.match(token.IDENTIFIER) {
		kind := self.prev

		if !self.scope.Has(kind.String()) {
			return nil, self.error("type '" + kind.String() + "' not found")
		}

		_type = self.scope.Get(kind.String())

		if self.match(token.QUESTION_MARK) {
			nilable = self.prev
		}
	}

	if isSlice {
		_type = reflect.NewSliceType(_type, -1)
	}

	if self.match(token.EQ) {
		init, err = self.expression()

		if err != nil {
			return nil, err
		}

		t, err := init.GetType()

		if err != nil {
			return nil, err
		}

		if _type != nil && !_type.Equals(t) {
			return nil, self.error("expected type '" + _type.Name() + "', received '" + t.Name() + "'")
		}

		_type = t
	}

	_, err = self.consume(token.SEMI_COLON, "expected ';'")

	if err != nil {
		return nil, err
	}

	self.scope.Define(name.String(), _type)
	return stmt.NewVar(keyword, name, _type, nilable, init), nil
}

func (self *Parser) _struct() (stmt.Stmt, *error.Error) {
	name, err := self.consume(token.IDENTIFIER, "expected struct name")

	if err != nil {
		return nil, err
	}

	if self.scope.Has(name.String()) {
		return nil, self.error("duplicate name")
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

	v := stmt.NewStruct(name, methods)
	self.scope.Define(name.String(), nil)
	return v, nil
}

func (self *Parser) _for() (stmt.Stmt, *error.Error) {
	parent := self.scope
	self.scope = NewChildScope(parent)

	defer func() {
		self.scope = parent
	}()

	var init stmt.Stmt = nil
	var cond expr.Expr = nil
	var inc expr.Expr = nil

	_, err := self.consume(token.LEFT_PAREN, "expected '('")

	if err != nil {
		return nil, err
	}

	if self.match(token.LET) || self.match(token.CONST) {
		init, err = self._var()
	}

	if err != nil {
		return nil, err
	}

	cond, err = self.expression()

	if err != nil {
		return nil, err
	}

	if init != nil {
		_, err = self.consume(token.SEMI_COLON, "expected ';'")

		if err != nil {
			return nil, err
		}

		inc, err = self.expression()

		if err != nil {
			return nil, err
		}
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
		cond = expr.NewLiteral(reflect.NewBool(true))
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

	_, err = expr.GetType()

	if err != nil {
		return nil, err
	}

	return stmt.NewExpr(expr), nil
}

func (self *Parser) fn() (stmt.Stmt, *error.Error) {
	var return_type reflect.Type = reflect.NewNilType()
	vars := []*stmt.Var{}
	name, err := self.consume(token.IDENTIFIER, "expected function name")

	if err != nil {
		return nil, err
	}

	if self.scope.HasLocal(name.String()) {
		return nil, self.error("duplicate name")
	}

	_, err = self.consume(token.LEFT_PAREN, "expected '('")

	if err != nil {
		return nil, err
	}

	parent := self.scope
	self.scope = NewChildScope(parent)

	if self.curr.Kind != token.RIGHT_PAREN {
		for {
			var nilable *token.Token = nil
			param, err := self.consume(token.IDENTIFIER, "expected parameter name")

			if err != nil {
				return nil, err
			}

			_type, err := self.consume(token.TYPE, "expected parameter type")

			if err != nil {
				return nil, err
			}

			if !self.scope.Has(_type.String()) {
				return nil, self.error("type '" + _type.String() + "' is undefined")
			}

			type_value := self.scope.Get(_type.String())

			if self.match(token.QUESTION_MARK) {
				nilable = self.prev
			}

			vars = append(vars, stmt.NewVar(
				nil,
				param,
				type_value,
				nilable,
				nil,
			))

			self.scope.Define(
				param.String(),
				self.scope.Get(_type.String()),
			)

			if !self.match(token.COMMA) {
				break
			}
		}
	}

	_, err = self.consume(token.RIGHT_PAREN, "expected ')'")

	if err != nil {
		return nil, err
	}

	if self.match(token.RETURN_TYPE) {
		t, err := self.consume(token.TYPE, "expected return type")

		if err != nil {
			return nil, err
		}

		if !self.scope.Has(t.String()) {
			return nil, self.error("type '" + t.String() + "' not found")
		}

		return_type = self.scope.Get(t.String())
	}

	fn := reflect.NewFnType(
		name.String(),
		utils.MapSlice(vars, func(v *stmt.Var) reflect.Param {
			return reflect.Param{
				Name: v.Name.String(),
				Type: v.Type,
			}
		}),
		return_type,
	)

	self.scope.Define(name.String(), fn)
	_, err = self.consume(token.LEFT_BRACE, "expected '{'")

	if err != nil {
		return nil, err
	}

	body, err := self.block()

	if err != nil {
		return nil, err
	}

	for _, line := range body {
		if ret, ok := line.(*stmt.Return); ok {
			t, err := ret.Value.GetType()

			if err != nil {
				return nil, err
			}

			if !return_type.Equals(t) {
				return nil, self.error(fmt.Sprintf(
					"expected return type '%s', received '%s'",
					return_type.Name(),
					t.Name(),
				))
			}
		}
	}

	v := stmt.NewFn(
		name,
		vars,
		return_type,
		body,
	)

	self.scope = parent
	self.scope.Define(name.String(), fn)

	return v, nil
}

func (self *Parser) block() ([]stmt.Stmt, *error.Error) {
	parent := self.scope
	self.scope = NewChildScope(parent)

	defer func() {
		self.scope = parent
	}()

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

func (self *Parser) use() (stmt.Stmt, *error.Error) {
	path := []*token.Token{}

	for {
		var name *token.Token = nil

		if self.match(token.STAR) {
			name = self.prev
		}

		if name == nil {
			n, err := self.consume(token.IDENTIFIER, "expected identifier")

			if err != nil {
				return nil, err
			}

			name = n
		}

		path = append(path, name)

		if name.Kind == token.STAR || !self.match(token.DOUBLE_COLON) {
			break
		}
	}

	_, err := self.consume(token.SEMI_COLON, "expected ';'")

	if err != nil {
		return nil, err
	}

	filePath := filepath.Dir(self.path) + "/"

	for i, n := range path {
		if n.Kind == token.IDENTIFIER {
			filePath += n.String()

			if i < len(path)-1 && path[i+1].Kind != token.STAR {
				filePath += "/"
			}
		}
	}

	if _, err := os.Stat(fmt.Sprintf("%s.q", filePath)); err == nil {
		filePath = fmt.Sprintf("%s.q", filePath)
	} else if _, err := os.Stat(fmt.Sprintf("%s/mod.q", filePath)); err == nil {
		filePath = fmt.Sprintf("%s/mod.q", filePath)
	} else {
		return nil, self.error("module not found")
	}

	parser, ok := cache[filePath]
	stmts := []stmt.Stmt{}

	if ok {
		stmts = parser.stmts
	} else {
		parser = New(filePath)
		_stmts, errs := parser.Parse()

		if errs != nil && len(errs) > 0 {
			self.errs = append(self.errs, errs...)
		}

		stmts = _stmts
		cache[filePath] = parser
	}

	v := stmt.NewUse(path, stmts)

	if path[len(path)-1].Kind == token.STAR {
		for key, _type := range parser.scope.types {
			self.scope.Define(key, _type)
		}
	} else {
		mod := reflect.NewModType()

		for key, _type := range parser.scope.types {
			mod.SetExport(key, _type)
		}

		self.scope.Define(path[len(path)-1].String(), mod)
	}

	return v, nil
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

	varType, err := e.GetType()

	if err != nil {
		return nil, err
	}

	if self.match(token.EQ) {
		value, err := self.assignment()

		if err != nil {
			return nil, err
		}

		assignType, err := value.GetType()

		if err != nil {
			return nil, err
		}

		if callable, ok := assignType.(reflect.CallableType); ok {
			assignType = callable.ReturnType()
		}

		if !varType.Equals(assignType) {
			return nil, self.error("expected type '" + varType.Name() + "', received '" + assignType.Name() + "'")
		}

		switch e.(type) {
		case *expr.Var:
			_var := e.(*expr.Var)

			if !self.scope.Has(_var.Name.String()) {
				return nil, self.error("undefined identifier")
			}

			return expr.NewAssign(_var.Name, value), nil
		case *expr.Get:
			get := e.(*expr.Get)

			if !self.scope.Has(get.Name.String()) {
				return nil, self.error("undefined identifier")
			}

			return expr.NewSet(get.Object, get.Name, value), nil
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
	var _type reflect.Type = nil
	e, err := self.primary()

	if err != nil {
		return nil, err
	}

	for {
		if self.match(token.LEFT_PAREN) {
			args := []expr.Expr{}

			switch e.(type) {
			case *expr.Var:
				v := e.(*expr.Var)
				_type = self.scope.Get(v.Name.String())
			case *expr.Get:
				v := e.(*expr.Get)
				t, err := v.GetType()

				if err != nil {
					return nil, err
				}

				_type = t
			}

			fn, ok := _type.(reflect.CallableType)

			if !ok {
				return nil, self.error("expected type 'fn', received '" + _type.Name() + "'")
			}

			if self.curr.Kind != token.RIGHT_PAREN {
				i := 0

				for {
					arg, err := self.expression()

					if err != nil {
						return nil, err
					}

					args = append(args, arg)

					if i > len(fn.Params())-1 {
						return nil, self.error("too many arguments")
					}

					t, err := arg.GetType()

					if err != nil {
						return nil, err
					}

					if !fn.Params()[i].Type.Equals(t) {
						return nil, self.error(fmt.Sprintf(
							"expected type '%s', received '%s'",
							fn.Params()[i].Type.Name(),
							t.Name(),
						))
					}

					if !self.match(token.COMMA) {
						break
					}

					i++
				}
			}

			paren, err := self.consume(token.RIGHT_PAREN, "expected ')'")

			if err != nil {
				return nil, err
			}

			if len(args) != len(fn.Params()) {
				return nil, self.error(fmt.Sprintf(
					"expected %d arguments, received %d",
					len(fn.Params()),
					len(args),
				))
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

		return expr.NewLiteral(reflect.NewBool(v)), nil
	} else if self.match(token.LINT) {
		v, err := self.prev.Int()

		if err != nil {
			return nil, self.error(err.Error())
		}

		return expr.NewLiteral(reflect.NewInt(v)), nil
	} else if self.match(token.LFLOAT) {
		v, err := self.prev.Float()

		if err != nil {
			return nil, self.error(err.Error())
		}

		return expr.NewLiteral(reflect.NewFloat(v)), nil
	} else if self.match(token.LSTRING) {
		return expr.NewLiteral(reflect.NewString(self.prev.String())), nil
	} else if self.match(token.LBYTE) {
		return expr.NewLiteral(reflect.NewByte(self.prev.Byte())), nil
	} else if self.match(token.NIL) {
		return expr.NewLiteral(reflect.NewNil()), nil
	} else if self.match(token.SELF) {
		if !self.scope.Has("self") {
			return nil, self.error("self is undefined")
		}

		return expr.NewSelf(self.prev, self.scope.Get("self")), nil
	} else if self.match(token.IDENTIFIER) {
		if !self.scope.Has(self.prev.String()) {
			return nil, self.error("undefined identifier '" + self.prev.String() + "'")
		}

		return expr.NewVar(self.prev, self.scope.Get(self.prev.String())), nil
	} else if self.match(token.LEFT_PAREN) {
		e, err := self.expression()

		if err != nil {
			return nil, err
		}

		self.consume(token.RIGHT_PAREN, "expected ')'")
		return expr.NewGrouping(e), nil
	} else if self.match(token.LEFT_BRACKET) {
		var _type reflect.Type = nil
		items := []expr.Expr{}

		if !self.match(token.RIGHT_BRACKET) {
			for {
				e, err := self.expression()

				if err != nil {
					return nil, err
				}

				t, err := e.GetType()

				if err != nil {
					return nil, err
				}

				if _type == nil {
					_type = t
				} else if !_type.Equals(t) {
					return nil, self.error(fmt.Sprintf(
						"expected type '%s', received '%s'",
						_type.Name(),
						t.Name(),
					))
				}

				items = append(items, e)

				if !self.match(token.COMMA) {
					break
				}
			}
		}

		_, err := self.consume(token.RIGHT_BRACKET, "expected ']'")

		if err != nil {
			return nil, err
		}

		return expr.NewSlice(_type, items), nil
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
		case token.USE:
		case token.STRUCT:
		case token.FN:
		case token.LET:
		case token.CONST:
		case token.FOR:
		case token.IF:
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
