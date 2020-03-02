package style

import (
	"image/color"

	"github.com/qlova/seed"
	"github.com/qlova/seed/style/css"
)

type cssable interface {
	CSS() css.Style
}

func toOption(f func(css.Style)) seed.Option {
	return func(any interface{}) {

		if css, ok := any.(cssable); ok {
			f(css.CSS())
		}
	}
}

func SetColor(color color.Color) seed.Option {
	return toOption(func(s css.Style) {
		s.SetBackgroundColor(css.Colour(color))
	})
}
