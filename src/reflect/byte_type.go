package reflect

type ByteType struct {
	members map[string]Type
}

func NewByteType() ByteType {
	return ByteType{
		members: map[string]Type{},
	}
}

func (self ByteType) Kind() Kind {
	return Byte
}

func (self ByteType) Name() string {
	return Byte.String()
}

func (self ByteType) String() string {
	return Byte.String()
}

func (self ByteType) Len() int {
	panic("method not supported")
}

func (self ByteType) Comparable() bool {
	return true
}

func (self ByteType) Numeric() bool {
	return true
}

func (self ByteType) Collection() bool {
	return false
}

func (self ByteType) Equals(t Type) bool {
	return t.Kind() == Byte
}

func (self ByteType) HasMember(name string) bool {
	_, ok := self.members[name]
	return ok
}

func (self ByteType) GetMember(name string) Type {
	return self.members[name]
}
