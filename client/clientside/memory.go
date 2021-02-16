package clientside

import (
	"encoding/base64"
	"math/big"
)

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

var address int64

func NewAddress() Address {
	address++
	return Address(base64.RawURLEncoding.EncodeToString(big.NewInt(address).Bytes()))
}
