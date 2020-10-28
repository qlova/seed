package link

import (
	"qlova.org/seed"
	"qlova.org/seed/use/html"
	"qlova.org/seed/use/html/attr"
)

//New returns a new HTML link element.
func New(options ...seed.Option) seed.Seed {
	return seed.New(html.SetTag("link"), seed.Options(options))
}

//Manifest returns a new HTML link element that links to a web manifest.
func Manifest(path string) seed.Seed {
	return New(attr.Set("rel", "manifest"), attr.Set("href", path))
}
