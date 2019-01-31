package listbox

import "fmt"
import "github.com/qlova/seed"

type Widget struct {
	seed.Seed
}

func New(values ...string) Widget {
	widget := seed.New()
	widget.SetTag("select")
	
	var content string
	
	for _, value := range values {
		content += fmt.Sprint("<option value='", value, "'>", value, "</option>")
	}
	
	widget.SetContent(content)

	return Widget{widget}
}

func AddTo(parent seed.Interface, values ...string) Widget {
	var widget = New(values...)
	parent.Root().Add(widget)
	return widget
}
