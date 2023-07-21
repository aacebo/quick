package reflect

type NilType struct {
	members map[string]Type
}

func NewNilType() NilType {
	return NilType{
		members: map[string]Type{},
	}
}

func (self NilType) Kind() Kind {
	return Nil
}

func (self NilType) Name() string {
	return Nil.String()
}

func (self NilType) String() string {
	return Nil.String()
}

func (self NilType) Len() int {
	panic("method not supported")
}

func (self NilType) Comparable() bool {
	return true
}

func (self NilType) Numeric() bool {
	return false
}

func (self NilType) Collection() bool {
	return false
}

func (self NilType) Equals(t Type) bool {
	return t.Kind() == Nil
}

func (self NilType) HasMember(name string) bool {
	_, ok := self.members[name]
	return ok
}

func (self NilType) GetMember(name string) Type {
	return self.members[name]
}
