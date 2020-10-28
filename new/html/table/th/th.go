package th

import (
	"qlova.org/seed"
	"qlova.org/seed/web/html"
)

//New returns an HTML 'th' element.
func New(options ...seed.Option) seed.Seed {
	return seed.New(html.SetTag("th"), seed.Options(options))
}
