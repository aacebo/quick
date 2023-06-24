use crate::token;
use crate::error;
use crate::expression;
use crate::statement;

pub struct Parser {
    _tokens: Vec<token::Token>,
    _errors: Vec<error::Error>,
    _it: usize,
}

impl Parser {
    pub fn new(tokens: Vec<token::Token>) -> Self {
        return Self {
            _tokens: tokens,
            _errors: Vec::new(),
            _it: 0,
        };
    }

    // fn assignment(&mut self) -> expression::Expression {
    //     let expr = 
    // }

    // fn primary(&mut self) -> expression::Expression {
    //     if self.next_until_not(&[token::Kind::False]) {
    //         return expression::Expression::Literal(Box::new(false));
    //     } else if self.next_until_not(&[token::Kind::True]) {
    //         return expression::Expression::Literal(Box::new(true));
    //     } else if self.next_until_not(&[token::Kind::Nil]) {
    //         return expression::Expression::Literal(Box::new('\0'));
    //     } else if self.next_until_not(&[token::Kind::Number]) {
    //         return expression::Expression::Literal(Box::new(self.prev().as_number().unwrap()));
    //     } else if self.next_until_not(&[token::Kind::String]) {
    //         return expression::Expression::Literal(Box::new(self.prev().as_string().unwrap()));
    //     } else if self.next_until_not(&[token::Kind::Super]) {
    //         let keyword = self.prev().clone();
    //         self.consume(token::Kind::Dot, "expected '.' after 'super'");
    //         let method = self.consume(token::Kind::Identifier, "expected superclass method name").clone();
    //         return expression::Expression::Super(keyword, method);
    //     } else if self.next_until_not(&[token::Kind::This]) {
    //         return expression::Expression::This(self.prev().clone());
    //     } else if self.next_until_not(&[token::Kind::Identifier]) {
    //         return expression::Expression::Variable(self.prev().clone());
    //     } else if self.next_until_not(&[token::Kind::LParen]) {

    //         return expression::Expression::Grouping();
    //     }

    //     panic!("expected expression");
    // }

    fn term(&mut self) -> expression::Expression {
        let mut expr = self.factor();

        while self.next_until_not(&[token::Kind::Minus, token::Kind::Plus]) {
            expr = expression::Expression::Binary(
                Box::new(expr),
                self.prev().clone(),
                Box::new(self.factor()),
            );
        }

        return expr;
    }

    fn factor(&mut self) -> expression::Expression {
        let mut expr = self.unary();

        while self.next_until_not(&[token::Kind::Slash, token::Kind::Star]) {
            expr = expression::Expression::Binary(
                Box::new(expr),
                self.prev().clone(),
                Box::new(self.unary())
            );
        }

        return expr;
    }

    fn unary(&mut self) -> expression::Expression {
        if self.next_until_not(&[token::Kind::Not, token::Kind::Minus]) {
            return expression::Expression::Unary(
                self.prev().clone(),
                self.unary()
            );
        }

    }

    fn next_until_not(&mut self, kinds: &[token::Kind]) -> bool {
        for kind in kinds {
            if self.is_kind(*kind) {
                self.next();
                return true;
            }
        }

        return false;
    }

    fn consume(&mut self, kind: token::Kind, message: &str) -> &token::Token {
        if self.is_kind(kind) {
            return self.next();
        }

        panic!("{message}");
    }

    fn next(&mut self) -> &token::Token {
        if !self.is_end() {
            self._it += 1;
        }

        return self.prev();
    }

    fn is_end(&self) -> bool {
        return self.peek().kind == token::Kind::Eof;
    }

    fn is_kind(&self, kind: token::Kind) -> bool {
        return self.peek().kind == kind;
    }

    fn peek(&self) -> &token::Token {
        return &self._tokens[self._it];
    }

    fn prev(&self) -> &token::Token {
        return &self._tokens[self._it - 1];
    }
}
