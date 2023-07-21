package reflect

func NewFn(
	name string,
	params []Param,
	returnType Type,
	value any,
) *Value {
	return &Value{
		_type: NewFnType(
			name,
			params,
			returnType,
		),
		_value:   value,
		_members: map[string]*Value{},
	}
}

func (self Value) FnType() FnType {
	return self._type.(FnType)
}

func (self Value) IsFn() bool {
	return self.Kind() == Fn
}

func (self Value) Fn() any {
	return self._value
}

func (self Value) FnString() string {
	return self.FnType().String()
}
