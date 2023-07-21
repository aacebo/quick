package reflect

type Value struct {
	_type    Type
	_value   any
	_members map[string]*Value
}

func (self Value) Type() Type {
	return self._type
}

func (self Value) Kind() Kind {
	return self._type.Kind()
}

func (self Value) Any() any {
	return self._value
}

func (self Value) ToString() string {
	if self.IsNil() {
		return self.NilString()
	}

	if self.IsBool() {
		return self.BoolString()
	}

	if self.IsByte() {
		return self.ByteString()
	}

	if self.IsInt() {
		return self.IntString()
	}

	if self.IsFloat() {
		return self.FloatString()
	}

	if self.IsFn() {
		return self.FnString()
	}

	if self.IsNativeFn() {
		return self.NativeFnString()
	}

	if self.IsMap() {
		return self.MapString()
	}

	if self.IsMod() {
		return self.ModString()
	}

	if self.IsSlice() {
		return self.SliceString()
	}

	return self.String()
}

func (self Value) HasMember(name string) bool {
	_, ok := self._members[name]
	return ok
}

func (self *Value) GetMember(name string) *Value {
	return self._members[name]
}

func (self *Value) SetMember(name string, value *Value) {
	v, ok := self._members[name]

	if ok && !v._type.Equals(value._type) {
		panic("invalid type")
	}

	self._members[name] = value
}

func (self *Value) DelMember(name string) {
	delete(self._members, name)
}
