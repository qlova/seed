package clientside

import "qlova.org/seed/client"

//Go2
/*
type PointerTo[T Variable] struct {
	T
}
*/

//PointerToBool can be sent as an argument to a client.Go style call.
type PointerToBool struct {
	*Bool
}

//AsArgument implements client.Argument.
func (p PointerToBool) AsArgument() client.Value {
	address, memory := p.Variable()
	return client.NewString(string(memory) + ":" + string(address))
}
