package value

type Numeric interface {
	Gt(Comparable) Bool
	GtEq(Comparable) Bool
	Lt(Comparable) Bool
	LtEq(Comparable) Bool

	Add(Comparable) Comparable
	Subtract(Comparable) Comparable
	Multiply(Comparable) Comparable
	Divide(Comparable) Comparable

	Inc()
	Dec()

	Comparable
}
