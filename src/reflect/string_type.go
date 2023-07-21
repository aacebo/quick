package reflect

type StringType struct {
	members map[string]Type
}

func NewStringType() StringType {
	return StringType{
		members: map[string]Type{
			"len": NewNativeFnType(
				"len",
				[]Param{},
				NewIntType(),
			),
		},
	}
}

func (self StringType) Kind() Kind {
	return String
}

func (self StringType) Name() string {
	return String.String()
}

func (self StringType) String() string {
	return String.String()
}

func (self StringType) Len() int {
	panic("method not supported")
}

func (self StringType) Comparable() bool {
	return true
}

func (self StringType) Numeric() bool {
	return false
}

func (self StringType) Collection() bool {
	return true
}

func (self StringType) Equals(t Type) bool {
	return t.Kind() == String
}

func (self StringType) HasMember(name string) bool {
	_, ok := self.members[name]
	return ok
}

func (self StringType) GetMember(name string) Type {
	return self.members[name]
}
