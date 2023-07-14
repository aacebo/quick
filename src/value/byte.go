package value

type Byte byte

func (self Byte) Kind() Kind {
	return BYTE
}

func (self Byte) Truthy() Bool {
	return self > 0
}

func (self Byte) String() string {
	return string(self)
}

func (self Byte) Eq(other Byte) Bool {
	return self == other
}

func (self Byte) Gt(other Byte) Bool {
	return self > other
}

func (self Byte) GtEq(other Byte) Bool {
	return self >= other
}

func (self Byte) Lt(other Byte) Bool {
	return self < other
}

func (self Byte) LtEq(other Byte) Bool {
	return self <= other
}

func (self Byte) Add(other Byte) Byte {
	return self + other
}

func (self Byte) Subtract(other Byte) Byte {
	return self - other
}

func (self Byte) Multiply(other Byte) Byte {
	return self * other
}

func (self Byte) Divide(other Byte) Byte {
	return self / other
}

func (self Byte) Inc() {
	self++
}

func (self Byte) Dec() {
	self--
}
