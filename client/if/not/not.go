package not

import (
	"qlova.org/seed/client"
	"qlova.org/seed/client/if/the"
	"qlova.org/seed/use/js"
)

//True returns a client.Bool that is true when b is false.
func True(b client.Bool) client.Bool {
	return the.Bool(js.NewValue("(!%v)", b), b)
}
