package clientside

import "qlova.org/seed/client"

//Memory is a type of client memory for SideValues.
type Memory string

//Memory types.
const (
	ShortTermMemory Memory = ""
	SessionMemory   Memory = "session"
	LongTermMemory  Memory = "storage"

	LocalMemory Memory = "local"
)

//Address is a Memory Address
type Address string

//Variable is a clientside variable.
type Variable interface {
	Variable() (Address, Memory)
	GetDefaultValue() client.Value
}
