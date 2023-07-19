package value

import (
	"fmt"
)

type Float float64

func (self Float) Kind() Kind {
	return FLOAT
}

func (self Float) ComplexKind() ComplexKind {
	return NONE
}

func (self Float) Truthy() Bool {
	return self > 0
}

func (self Float) Name() string {
	return "float"
}

func (self Float) String() string {
	return fmt.Sprintf("%f", self)
}

func (self Float) TypeEq(other Value) bool {
	return other.Kind() == FLOAT
}

func (self Float) Eq(other Comparable) Bool {
	return self == other.(Comparable)
}

func (self Float) Gt(other Numeric) Bool {
	return self > other.(Float)
}

func (self Float) GtEq(other Numeric) Bool {
	return self >= other.(Float)
}

func (self Float) Lt(other Numeric) Bool {
	return self < other.(Float)
}

func (self Float) LtEq(other Numeric) Bool {
	return self <= other.(Float)
}

func (self Float) Add(other Numeric) Numeric {
	return self + other.(Float)
}

func (self Float) Subtract(other Numeric) Numeric {
	return self - other.(Float)
}

func (self Float) Multiply(other Numeric) Numeric {
	return self * other.(Float)
}

func (self Float) Divide(other Numeric) Numeric {
	return self / other.(Float)
}

func (self Float) Inc() {
	self++
}

func (self Float) Dec() {
	self--
}
