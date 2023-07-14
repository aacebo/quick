package scope

import "quick/src/value"

type Scope struct {
	parent *Scope
	values map[string]*value.Value
}

func New() *Scope {
	return &Scope{
		parent: nil,
		values: map[string]*value.Value{},
	}
}

func NewChild(parent *Scope) *Scope {
	return &Scope{
		parent: parent,
		values: map[string]*value.Value{},
	}
}

func (self Scope) HasLocal(key string) bool {
	_, ok := self.values[key]
	return ok
}

func (self Scope) Has(key string) bool {
	if self.HasLocal(key) {
		return true
	}

	if self.parent != nil {
		return self.parent.Has(key)
	}

	return false
}

func (self Scope) Get(key string) *value.Value {
	if self.HasLocal(key) {
		return self.values[key]
	}

	if self.parent != nil {
		return self.parent.Get(key)
	}

	return nil
}

func (self *Scope) Set(key string, value value.Value) {
	self.values[key] = &value
}
