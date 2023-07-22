package reflect

type Value struct {
	_type  Type
	_value any
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

func (self Value) HasMember(name string) bool {
	_, ok := members[self.Kind()][name]

	if !ok && self.IsMod() {
		return self.HasExport(name)
	}

	return ok
}

func (self *Value) GetMember(name string) *Value {
	cb, ok := members[self.Kind()][name]

	if !ok && self.IsMod() {
		return self.GetExport(name)
	}

	return cb(self)
}
