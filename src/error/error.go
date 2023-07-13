package error

type Error struct {
	Ln      int
	Start   int
	End     int
	Message string
}

func New(ln int, start int, end int, message string) *Error {
	return &Error{
		ln,
		start,
		end,
		message,
	}
}
