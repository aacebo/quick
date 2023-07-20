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

func (self Value) HasExport(name string) bool {
	_, ok := self.Mod()[name]
	return ok
}

func (self Value) GetExport(name string) Value {
	return self.Mod()[name]
}

func (self *Value) SetExport(name string, value Value) {
	v := self.Mod()
	v[name] = value
	self._value = v
}

func (self *Value) DelExport(name string) {
	v := self.Mod()
	delete(v, name)
	self._value = v
}
