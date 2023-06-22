use super::Expression;

pub struct Grouping<'a> {
    pub exp: &'a dyn Expression,
}

impl<'a> Grouping<'a> {
    pub fn new(exp: &'a dyn Expression) -> Self {
        return Self { exp };
    }
}
