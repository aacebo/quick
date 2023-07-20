package reflect

type Type interface {
	Kind() Kind
	Name() string
	String() string
	Len() int
	Comparable() bool
	Numeric() bool
	Collection() bool
	Equals(Type) bool
}

type ComparableType interface {
	Type
	comparable
}
