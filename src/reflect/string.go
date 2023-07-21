package reflect

func NewString(value string) *Value {
	self := &Value{
		_type:  NewStringType(),
		_value: value,
	}

	self._members = map[string]*Value{
		"len": NewNativeFn("len", []Param{}, NewIntType(), func(args []*Value) *Value {
			return NewInt(self.Len())
		}),
		"at": NewNativeFn("at", []Param{{Name: "i", Type: NewIntType()}}, NewByteType(), func(args []*Value) *Value {
			return NewByte(self.String()[args[0].Int()])
		}),
	}

	return self
}

func (self Value) StringType() StringType {
	return self._type.(StringType)
}

func (self Value) IsString() bool {
	return self.Kind() == String
}

func (self Value) String() string {
	return self._value.(string)
}

func (self *Value) SetString(value string) {
	self._value = value
}

func (self Value) SubString(i int, j int) string {
	return self.String()[i:j]
}

func (self *Value) Append(value string) {
	self._value = self.String() + value
}