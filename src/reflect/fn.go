package reflect

import "quick/src/ast/stmt"

func NewFn(
	name string,
	params []Param,
	returnType Type,
	value *stmt.Fn,
) Value {
	return Value{
		_type: FnType{
			name:       name,
			params:     params,
			returnType: returnType,
		},
		_value: value,
	}
}

func (self Value) FnType() FnType {
	return self._type.(FnType)
}

func (self Value) IsFn() bool {
	return self.Kind() == Fn
}

func (self Value) Fn() *stmt.Fn {
	return self._value.(*stmt.Fn)
}
