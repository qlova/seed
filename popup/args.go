package popup

import (
	"qlova.org/seed/client/clientargs"
	"qlova.org/seed/js"
)

//PassGoValuesToClient can be embedded inside of a popup to enable the passthrough of Go arguments.
//these arguments must be decoded within the Popup(Manager) Seed method using the methods of PassGoValuesToClient.
type PassGoValuesToClient struct {
	clientargs.PassGoValues
}

//parseArgs returns the page arguments as a js.Object.
func parseArgs(popup Popup) (Popup, js.AnyObject) {
	if popup == nil {
		return popup, js.NewObject(nil)
	}

	NewPopup, object := clientargs.Parse(popup, js.Object{js.NewValue(`q.get(%v).args`, js.NewString(ID(popup)))})

	return NewPopup.(Popup), object
}
