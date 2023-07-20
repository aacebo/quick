package reflect

type ModType struct {
	exports map[string]Type
}

func (self ModType) Kind() Kind {
	return Mod
}

func (self ModType) Name() string {
	return Mod.String()
}

func (self ModType) String() string {
	return Mod.String()
}

func (self ModType) Comparable() bool {
	return false
}

func (self ModType) Numeric() bool {
	return false
}

func (self ModType) Collection() bool {
	return true
}

func (self ModType) Len() int {
	return len(self.exports)
}
