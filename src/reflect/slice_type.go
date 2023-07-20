package reflect

import "fmt"

type SliceType struct {
	_type  Type
	length int
}

func (self SliceType) Kind() Kind {
	return Slice
}

func (self SliceType) Name() string {
	return Slice.String()
}

func (self SliceType) String() string {
	return fmt.Sprintf(
		"[]%s",
		self._type.Name(),
	)
}

func (self SliceType) Len() int {
	return self.length
}

func (self SliceType) Comparable() bool {
	return false
}

func (self SliceType) Numeric() bool {
	return false
}

func (self SliceType) Collection() bool {
	return true
}

func (self SliceType) Type() Type {
	return self._type
}

func (self SliceType) Equals(t Type) bool {
	if t.Kind() != Slice {
		return false
	}

	if !self._type.Equals(t) {
		return false
	}

	return true
}
