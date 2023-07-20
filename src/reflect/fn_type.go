package reflect

import "fmt"

type Param struct {
	Name string
	Type Type
}

type FnType struct {
	name       string
	params     []Param
	returnType Type
}

func (self FnType) Kind() Kind {
	return Fn
}

func (self FnType) Name() string {
	return self.name
}

func (self FnType) String() string {
	return fmt.Sprintf("<fn %s>", self.name)
}

func (self FnType) Comparable() bool {
	return false
}

func (self FnType) Numeric() bool {
	return false
}

func (self FnType) Collection() bool {
	return false
}

func (self FnType) Len() int {
	panic("method not supported")
}

func (self FnType) Params() []Param {
	return self.params
}

func (self FnType) ReturnType() Type {
	return self.returnType
}
