package token

type Token struct {
	Kind  Kind
	Ln    int
	Start int
	End   int
	Value []byte
}

func New(kind Kind, ln int, start int, end int, value []byte) *Token {
	return &Token{
		kind,
		ln,
		start,
		end,
		value,
	}
}

func (self Token) ToString() string {
	return string(self.Value)
}
