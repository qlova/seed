package title

import (
	"qlova.org/seed"
	"qlova.org/seed/web/html"
)

//New returns a new HTML title tag with the given title.
func New(title string, options ...seed.Option) seed.Seed {
	return seed.New(
		html.SetTag("title"),
		html.SetInnerText(title),
	)
}
