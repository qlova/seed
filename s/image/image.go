package image

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/asset"
	"github.com/qlova/seed/html"
	"github.com/qlova/seed/html/attr"
)

//New returns a new image widget.
func New(src string, options ...seed.Option) seed.Seed {
	src = asset.Path(src)

	return seed.New(
		html.SetTag("img"),
		attr.Set("src", src),
		attr.Set("alt", src),

		asset.New(src),

		seed.Options(options),
	)
}
