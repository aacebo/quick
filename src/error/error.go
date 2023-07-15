package error

import "fmt"

type Error struct {
	Path    string
	Ln      int
	Start   int
	End     int
	Message string
}

func New(path string, ln int, start int, end int, message string) *Error {
	return &Error{
		Path:    path,
		Ln:      ln,
		Start:   start,
		End:     end,
		Message: message,
	}
}

func (self Error) String() string {
	return fmt.Sprintf(
		"[%s] [ln: %d, start: %d, end: %d] => %s",
		self.Path,
		self.Ln,
		self.Start,
		self.End,
		self.Message,
	)
}
