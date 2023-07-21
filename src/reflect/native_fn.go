package reflect

func NewNativeFn(
	name string,
	params []Param,
	returnType Type,
	value func([]*Value) *Value,
) *Value {
	return &Value{
		_type: NewNativeFnType(
			name,
			params,
			returnType,
		),
		_value:   value,
		_members: map[string]*Value{},
	}
}

func (self Value) NativeFnType() NativeFnType {
	return self._type.(NativeFnType)
}

func (self Value) IsNativeFn() bool {
	return self.Kind() == NativeFn
}

func (self Value) NativeFn() func([]*Value) *Value {
	return self._value.(func([]*Value) *Value)
}

func (self Value) NativeFnString() string {
	return self.NativeFnType().String()
}
