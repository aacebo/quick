package value

import (
	"fmt"
)

type Int int64

func (self Int) Kind() Kind {
	return INT
}

func (self Int) ComplexKind() ComplexKind {
	return NONE
}

func (self Int) Truthy() Bool {
	return self > 0
}

func (self Int) Name() string {
	return "int"
}

func (self Int) String() string {
	return fmt.Sprintf("%d", self)
}

func (self Int) TypeEq(other Value) bool {
	return other.Kind() == INT
}

func (self Int) Eq(other Comparable) Bool {
	return self == other.(Int)
}

func (self Int) Gt(other Numeric) Bool {
	return self > other.(Int)
}

func (self Int) GtEq(other Numeric) Bool {
	return self >= other.(Int)
}

func (self Int) Lt(other Numeric) Bool {
	return self < other.(Int)
}

func (self Int) LtEq(other Numeric) Bool {
	return self <= other.(Int)
}

func (self Int) Add(other Numeric) Numeric {
	return self + other.(Int)
}

func (self Int) Subtract(other Numeric) Numeric {
	return self - other.(Int)
}

func (self Int) Multiply(other Numeric) Numeric {
	return self * other.(Int)
}

func (self Int) Divide(other Numeric) Numeric {
	return self / other.(Int)
}

func (self Int) Inc() {
	self++
}

func (self Int) Dec() {
	self--
}
