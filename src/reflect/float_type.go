package reflect

type FloatType struct{}

func (self FloatType) Kind() Kind {
	return Float
}

func (self FloatType) Name() string {
	return Float.String()
}

func (self FloatType) String() string {
	return Float.String()
}

func (self FloatType) Len() int {
	panic("method not supported")
}

func (self FloatType) Comparable() bool {
	return true
}

func (self FloatType) Numeric() bool {
	return true
}

func (self FloatType) Collection() bool {
	return false
}

func (self FloatType) Equals(t Type) bool {
	return t.Kind() == Float
}
