package reflect

func NewSlice[T Type](_type T) Value {
	return Value{
		_type: SliceType{
			_type: _type,
		},
		_value: []Value{},
	}
}

func (self Value) SliceType() SliceType {
	return self._type.(SliceType)
}

func (self Value) IsSlice() bool {
	return self.Kind() == Slice
}

func (self Value) Slice() []Value {
	return self._value.([]Value)
}

func (self Value) At(i int) Value {
	return self.Slice()[i]
}

func (self Value) SubSlice(i int, j int) []Value {
	return self.Slice()[i:j]
}

func (self *Value) Push(value Value) {
	v := self.Slice()
	v = append(v, value)
	self._value = v
}
