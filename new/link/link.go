package link

import (
	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/new/html/a"
	"qlova.org/seed/use/html/attr"
)

//New returns a new link.
func New(options ...seed.Option) seed.Seed {
	return a.New(options...)
}

//Set the url of the link.
func Set(url string) seed.Option {
	return attr.Set("href", url)
}

//SetTo sets the url of the link to a client.String
func SetTo(url client.String) seed.Option {
	return attr.SetTo("href", url)
}
