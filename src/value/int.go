package value

import (
	"fmt"
)

type Int int64

func (self Int) Kind() Kind {
	return INT
}

func (self Int) Truthy() Bool {
	return self > 0
}

func (self Int) String() string {
	return fmt.Sprintf("%d", self)
}

func (self Int) Eq(other Int) Bool {
	return self == other
}

func (self Int) Gt(other Int) Bool {
	return self > other
}

func (self Int) GtEq(other Int) Bool {
	return self >= other
}

func (self Int) Lt(other Int) Bool {
	return self < other
}

func (self Int) LtEq(other Int) Bool {
	return self <= other
}

func (self Int) Add(other Int) Int {
	return self + other
}

func (self Int) Subtract(other Int) Int {
	return self - other
}

func (self Int) Multiply(other Int) Int {
	return self * other
}

func (self Int) Divide(other Int) Int {
	return self / other
}

func (self Int) Inc() {
	self++
}

func (self Int) Dec() {
	self--
}
