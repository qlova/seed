package filepicker

import "strconv"
import "github.com/qlova/seed"
import "github.com/qlova/seed/script"

type Widget struct {
	seed.Seed
}

func New(types ...string) Widget {
	widget := seed.New()
	widget.SetTag("Input")

	if len(types) > 0 {
		widget.SetAttributes(`type="file" accept="`+types[0]+`"`)
	} else {
		widget.SetAttributes(`type="file" accept="*"`)
	}

	return  Widget{widget}
}

func AddTo(parent seed.Interface, types ...string) Widget {
	var widget = New(types...)
	parent.Root().Add(widget)
	return widget
}

type Script struct {
	script.Seed
}

func (w Widget) Script(q seed.Script) Script {
	return Script{w.Seed.Script(q)}
}

func (s Script) AttachTo(request string, index int) string {
	
	return "for (let i = 0; i < "+s.Element()+".files.length; i++) "+request+
		`.append("attachment-`+strconv.Itoa(index)+`-"+(i+1), `+s.Element()+`.files[i], `+s.Element()+`.files[i].name);`
}