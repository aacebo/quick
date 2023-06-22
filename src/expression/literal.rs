pub struct Literal<T> {
    pub value: T,
}

impl<T> Literal<T> {
    pub fn new(value: T) -> Self {
        return Self { value };
    }
}
