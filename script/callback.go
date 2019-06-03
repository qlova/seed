package script

import qlova "github.com/qlova/script"

type Attachable interface {
	AttachTo(request string, index int) string
}

func (q Script) Attach(attachables ...Attachable) callWithFormData {
	var variable = Unique()

	q.Javascript(`var `+variable+" = new FormData();")

	for i, attachable := range attachables {
		q.Javascript(attachable.AttachTo(variable, i+1))
	}

	return callWithFormData{variable, q}
}

type callWithFormData struct {
	formdata string
	q Script
}

func (c callWithFormData) Call(f interface{}, args ...qlova.Type) Promise {	
	return c.q.rpc(f, c.formdata, args...)
}