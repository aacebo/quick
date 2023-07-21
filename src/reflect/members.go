package reflect

type MemberCallback func(*Value) *Value

var members = map[Kind]map[string]MemberCallback{
	Bool:     {},
	Byte:     {},
	Float:    {},
	Fn:       {},
	Int:      {},
	Map:      {},
	Mod:      {},
	NativeFn: {},
	Nil:      {},
	Slice:    {},
	String: {
		"at":    stringMemberAt(),
		"len":   stringMemberLen(),
		"slice": stringMemberSlice(),
	},
}

var memberTypes = map[Kind]map[string]Type{
	Bool:     {},
	Byte:     {},
	Float:    {},
	Fn:       {},
	Int:      {},
	Map:      {},
	Mod:      {},
	NativeFn: {},
	Nil:      {},
	Slice:    {},
	String: {
		"at":  NewNativeFnType("at", []Param{{Name: "i", Type: NewIntType()}}, NewByteType()),
		"len": NewNativeFnType("len", []Param{}, NewIntType()),
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
