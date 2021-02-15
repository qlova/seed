package datebox

import (
	"time"

	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/use/html/attr"

	"qlova.org/seed/new/textbox"
)

//New returns a new datebox widget.
func New(options ...seed.Option) seed.Seed {
	return textbox.New(attr.Set("type", "date"), seed.Options(options))
}

//SetMin sets minimum date range constraint
func SetMin(min time.Time) seed.Option {
	return attr.Set("min", min.Format("2006-01-02"))
}

//SetMinTo sets dynamic minimum date range constraint
func SetMinTo(min client.Date) seed.Option {
	return attr.SetTo("min", min)
}

//SetMax sets maximum date range constraint
func SetMax(max time.Time) seed.Option {
	return attr.Set("max", max.Format("2006-01-02"))
}

//SetMaxTo sets dynamic maximum date range constraint
func SetMaxTo(max client.Date) seed.Option {
	return attr.SetTo("min", max)
}
