package value

type Comparable interface {
	Eq(Comparable) Bool

	Value
}

type KeyComparable interface {
	comparable
	Comparable
}
