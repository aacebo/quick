use super::Expression;
use crate::token;

pub struct Unary<'a> {
    pub op: token::Token,
    pub right: &'a dyn Expression,
}

impl<'a> Unary<'a> {
    pub fn new(op: token::Token, right: &'a dyn Expression) -> Self {
        return Self { op, right };
    }
}
