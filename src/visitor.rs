use crate::expression;
use crate::statement;

pub trait Visitor<T> {
    fn visit_statement(&mut self, s: &statement::Statement) -> T;
    fn visit_expression(&mut self, e: &expression::Expression) -> T;
}
