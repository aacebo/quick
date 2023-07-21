package reflect

func NewMod() *Value {
	return &Value{
		_type:    NewModType(),
		_value:   nil,
		_members: map[string]*Value{},
	}
}

func (self Value) ModType() ModType {
	return self._type.(ModType)
}

func (self Value) IsMod() bool {
	return self.Kind() == Mod
}

func (self Value) Mod() map[string]*Value {
	return self._members
}

func (self Value) ModString() string {
	return self.ModType().String()
}
