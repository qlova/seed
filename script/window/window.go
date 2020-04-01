package window

import "github.com/qlova/seed/script"

//Alert creates a popup window with the provided text message.
func Alert(msg script.String) script.Script {
	return func(q script.Ctx) {
		q.Javascript(`window.alert(%v)`, msg)
	}
}
