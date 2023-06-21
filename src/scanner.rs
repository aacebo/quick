use crate::error;
use crate::token;

#[derive(Debug)]
pub struct Scanner {
    pub tokens: Vec<token::Token>,
    pub errors: Vec<error::Error>,

    _src: Vec<u8>,
    _left: usize,
    _right: usize,
    _ln: usize,
}

impl Scanner {
    pub fn new(src: Vec<u8>) -> Self {
        let mut v = Self {
            tokens: Vec::new(),
            errors: Vec::new(),
            _src: src,
            _left: 0,
            _right: 0,
            _ln: 0,
        };

        while v._right < v._src.len() {
            v._left = v._right;
            let c = v._src[v._right] as char;
            v._right += 1;

            match c {
                // ignore whitespace
                ' ' | '\r' | '\t' => {}
                '\n' => v._ln += 1,
                '(' => v.push(token::Kind::LParen),
                ')' => v.push(token::Kind::RParen),
                '{' => v.push(token::Kind::LBrace),
                '}' => v.push(token::Kind::RBrace),
                ',' => v.push(token::Kind::Comma),
                '.' => v.push(token::Kind::Dot),
                ';' => v.push(token::Kind::SemiColon),
                '/' => {
                    if v.peek() == '/' {
                        v.on_comment();
                    } else if v.peek() == '=' {
                        v.push(token::Kind::SlashEq);
                        v._right += 1;
                    } else {
                        v.push(token::Kind::Slash);
                    }
                }
                '"' => v.on_string(),
                _ => {
                    if c.is_numeric() {
                        v.on_number();
                    } else if c.is_alphabetic() {
                        v.on_identifier();
                    } else {
                        v.errors.push(error::Error::new(
                            error::Kind::UnexpectedChar,
                            v._ln,
                            v._left,
                            v._right,
                        ));
                    }
                }
            };
        }

        return v;
    }

    fn peek(&self) -> char {
        if self._right >= self._src.len() {
            return '\0';
        }

        return self._src[self._right] as char;
    }

    fn push(&mut self, kind: token::Kind) {
        self.tokens.push(token::Token::new(
            kind,
            self._ln,
            self._left,
            self._right,
            self._src[self._left..self._right].to_vec(),
        ));
    }

    fn on_comment(&mut self) {
        while self.peek() != '\n' && self.peek() != '\0' {
            self._right += 1;
        }

        self._ln += 1;
        self._right += 1;
    }

    fn on_string(&mut self) {
        while self.peek() != '"' && self.peek() != '\0' {
            if self.peek() == '\n' {
                self._ln += 1;
            }

            self._right += 1;
        }

        if self._right == self._src.len() {
            self.errors.push(error::Error::new(
                error::Kind::UnterminatedString,
                self._ln,
                self._left,
                self._right,
            ));

            return;
        }

        self._right += 1;
        self.push(token::Kind::String);
    }

    fn on_number(&mut self) {
        while self.peek().is_numeric() {
            self._right += 1;
        }

        if self.peek() == '.' {
            self._right += 1;

            while self.peek().is_numeric() {
                self._right += 1;
            }
        }

        self.push(token::Kind::Number);
    }

    fn on_identifier(&mut self) {
        while self.peek().is_alphanumeric() {
            self._right += 1;
        }

        let identifier = match std::str::from_utf8(&self._src[self._left..self._right]) {
            Ok(v) => v,
            Err(e) => panic!("{}", e),
        };

        match token::Kind::from_keyword(identifier) {
            Some(kind) => self.push(kind),
            None => self.push(token::Kind::Identifier),
        };
    }
}
