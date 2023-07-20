package interpreter

import (
	"quick/src/reflect"
)

type Scope struct {
	parent *Scope
	values map[string]*reflect.Value
}

func NewScope() *Scope {
	return &Scope{
		parent: nil,
		values: map[string]*reflect.Value{},
	}
}

func NewChildScope(parent *Scope) *Scope {
	return &Scope{
		parent: parent,
		values: map[string]*reflect.Value{},
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

func (self Scope) GetLocal(key string) *reflect.Value {
	return self.values[key]
}

func (self Scope) Get(key string) *reflect.Value {
	if self.HasLocal(key) {
		return self.GetLocal(key)
	}

	if self.parent != nil {
		return self.parent.Get(key)
	}

	return nil
}

func (self *Scope) Set(key string, value *reflect.Value) {
	self.values[key] = value
}
