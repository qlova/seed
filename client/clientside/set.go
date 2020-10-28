package clientside

import (
	"qlova.org/seed/client"
	"qlova.org/seed/use/js"
)

//Set implements a type-safe clientside equivalant to map[interface{}]struct{}
type Set struct {
	Name string

	address Address
	Memory  Memory

	Value map[interface{}]struct{}
}

//GetDefaultValue implements Variable
func (s *Set) GetDefaultValue() client.Value {
	return js.NewValue("new Set()")
}

//Variable implements Variable
func (s *Set) Variable() (Address, Memory) {
	if s.address == "" {
		if s.Name != "" {
			s.address = Address(s.Name)
		} else {
			s.address = NewAddress()
		}
	}
	return s.address, s.Memory
}

//GetSet implements js.AnySet
func (s *Set) GetSet() js.Set {
	address, memory := s.Variable()
	return js.Set{Value: js.Call(js.Func("q.getset"),
		client.NewString(string(address)), client.NewString(string(memory)))}
}

//GetValue implements client.Value
func (s *Set) GetValue() js.Value {
	return s.GetSet().Value
}

//GetBool implements client.Value
func (s *Set) GetBool() js.Bool {
	return s.GetSet().GetBool()
}

//Add adds an item to the set.
func (s *Set) Add(item client.Value) client.Script {
	return s.GetSet().Run("add", item)
}

//Has returns true if the set contains the given item.
func (s *Set) Has(item client.Value) client.Bool {
	return s.GetSet().Call("has", item)
}

//Remove removes the item from the set if it exists.
func (s *Set) Remove(item client.Value) client.Script {
	return s.GetSet().Run("delete", item)
}

//Empty sets the set to the empty set.
func (s *Set) Empty() client.Script {
	return s.GetSet().Run("clear")
}
