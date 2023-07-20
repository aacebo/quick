package parser

import (
	"quick/src/reflect"
)

type Scope struct {
	parent *Scope
	types  map[string]reflect.Type
}

func NewScope() *Scope {
	return &Scope{
		parent: nil,
		types: map[string]reflect.Type{
			"byte":   reflect.NewByteType(),
			"bool":   reflect.NewBoolType(),
			"int":    reflect.NewIntType(),
			"float":  reflect.NewFloatType(),
			"string": reflect.NewStringType(),
		},
	}
}

func NewChildScope(parent *Scope) *Scope {
	return &Scope{
		parent: parent,
		types:  map[string]reflect.Type{},
	}
}

func (self Scope) HasLocal(key string) bool {
	_, ok := self.types[key]
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

func (self Scope) GetLocal(key string) reflect.Type {
	return self.types[key]
}

func (self Scope) Get(key string) reflect.Type {
	if self.HasLocal(key) {
		return self.GetLocal(key)
	}

	if self.parent != nil {
		return self.parent.Get(key)
	}

	return nil
}

func (self *Scope) Set(key string, _type reflect.Type) {
	self.types[key] = _type
}
