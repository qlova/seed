package clientside

import (
	"errors"
	"strings"

	"qlova.org/seed/client"
	"qlova.org/seed/use/js"
)

//Pointer is a pointer to a memory address.
type Pointer struct {
	*MemoryAddress
}

//GetBool implements client.Bool
func (v Pointer) GetBool() js.Bool {
	return v.GetValue().GetBool()
}

//GetValue implements client.Value
func (v Pointer) GetValue() js.Value {
	address, memory := v.Variable()

	return js.Call(js.Func("q.getvar"),
		client.NewString(string(address)), client.NewString(string(memory)))
}

//AsArgument implements client.Argument.
func (p Pointer) AsArgument() client.Value {
	address, memory := p.Variable()
	return client.NewString(string(memory) + ":" + string(address))
}

//Variable is a clientside variable.
type Variable interface {
	Variable() (Address, Memory)
	GetDefaultValue() client.Value
}

//MemoryAddress describes a memory and address on the client.
type MemoryAddress struct {
	Name string

	address Address
	Memory  Memory
}

//Pointer returns a pointer to this bool, suitable for passing as an argument.
func (ma *MemoryAddress) Pointer() Pointer {
	return Pointer{ma}
}

//Parse parses a bool from a clientside.PointerToBool.AsArgument
func (ma *MemoryAddress) Parse(val string) error {
	splits := strings.Split(val, ":")
	if len(splits) != 2 {
		return errors.New("invalid clientside.Pointer format: " + val)
	}

	ma.Memory = Memory(splits[0])
	ma.address = Address(splits[1])

	return nil
}

//Variable implements Variable
func (ma *MemoryAddress) Variable() (Address, Memory) {
	if ma.address == "" {
		if ma.Name != "" {
			ma.address = Address(ma.Name)
		} else {
			ma.address = NewAddress()
		}
	}
	return ma.address, ma.Memory
}
