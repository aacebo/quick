package reflect

func NewNil() Value {
	return Value{
		_type:  NilType{},
		_value: nil,
	}
}

func (self Value) NilType() NilType {
	return self._type.(NilType)
}

func (self Value) IsNil() bool {
	return self.Kind() == Nil
}

func (self *Value) SetNil() {
	self._value = nil
}
