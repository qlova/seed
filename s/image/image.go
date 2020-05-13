package image

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/asset"
	"github.com/qlova/seed/css"
	"github.com/qlova/seed/html"
	"github.com/qlova/seed/html/attr"
	"github.com/qlova/seed/state"
)

//New returns a new image widget.
func New(src state.AnyString, options ...seed.Option) seed.Seed {
	return seed.New(
		html.SetTag("img"),
		state.SetSource(src),
		seed.Options(options),
	)
}

//SetFallback sets a fallback image to be used, this must be a local, cacheable image.
func SetFallback(src string) seed.Option {
	src = asset.Path(src)

	return seed.Options{
		attr.Set(`onError`, `this.src='`+src+`'`),
		asset.New(src),
	}
}

//Crop crops the image to fill its bounding box.
func Crop() seed.Option {
	return css.SetObjectFit(css.Cover)
}
