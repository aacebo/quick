package reflect

type NilType struct{}

func (self NilType) Kind() Kind {
	return Nil
}

func (self NilType) Name() string {
	return Nil.String()
}

func (self NilType) String() string {
	return Nil.String()
}

func (self NilType) Len() int {
	panic("method not supported")
}

func (self NilType) Comparable() bool {
	return true
}

func (self NilType) Numeric() bool {
	return false
}

func (self NilType) Collection() bool {
	return false
}

func (self NilType) Equals(t Type) bool {
	return t.Kind() == Nil
}
