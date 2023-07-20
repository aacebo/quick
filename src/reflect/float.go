package reflect

func NewFloat(value float64) Value {
	return Value{
		_type:  FloatType{},
		_value: value,
	}
}

func (self Value) FloatType() FloatType {
	return self._type.(FloatType)
}

func (self Value) IsFloat() bool {
	return self.Kind() == Float
}

func (self Value) Float() float64 {
	return self._value.(float64)
}

func (self *Value) SetFloat(value float64) {
	self._value = value
}
