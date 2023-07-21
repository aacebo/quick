package reflect

func NewString(value string) *Value {
	self := &Value{
		_type:  NewStringType(),
		_value: value,
	}

	self._members = map[string]*Value{
		"len": self.StringMemberLen(),
		"at":  self.StringMemberAt(),
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
