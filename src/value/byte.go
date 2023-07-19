package value

type Byte byte

func (self Byte) Kind() Kind {
	return BYTE
}

func (self Byte) ComplexKind() ComplexKind {
	return NONE
}

func (self Byte) Truthy() Bool {
	return self > 0
}

func (self Byte) Name() string {
	return "byte"
}

func (self Byte) String() string {
	return string(self)
}

func (self Byte) TypeEq(other Value) bool {
	return other.Kind() == BYTE
}

func (self Byte) Eq(other Comparable) Bool {
	return self == other.(Comparable)
}

func (self Byte) Gt(other Numeric) Bool {
	return self > other.(Byte)
}

func (self Byte) GtEq(other Numeric) Bool {
	return self >= other.(Byte)
}

func (self Byte) Lt(other Numeric) Bool {
	return self < other.(Byte)
}

func (self Byte) LtEq(other Numeric) Bool {
	return self <= other.(Byte)
}

func (self Byte) Add(other Numeric) Numeric {
	return self + other.(Byte)
}

func (self Byte) Subtract(other Numeric) Numeric {
	return self - other.(Byte)
}

func (self Byte) Multiply(other Numeric) Numeric {
	return self * other.(Byte)
}

func (self Byte) Divide(other Numeric) Numeric {
	return self / other.(Byte)
}

func (self Byte) Inc() {
	self++
}

func (self Byte) Dec() {
	self--
}
