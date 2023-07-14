package value

type Map[K Comparable, V Value] map[K]V

func (self Map[K, V]) Kind() Kind {
	return COMPLEX
}

func (self Map[K, V]) ComplexKind() ComplexKind {
	return MAP
}

func (self Map[K, V]) String() string {
	str := "{"

	for k, v := range self {
		str += (k.String() + ": " + v.String() + ", ")
	}

	return str + "}"
}
