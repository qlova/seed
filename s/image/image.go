package image

import (
	"fmt"

	"qlova.org/seed"
	"qlova.org/seed/asset"
	"qlova.org/seed/asset/assets"
	"qlova.org/seed/css"
	"qlova.org/seed/html"
	"qlova.org/seed/html/attr"
	"qlova.org/seed/sum"
)

//New returns a new image widget.
func New(src sum.String, options ...seed.Option) seed.Seed {
	src = asset.Path(src)

	return seed.New(
		html.SetTag("img"),
		attr.Set("src", src),

		seed.Options(options),
	)
}

func SetSize(size float32) seed.Option {
	return seed.Options{
		//style.SetWidth(style.Unit(complex(size, 0)) * unit.Px),
		//style.SetHeight(style.Unit(complex(size, 0)) * unit.Px),

		attr.Set("width", fmt.Sprint(size)),
		attr.Set("height", fmt.Sprint(size)),
	}
}

//SetFallback sets a fallback image to be used, this must be a local, cacheable image.
func SetFallback(src string) seed.Option {
	src = asset.Path(src).(string)

	return seed.Options{
		attr.Set(`onError`, `this.src='`+src+`'`),
		assets.New(src),
	}
}

//Crop crops the image to fill its bounding box.
func Crop() seed.Option {
	return css.SetObjectFit(css.Cover)
}
