package reflect

type Type interface {
	Kind() Kind
	Name() string
	String() string
	Comparable() bool
	Numeric() bool
	Collection() bool
	Len() int
}

type ComparableType interface {
	Type
	comparable
}
