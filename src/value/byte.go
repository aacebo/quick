package value

type Byte byte

func (self Byte) Kind() Kind {
	return BYTE
}

func (self Byte) String() string {
	return string(self)
}
