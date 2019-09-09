package script

import qlova "github.com/qlova/script"

//Attachable is a something that can be attached to a Go call.
type Attachable interface {
	AttachTo(request string, index int) string
}

//Attach attaches Attachables and returns an AttachCall.
func (q Script) Attach(attachables ...Attachable) Attached {
	var variable = Unique()

	q.Javascript(`var ` + variable + " = new FormData();")

	for i, attachable := range attachables {
		q.Javascript(attachable.AttachTo(variable, i+1))
	}

	return Attached{variable, q}
}

//Attached has attachments and these will be passed to the Go function that is called.
type Attached struct {
	formdata string
	q        Script
}

//Go calls a Go function f, with args. Returns a promise.
func (c Attached) Go(f interface{}, args ...qlova.Type) Promise {
	return c.q.rpc(f, c.formdata, args...)
}
