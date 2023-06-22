use crate::expression;
use crate::token;

pub enum Statement {
    Block(Vec<Statement>),
    Class(token::Token, expression::Expression, Vec<Statement>),
    Expression(expression::Expression),
    Function(token::Token, Vec<token::Token>, Vec<Statement>),
    If(expression::Expression, Box<Statement>, Box<Statement>),
    Return(token::Token, expression::Expression),
    Let(token::Token, expression::Expression),
    For(expression::Expression, Box<Statement>),
}
