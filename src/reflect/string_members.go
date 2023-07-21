package reflect

func (self Value) StringMemberLen() *Value {
	return NewNativeFn("len", []Param{}, NewIntType(), func(args []*Value) *Value {
		return NewInt(self.Len())
	})
}

func (self Value) StringMemberAt() *Value {
	return NewNativeFn("at", []Param{{Name: "i", Type: NewIntType()}}, NewByteType(), func(args []*Value) *Value {
		return NewByte(self.String()[args[0].Int()])
	})
}
