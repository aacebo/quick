package value

type Bool bool

func (self Bool) Kind() Kind {
	return BOOL
}

func (self Bool) String() string {
	if self == true {
		return "true"
	}

	return "false"
}
