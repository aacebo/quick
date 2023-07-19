package value

type Slice[T Value] struct {
	Of     T
	Values []T
}

func NewSlice[T Value](of T, values []T) *Slice[T] {
	return &Slice[T]{
		Of:     of,
		Values: values,
	}
}

func (self Slice[T]) Kind() Kind {
	return COMPLEX
}

func (self Slice[T]) ComplexKind() ComplexKind {
	return SLICE
}

func (self Slice[T]) Truthy() Bool {
	return len(self.Values) > 0
}

func (self Slice[T]) Name() string {
	return "[]" + self.Of.Name()
}

func (self Slice[T]) String() string {
	str := "["

	for i, v := range self.Values {
		str += v.String()

		if i < len(self.Values)-1 {
			str += ", "
		}
	}

	return str + "]"
}

func (self Slice[T]) TypeEq(other Value) bool {
	if other.Kind() != COMPLEX || other.ComplexKind() != SLICE {
		return false
	}

	_, ok := other.(*Slice[T])
	return ok
}
