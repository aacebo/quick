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
		"at":      stringMemberAt(),
		"len":     stringMemberLen(),
		"replace": stringMemberReplace(),
		"slice":   stringMemberSlice(),
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
