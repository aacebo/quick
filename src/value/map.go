package value

type Map[K KeyComparable, V Value] map[K]V

func (self Map[K, V]) Kind() Kind {
	return COMPLEX
}

func (self Map[K, V]) ComplexKind() ComplexKind {
	return MAP
}

func (self Map[K, V]) Truthy() Bool {
	return self != nil
}

func (self Map[K, V]) Name() string {
	return "map"
}

func (self Map[K, V]) String() string {
	str := "{"

	for k, v := range self {
		str += (k.String() + ": " + v.String() + ", ")
	}

	return str + "}"
}
