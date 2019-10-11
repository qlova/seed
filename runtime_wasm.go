package seed

import "syscall/js"

//TODO Clean this up.
var Production bool
var exporting bool

func cleanup() {

}

//Launch launches this app on the browser.
func (runtime Runtime) Launch(args ...string) {
	var window = js.Global()
	var document = window.Get("document")
	document.Call("write", string(runtime.app.Render(Default)))
	document.Call("dispatchEvent", window.Get("Event").New("DOMContentLoaded"))
}
