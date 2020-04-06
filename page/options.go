package page

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/html/attr"
)

//SetTitle sets the title of this page.
func SetTitle(title string) seed.Option {
	return attr.Set("data-title", title)
}

//SetPath sets the url path of this page.
func SetPath(path string) seed.Option {
	return attr.Set("data-path", path)
}
