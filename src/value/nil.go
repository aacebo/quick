package value

type Nil struct{}

func (self Nil) Kind() Kind {
	return NIL
}

func (self Nil) String() string {
	return "<nil>"
}
