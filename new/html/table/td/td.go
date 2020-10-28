package td

import (
	"qlova.org/seed"
	"qlova.org/seed/web/html"
)

//New returns an HTML 'td' element.
func New(options ...seed.Option) seed.Seed {
	return seed.New(html.SetTag("td"), seed.Options(options))
}
