package value

type Slice[T Value] []T

func (self Slice[T]) Kind() Kind {
	return COMPLEX
}

func (self Slice[T]) ComplexKind() ComplexKind {
	return SLICE
}

func (self Slice[T]) Truthy() Bool {
	return self != nil
}

func (self Slice[T]) Name() string {
	return "slice"
}

func (self Slice[T]) String() string {
	str := "["

	for i, v := range self {
		str += v.String()

		if i < len(self)-1 {
			str += ", "
		}
	}

	return str + "]"
}
