package link

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/html"
	"github.com/qlova/seed/html/attr"
)

//New returns a new HTML link element.
func New(options ...seed.Option) seed.Seed {
	return seed.New(html.SetTag("link").And(options...))
}

//Manifest returns a new HTML link element that links to a web manifest.
func Manifest(path string) seed.Seed {
	return New(attr.Set("rel", "manifest"), attr.Set("href", path))
}
