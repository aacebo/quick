package value

type Bool bool

func (self Bool) Kind() Kind {
	return BOOL
}

func (self Bool) Truthy() Bool {
	return self == true
}

func (self Bool) String() string {
	if self == true {
		return "true"
	}

	return "false"
}

func (self Bool) Eq(other Bool) Bool {
	return self == other
}
