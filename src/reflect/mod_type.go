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

func (self ModType) Len() int {
	return len(self.exports)
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

func (self ModType) Equals(t Type) bool {
	if t.Kind() != Mod {
		return false
	}

	mod := t.(ModType)

	for key, _type := range self.exports {
		et, ok := mod.exports[key]

		if !ok || !_type.Equals(et) {
			return false
		}
	}

	return true
}

func (self ModType) HasExport(name string) bool {
	_, ok := self.exports[name]
	return ok
}

func (self ModType) GetExport(name string) Type {
	return self.exports[name]
}
