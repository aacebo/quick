use crate::token;
use std::any;

pub enum Expression {
    Assign(token::Token, Box<Expression>),
    Binary(Box<Expression>, token::Token, Box<Expression>),
    Call(Box<Expression>, token::Token, Vec<Expression>),
    Get(Box<Expression>, token::Token),
    Grouping(Box<Expression>),
    Literal(Box<dyn any::Any>),
    Logical(Box<Expression>, token::Token, Box<Expression>),
    Set(Box<Expression>, token::Token, Box<Expression>),
    Super(token::Token, token::Token),
    This(token::Token),
    Unary(token::Token, Box<Expression>),
    Variable(token::Token),
}
