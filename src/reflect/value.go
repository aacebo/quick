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
