package reflect

func NewMod() Value {
	return Value{
		_type: ModType{
			exports: map[string]Type{},
		},
		_value: map[string]Value{},
	}
}

func (self Value) ModType() ModType {
	return self._type.(ModType)
}

func (self Value) IsMod() bool {
	return self.Kind() == Mod
}

func (self Value) Mod() map[string]Value {
	return self._value.(map[string]Value)
}
