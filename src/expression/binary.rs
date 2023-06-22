use super::Expression;
use crate::token;

pub struct Binary<'a> {
    pub left: &'a dyn Expression,
    pub op: token::Token,
    pub right: &'a dyn Expression,
}

impl<'a> Binary<'a> {
    pub fn new(left: &'a dyn Expression, op: token::Token, right: &'a dyn Expression) -> Self {
        return Self { left, op, right };
    }
}
