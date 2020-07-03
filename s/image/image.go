package image

import (
	"fmt"

	"qlova.org/seed"
	"qlova.org/seed/asset"
	"qlova.org/seed/asset/assets"
	"qlova.org/seed/client"
	"qlova.org/seed/css"
	"qlova.org/seed/html"
	"qlova.org/seed/html/attr"
	"qlova.org/seed/sum"
	"qlova.org/seed/units"
)

//New returns a new image widget.
func New(src sum.String, options ...seed.Option) seed.Seed {
	if src != nil {
		src = asset.Path(src)
	}

	return seed.New(
		html.SetTag("img"),
		attr.Set("src", src),

		seed.Options(options),
	)
}

//Set sets the image path.
func Set(path string) seed.Option {
	return attr.Set("src", asset.Path(path))
}

//SetTo sets the image path.
func SetTo(path client.String) seed.Option {
	return attr.Set("src", asset.Path(path))
}

func SetSize(size float32) seed.Option {
	return seed.Options{
		//style.SetWidth(style.Unit(complex(size, 0)) * unit.Px),
		//style.SetHeight(style.Unit(complex(size, 0)) * unit.Px),

		attr.Set("width", fmt.Sprint(size)),
		attr.Set("height", fmt.Sprint(size)),
	}
}

//SetOffset sets the offset of this image within its container.
func SetOffset(x, y units.Unit) css.Rule {
	return css.Set("object-position", fmt.Sprintf("%v %v", css.Measure(x), css.Measure(y)))
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
