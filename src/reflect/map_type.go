package reflect

import "fmt"

type MapType struct {
	key   Type
	value Type
}

func NewMapType(key Type, value Type) *MapType {
	return &MapType{
		key:   key,
		value: value,
	}
}

func (self MapType) Kind() Kind {
	return Map
}

func (self MapType) Name() string {
	return Map.String()
}

func (self MapType) String() string {
	return fmt.Sprintf(
		"%s[%s, %s]",
		Map.String(),
		self.key.Name(),
		self.value.Name(),
	)
}

func (self MapType) Comparable() bool {
	return false
}

func (self MapType) Numeric() bool {
	return false
}

func (self MapType) Collection() bool {
	return true
}

func (self MapType) Len() int {
	panic("method not supported")
}

func (self MapType) Key() Type {
	return self.key
}

func (self MapType) Value() Type {
	return self.value
}
