package value

import (
	"fmt"
)

type Float float64

func (self Float) Kind() Kind {
	return FLOAT
}

func (self Float) Truthy() Bool {
	return self > 0
}

func (self Float) String() string {
	return fmt.Sprintf("%f", self)
}

func (self Float) Eq(other Float) Bool {
	return self == other
}

func (self Float) Gt(other Float) Bool {
	return self > other
}

func (self Float) GtEq(other Float) Bool {
	return self >= other
}

func (self Float) Lt(other Float) Bool {
	return self < other
}

func (self Float) LtEq(other Float) Bool {
	return self <= other
}

func (self Float) Add(other Float) Float {
	return self + other
}

func (self Float) Subtract(other Float) Float {
	return self - other
}

func (self Float) Multiply(other Float) Float {
	return self * other
}

func (self Float) Divide(other Float) Float {
	return self / other
}

func (self Float) Inc() {
	self++
}

func (self Float) Dec() {
	self--
}
