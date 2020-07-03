package line

import (
	"qlova.org/seed"
	"qlova.org/seed/s/html/div"
)

//New returns a new line.
func New(options ...seed.Option) seed.Seed {
	return div.New(options...)
}
