package progressbar

import (
	"fmt"

	"qlova.org/seed"
	"qlova.org/seed/client/clientfmt"
	"qlova.org/seed/client/clientside"
	"qlova.org/seed/new/html/progress"
	"qlova.org/seed/use/html/attr"
)

//New returns a new progressbar.
func New(options ...seed.Option) seed.Seed {
	return progress.New(
		attr.Set("max", "1"),
		seed.Options(options),
	)
}

//Set sets the progress to reflect the value.
func Set(value float64) seed.Option {
	return attr.Set("value", fmt.Sprint(value))
}

//SetTo sets the progress to reflect the value.
func SetTo(value *clientside.Float64) seed.Option {
	return attr.SetTo("value", clientfmt.Sprintf("%v", value))
}
