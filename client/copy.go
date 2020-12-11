package client

import "qlova.org/seed/use/js"

//Copy copies the given string to the client's clipboard.
func Copy(s String) Script {
	return js.Func("navigator.clipboard.writeText").Run(s)
}
