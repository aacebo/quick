package reflect

func stringMemberLen() MemberCallback {
	return func(self *Value) *Value {
		return NewNativeFn("len", []Param{}, NewIntType(), func(args []*Value) *Value {
			return NewInt(self.Len())
		})
	}
}

func stringMemberAt() MemberCallback {
	return func(self *Value) *Value {
		return NewNativeFn("at", []Param{{Name: "i", Type: NewIntType()}}, NewByteType(), func(args []*Value) *Value {
			return NewByte(self.String()[args[0].Int()])
		})
	}
}

func stringMemberSlice() MemberCallback {
	return func(self *Value) *Value {
		return NewNativeFn("slice", []Param{
			{Name: "start", Type: NewIntType()},
			{Name: "end", Type: NewIntType()},
		}, NewStringType(), func(args []*Value) *Value {
			return NewString(self.String()[args[0].Int():args[1].Int()])
		})
	}
}
