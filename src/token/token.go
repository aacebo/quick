package token

import "strconv"

type Token struct {
	Kind  Kind
	Path  string
	Ln    int
	Start int
	End   int
	Value []byte
}

func New(kind Kind, path string, ln int, start int, end int, value []byte) *Token {
	return &Token{
		Kind:  kind,
		Path:  path,
		Ln:    ln,
		Start: start,
		End:   end,
		Value: value,
	}
}

func (self Token) Type() Kind {
	return Types[string(self.Value)]
}

func (self Token) Byte() byte {
	return self.Value[0]
}

func (self Token) String() string {
	return string(self.Value)
}

func (self Token) Int() (int, error) {
	return strconv.Atoi(string(self.Value))
}

func (self Token) Float() (float64, error) {
	return strconv.ParseFloat(string(self.Value), 64)
}

func (self Token) Bool() (bool, error) {
	return strconv.ParseBool(string(self.Value))
}
