package page

import (
	"qlova.org/seed"
	"qlova.org/seed/web/html/attr"
)

//SetTitle sets the title of this page.
func SetTitle(title string) seed.Option {
	return attr.Set("data-title", title)
}

//SetPath sets the url path of this page.
func SetPath(path string) seed.Option {
	return attr.Set("data-path", path)
}
