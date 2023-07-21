package reflect

type FloatType struct {
	members map[string]Type
}

func NewFloatType() FloatType {
	return FloatType{
		members: map[string]Type{},
	}
}

func (self FloatType) Kind() Kind {
	return Float
}

func (self FloatType) Name() string {
	return Float.String()
}

func (self FloatType) String() string {
	return Float.String()
}

func (self FloatType) Len() int {
	panic("method not supported")
}

func (self FloatType) Comparable() bool {
	return true
}

func (self FloatType) Numeric() bool {
	return true
}

func (self FloatType) Collection() bool {
	return false
}

func (self FloatType) Equals(t Type) bool {
	return t.Kind() == Float
}

func (self FloatType) HasMember(name string) bool {
	_, ok := self.members[name]
	return ok
}

func (self FloatType) GetMember(name string) Type {
	return self.members[name]
}
