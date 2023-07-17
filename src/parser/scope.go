package parser

import "quick/src/value"

type Scope struct {
	parent *Scope
	values map[string]*value.Definition
}

func NewScope() *Scope {
	return &Scope{
		parent: nil,
		values: map[string]*value.Definition{},
	}
}

func NewChildScope(parent *Scope) *Scope {
	return &Scope{
		parent: parent,
		values: map[string]*value.Definition{},
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

func (self Scope) GetLocal(key string) *value.Definition {
	return self.values[key]
}

func (self Scope) Get(key string) *value.Definition {
	if self.HasLocal(key) {
		return self.GetLocal(key)
	}

	if self.parent != nil {
		return self.parent.Get(key)
	}

	return nil
}

func (self *Scope) Set(key string, def *value.Definition) {
	self.values[key] = def
}
