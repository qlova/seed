package rich_test

import (
	"image/color"
	"testing"

	"qlova.org/seed/new/text/rich"
	"qlova.org/should"
)

func Test_Text(t *testing.T) {
	should.Be("<span style='color:#000000;'><strong><em>Hello World</em></strong></span>")(
		rich.Text("Hello World").Italic().Bold().In(color.Black).HTML(),
	).Test(t)

	should.Be("<span style='color:#000000;'><strong><em>Hello </em></strong></span>World")(
		(rich.Text("Hello ").Italic().Bold().In(color.Black) + rich.Text("World")).HTML(),
	).Test(t)

	should.Be("<strong>Hello</strong><em> World</em>")(
		(rich.Text("Hello").Bold() + rich.Text(" World").Italic()).HTML(),
	).Test(t)

	should.Be("<img style='margin-top: 0.1em;vertical-align:text-top;height:1em;font-size:inherit;' src='img.png'>Hello World")(
		(rich.Icon("img.png") + rich.Text("Hello World")).HTML(),
	).Test(t)
}
