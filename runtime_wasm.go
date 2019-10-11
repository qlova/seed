package seed

import (
	"bytes"
	"net/http"
	"net/url"
	"syscall/js"

	"github.com/qlova/seed/script"
	"github.com/qlova/seed/user"
)

//TODO Clean this up.
var Production bool
var exporting bool

func cleanup() {

}

type fakeResponseWriter struct {
	bytes.Buffer
	header http.Header
	status int
}

func (f *fakeResponseWriter) Header() http.Header {
	return f.header
}

func (f *fakeResponseWriter) WriteHeader(status int) {
	f.status = status
}

//Launch launches this app on the browser.
func (runtime Runtime) Launch(args ...string) {
	var window = js.Global()
	var document = window.Get("document")
	document.Call("write", string(runtime.app.Render(Default)))

	window.Set("rpc", window.Get("Object").New())
	var rpc = window.Get("rpc")

	for name, function := range script.Exports {
		rpc.Set("/call/"+name, js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			var f = function.Interface().(func(user.User))
			var u user.User
			var response fakeResponseWriter

			var request http.Request
			request.Form = url.Values{}

			//Parse FormData()
			if len(args) > 0 {
				var formdata = args[0]
				if formdata.Truthy() {
					var iterator = formdata.Call("entries")
					for {
						var next = iterator.Call("next")
						if next.Get("done").Truthy() {
							break
						}

						var value = next.Get("value")

						request.Form.Set(value.Index(0).String(), value.Index(1).String())
					}
				}
			}

			u = u.FromHandler(&response, &request)
			f(u)
			u.Close()

			return response.String()
		}))
	}

	document.Call("dispatchEvent", window.Get("Event").New("DOMContentLoaded"))

	var signal = make(chan bool)
	<-signal
}
