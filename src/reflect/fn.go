package reflect

func NewFn(name string, params []Param, rtype Type) Value {
	return Value{
		_type: FnType{
			name:       name,
			params:     params,
			returnType: rtype,
		},
		_value: nil,
	}
}

func (self Value) FnType() FnType {
	return self._type.(FnType)
}

func (self Value) IsFn() bool {
	return self.Kind() == Fn
}
