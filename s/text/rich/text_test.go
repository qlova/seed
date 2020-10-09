package rich_test

import (
	"image/color"
	"testing"

	"qlova.org/seed/s/text/rich"
	"qlova.org/should"
)

func Test_Text(t *testing.T) {
	should.Be("<span style='color:#000000;'><strong><em>Hello World</em></strong></span>")(
		rich.Text("Hello World").Italic().Bold().In(color.Black).HTML(),
	).Test(t)

	should.Be("<strong>Hello</strong><em> World</em>")(
		(rich.Text("Hello").Bold() + rich.Text(" World").Italic()).HTML(),
	).Test(t)

	should.Be("<img style='height:1em;' src='img.png'>Hello World")(
		(rich.Icon("img.png") + rich.Text("Hello World")).HTML(),
	).Test(t)
}
