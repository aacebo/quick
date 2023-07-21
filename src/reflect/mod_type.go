package reflect

type ModType struct {
	members map[string]Type
}

func NewModType() ModType {
	return ModType{
		members: map[string]Type{},
	}
}

func (self ModType) Kind() Kind {
	return Mod
}

func (self ModType) Name() string {
	return Mod.String()
}

func (self ModType) String() string {
	return "<mod>"
}

func (self ModType) Len() int {
	return len(self.members)
}

func (self ModType) Comparable() bool {
	return false
}

func (self ModType) Numeric() bool {
	return false
}

func (self ModType) Collection() bool {
	return true
}

func (self ModType) Equals(t Type) bool {
	if t.Kind() != Mod {
		return false
	}

	mod := t.(ModType)

	for key, _type := range self.members {
		et, ok := mod.members[key]

		if !ok || !_type.Equals(et) {
			return false
		}
	}

	return true
}

func (self ModType) HasMember(name string) bool {
	_, ok := self.members[name]
	return ok
}

func (self ModType) GetMember(name string) Type {
	return self.members[name]
}

func (self *ModType) SetMember(name string, _type Type) {
	self.members[name] = _type
}
