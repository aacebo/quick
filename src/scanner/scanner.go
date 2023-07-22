package scanner

import (
	"quick/src/error"
	"quick/src/token"

	"golang.org/x/exp/slices"
)

type Scanner struct {
	path  string
	src   []byte
	ln    int
	left  int
	right int
}

func New(path string, src []byte) *Scanner {
	return &Scanner{
		path:  path,
		src:   src,
		ln:    0,
		left:  0,
		right: 0,
	}
}

func (self *Scanner) Next() (*token.Token, *error.Error) {
	if self.right >= len(self.src) {
		return self.create(token.EOF), nil
	}

	self.left = self.right
	b := self.src[self.right]
	self.right++

	switch b {
	case ' ':
	case '\r':
	case '\t':
		// ignore whitespace
		break
	case '\n':
		self.ln++
		break
	case '(':
		return self.create(token.LEFT_PAREN), nil
	case ')':
		return self.create(token.RIGHT_PAREN), nil
	case '{':
		return self.create(token.LEFT_BRACE), nil
	case '}':
		return self.create(token.RIGHT_BRACE), nil
	case '[':
		return self.create(token.LEFT_BRACKET), nil
	case ']':
		return self.create(token.RIGHT_BRACKET), nil
	case ',':
		return self.create(token.COMMA), nil
	case '.':
		return self.create(token.DOT), nil
	case ':':
		if self.peek() == ':' {
			self.right++
			return self.create(token.DOUBLE_COLON), nil
		}

		return self.create(token.COLON), nil
	case ';':
		return self.create(token.SEMI_COLON), nil
	case '?':
		return self.create(token.QUESTION_MARK), nil
	case '|':
		if self.peek() != '|' {
			return nil, self.error("expected '|'")
		}

		self.right++
		return self.create(token.OR), nil
	case '&':
		if self.peek() != '&' {
			return nil, self.error("expected '&'")
		}

		self.right++
		return self.create(token.AND), nil
	case '+':
		if self.peek() == '=' {
			self.right++
			return self.create(token.PLUS_EQ), nil
		}

		return self.create(token.PLUS), nil
	case '-':
		if self.peek() == '=' {
			self.right++
			return self.create(token.MINUS_EQ), nil
		} else if self.peek() == '>' {
			self.right++
			return self.create(token.RETURN_TYPE), nil
		} else if self.isInt(self.peek()) {
			self.right++
			return self.onNumeric()
		}

		return self.create(token.MINUS), nil
	case '*':
		if self.peek() == '=' {
			self.right++
			return self.create(token.STAR_EQ), nil
		}

		return self.create(token.STAR), nil
	case '/':
		if self.peek() == '/' {
			return self.onComment()
		} else if self.peek() == '=' {
			self.right++
			return self.create(token.SLASH_EQ), nil
		}

		return self.create(token.SLASH), nil
	case '!':
		if self.peek() == '=' {
			self.right++
			return self.create(token.NOT_EQ), nil
		}

		return self.create(token.NOT), nil
	case '=':
		if self.peek() == '=' {
			self.right++
			return self.create(token.EQ_EQ), nil
		}

		return self.create(token.EQ), nil
	case '>':
		if self.peek() == '=' {
			self.right++
			return self.create(token.GT_EQ), nil
		}

		return self.create(token.GT), nil
	case '<':
		if self.peek() == '=' {
			self.right++
			return self.create(token.LT_EQ), nil
		}

		return self.create(token.LT), nil
	case '\'':
		return self.onByte()
	case '"':
		return self.onString()
	default:
		if self.isInt(b) {
			return self.onNumeric()
		} else if self.isAlpha(b) {
			return self.onIdentifier()
		}

		return nil, self.error("unexpected character")
	}

	return self.Next()
}

func (self *Scanner) onComment() (*token.Token, *error.Error) {
	for self.peek() != '\n' && self.peek() != 0 {
		self.right++
	}

	self.ln++
	self.right++
	return self.Next()
}

func (self *Scanner) onByte() (*token.Token, *error.Error) {
	self.right++

	if self.peek() != '\'' {
		return nil, self.error("unterminated byte")
	}

	self.left++
	token := self.create(token.LBYTE)
	self.right++
	return token, nil
}

func (self *Scanner) onString() (*token.Token, *error.Error) {
	for self.peek() != '"' && self.peek() != 0 {
		if self.peek() == '\n' {
			self.ln++
		} else if self.peek() == '\\' {
			err := self.onEscape()

			if err != nil {
				return nil, err
			}
		}

		self.right++
	}

	if self.right == len(self.src) {
		return nil, self.error("unterminated string")
	}

	self.left++
	token := self.create(token.LSTRING)
	self.right++
	return token, nil
}

func (self *Scanner) onEscape() *error.Error {
	self.right++

	defer func() {
		self.right--
	}()

	switch self.peek() {
	case 'a': // bell
		self.src = slices.Replace(self.src, self.right-1, self.right+1, '\a')
	case 'b': // backspace
		self.src = slices.Replace(self.src, self.right-1, self.right+1, '\b')
	case 'f': // form feed
		self.src = slices.Replace(self.src, self.right-1, self.right+1, '\f')
	case 'n': // new line
		self.src = slices.Replace(self.src, self.right-1, self.right+1, '\n')
	case 'r': // carriage return
		self.src = slices.Replace(self.src, self.right-1, self.right+1, '\r')
	case 't': // horizontal tab
		self.src = slices.Replace(self.src, self.right-1, self.right+1, '\t')
	case 'v': // verical tab
		self.src = slices.Replace(self.src, self.right-1, self.right+1, '\v')
	case '\'': // single quote
		self.src = slices.Replace(self.src, self.right-1, self.right+1, '\'')
	case '"': // double quote
		self.src = slices.Replace(self.src, self.right-1, self.right+1, '"')
	case '\\': // back slash
		self.src = slices.Replace(self.src, self.right-1, self.right+1, '\\')
	default:
		return self.error("unknown escape sequence")
	}

	return nil
}

func (self *Scanner) onNumeric() (*token.Token, *error.Error) {
	kind := token.LINT

	for self.isInt(self.peek()) {
		self.right++
	}

	if self.peek() == '.' {
		kind = token.LFLOAT
		self.right++

		for self.isInt(self.peek()) {
			self.right++
		}
	}

	return self.create(kind), nil
}

func (self *Scanner) onIdentifier() (*token.Token, *error.Error) {
	for self.isAlpha(self.peek()) || self.isInt(self.peek()) {
		self.right++
	}

	name := self.src[self.left:self.right]

	if kind, ok := token.Keywords[string(name)]; ok {
		return self.create(kind), nil
	}

	return self.create(token.IDENTIFIER), nil
}

func (self Scanner) peek() byte {
	if self.right >= len(self.src) {
		return 0
	}

	return self.src[self.right]
}

func (self Scanner) isInt(b byte) bool {
	return b >= '0' && b <= '9'
}

func (self Scanner) isAlpha(b byte) bool {
	return (b >= 'a' && b <= 'z') ||
		(b >= 'A' && b <= 'Z') ||
		(b == '_')
}

func (self Scanner) create(kind token.Kind) *token.Token {
	return token.New(
		kind,
		self.path,
		self.ln,
		self.left,
		self.right,
		self.src[self.left:self.right],
	)
}

func (self Scanner) error(message string) *error.Error {
	return error.New(
		self.path,
		self.ln,
		self.left,
		self.right,
		message,
	)
}
