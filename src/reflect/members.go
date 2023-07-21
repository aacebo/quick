package reflect

type MemberCallback func(*Value) *Value

var members = map[Kind]map[string]MemberCallback{
	Bool: {
		"to_string": memberToString(),
	},
	Byte: {
		"to_string": memberToString(),
	},
	Float: {
		"to_string": memberToString(),
	},
	Fn: {
		"to_string": memberToString(),
	},
	Int: {
		"to_string": memberToString(),
	},
	Map: {
		"to_string": memberToString(),
	},
	Mod: {
		"to_string": memberToString(),
	},
	NativeFn: {
		"to_string": memberToString(),
	},
	Nil: {
		"to_string": memberToString(),
	},
	Slice: {
		"to_string": memberToString(),
	},
	String: {
		"at":      stringMemberAt(),
		"len":     stringMemberLen(),
		"replace": stringMemberReplace(),
		"slice":   stringMemberSlice(),
	},
}

var memberTypes = map[Kind]map[string]Type{
	Bool: {
		"to_string": NewNativeFnType("to_string", []Param{}, NewStringType()),
	},
	Byte: {
		"to_string": NewNativeFnType("to_string", []Param{}, NewStringType()),
	},
	Float: {
		"to_string": NewNativeFnType("to_string", []Param{}, NewStringType()),
	},
	Fn: {
		"to_string": NewNativeFnType("to_string", []Param{}, NewStringType()),
	},
	Int: {
		"to_string": NewNativeFnType("to_string", []Param{}, NewStringType()),
	},
	Map: {
		"to_string": NewNativeFnType("to_string", []Param{}, NewStringType()),
	},
	Mod: {
		"to_string": NewNativeFnType("to_string", []Param{}, NewStringType()),
	},
	NativeFn: {
		"to_string": NewNativeFnType("to_string", []Param{}, NewStringType()),
	},
	Nil: {
		"to_string": NewNativeFnType("to_string", []Param{}, NewStringType()),
	},
	Slice: {
		"to_string": NewNativeFnType("to_string", []Param{}, NewStringType()),
	},
	String: {
		"at":  NewNativeFnType("at", []Param{{Name: "i", Type: NewIntType()}}, NewByteType()),
		"len": NewNativeFnType("len", []Param{}, NewIntType()),
		"replace": NewNativeFnType(
			"replace",
			[]Param{
				{Name: "from", Type: NewStringType()},
				{Name: "to", Type: NewStringType()},
			},
			NewStringType(),
		),
		"slice": NewNativeFnType(
			"slice",
			[]Param{
				{Name: "start", Type: NewIntType()},
				{Name: "end", Type: NewIntType()},
			},
			NewStringType(),
		),
	},
}

func memberToString() MemberCallback {
	return func(self *Value) *Value {
		return NewNativeFn("to_string", []Param{}, NewStringType(), func(args []*Value) *Value {
			return NewString(self.ToString())
		})
	}
}
