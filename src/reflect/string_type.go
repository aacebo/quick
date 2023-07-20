package reflect

type StringType struct{}

func NewStringType() StringType {
	return StringType{}
}

func (self StringType) Kind() Kind {
	return String
}

func (self StringType) Name() string {
	return String.String()
}

func (self StringType) String() string {
	return String.String()
}

func (self StringType) Len() int {
	panic("method not supported")
}

func (self StringType) Comparable() bool {
	return true
}

func (self StringType) Numeric() bool {
	return false
}

func (self StringType) Collection() bool {
	return true
}

func (self StringType) Equals(t Type) bool {
	return t.Kind() == String
}
