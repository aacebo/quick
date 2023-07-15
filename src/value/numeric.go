package value

type Numeric interface {
	Gt(Numeric) Bool
	GtEq(Numeric) Bool
	Lt(Numeric) Bool
	LtEq(Numeric) Bool

	Add(Numeric) Numeric
	Subtract(Numeric) Numeric
	Multiply(Numeric) Numeric
	Divide(Numeric) Numeric

	Inc()
	Dec()

	Comparable
}
