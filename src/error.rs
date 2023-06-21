#[derive(Debug)]
pub enum Kind {
    UnexpectedChar,
    UnterminatedString,
}

impl Kind {
    fn to_str(&self) -> &str {
        return match self {
            Self::UnexpectedChar => "unexpected character",
            Self::UnterminatedString => "unterminated string",
        };
    }
}

#[derive(Debug)]
pub struct Error {
    pub kind: Kind,
    pub ln: usize,
    pub start: usize,
    pub end: usize,
}

impl Error {
    pub fn new(kind: Kind, ln: usize, start: usize, end: usize) -> Self {
        return Self {
            kind,
            ln,
            start,
            end,
        };
    }
}

impl std::fmt::Display for Error {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        return write!(
            f,
            "[ln: {}, start: {}, end: {}] - {}",
            self.ln,
            self.start,
            self.end,
            self.kind.to_str()
        );
    }
}
