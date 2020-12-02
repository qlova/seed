package compound

import "qlova.org/seed/client"

//Expression is a client.Value that will change when any of its variables change.
type Expression interface {
	client.Value

	Variables() []client.Value
}
