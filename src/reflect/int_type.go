package reflect

type IntType struct {
	members map[string]Type
}

func NewIntType() IntType {
	return IntType{
		members: map[string]Type{},
	}
}

func (self IntType) Kind() Kind {
	return Int
}

func (self IntType) Name() string {
	return Int.String()
}

func (self IntType) String() string {
	return Int.String()
}

func (self IntType) Len() int {
	panic("method not supported")
}

func (self IntType) Comparable() bool {
	return true
}

func (self IntType) Numeric() bool {
	return true
}

func (self IntType) Collection() bool {
	return false
}

func (self IntType) Equals(t Type) bool {
	return t.Kind() == Int
}

func (self IntType) HasMember(name string) bool {
	_, ok := self.members[name]
	return ok
}

func (self IntType) GetMember(name string) Type {
	return self.members[name]
}
