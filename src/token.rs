#[allow(dead_code)]
#[derive(Debug)]
pub enum Kind {
    Comma,
    Dot,
    SemiColon,
    LParen,
    RParen,
    LBrace,
    RBrace,

    // arithmetic
    Minus,
    MinusEq,
    Plus,
    PlusEq,
    Slash,
    SlashEq,
    Star,
    StarEq,

    // logical
    Not,
    NotEq,
    Eq,
    EqEq,
    Gt,
    GtEq,
    Lt,
    LtEq,
    And,
    Or,

    // literals
    Identifier,
    String,
    Number,
    Boolean,
    Nil,

    // keywords
    If,
    Else,
    For,
    Let,
    Fn,
    Return,
    Class,
    Super,
    This,
    Pub,
    Use,

    Eof,
}

impl Kind {
    pub fn from_keyword(keyword: &str) -> Option<Self> {
        return match keyword {
            "if" => Some(Self::If),
            "else" => Some(Self::Else),
            "for" => Some(Self::For),
            "let" => Some(Self::Let),
            "fn" => Some(Self::Fn),
            "return" => Some(Self::Return),
            "class" => Some(Self::Class),
            "super" => Some(Self::Super),
            "this" => Some(Self::This),
            "pub" => Some(Self::Pub),
            "use" => Some(Self::Use),
            _ => None,
        };
    }

    #[allow(dead_code)]
    pub fn as_str(&self) -> &'static str {
        return match self {
            Self::Comma => ",",
            Self::Dot => ".",
            Self::SemiColon => ";",
            Self::LParen => "(",
            Self::RParen => ")",
            Self::LBrace => "{",
            Self::RBrace => "}",
            Self::Minus => "-",
            Self::MinusEq => "-=",
            Self::Plus => "+",
            Self::PlusEq => "+=",
            Self::Slash => "/",
            Self::SlashEq => "/=",
            Self::Star => "*",
            Self::StarEq => "*=",
            Self::Not => "!",
            Self::NotEq => "!=",
            Self::Eq => "=",
            Self::EqEq => "==",
            Self::Gt => ">",
            Self::GtEq => ">=",
            Self::Lt => "<",
            Self::LtEq => "<=",
            Self::And => "&&",
            Self::Or => "||",
            Self::String => "string",
            Self::Number => "number",
            Self::Boolean => "boolean",
            Self::Nil => "nil",
            Self::If => "if",
            Self::Else => "else",
            Self::For => "for",
            Self::Let => "let",
            Self::Fn => "fn",
            Self::Return => "return",
            Self::Class => "class",
            Self::Super => "super",
            Self::This => "this",
            Self::Pub => "pub",
            Self::Use => "use",
            Self::Eof => "\0",
            _ => "",
        };
    }
}

#[derive(Debug)]
pub struct Token {
    pub kind: Kind,
    pub ln: usize,
    pub start: usize,
    pub end: usize,
    pub value: Vec<u8>,
}

impl Token {
    pub fn new(kind: Kind, ln: usize, start: usize, end: usize, value: Vec<u8>) -> Self {
        return Self {
            kind,
            ln,
            start,
            end,
            value,
        };
    }
}

impl std::fmt::Display for Token {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        return write!(
            f,
            "{} => {:?}",
            self.kind.as_str(),
            std::str::from_utf8(&self.value).unwrap()
        );
    }
}
