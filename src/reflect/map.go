package reflect

func NewMap[K ComparableType, V Type](key K, value V) Value {
	return Value{
		_type: MapType{
			key:   key,
			value: value,
		},
		_value: map[K]Value{},
	}
}

func (self Value) MapType() MapType {
	return self._type.(MapType)
}

func (self Value) IsMap() bool {
	return self.Kind() == Map
}
