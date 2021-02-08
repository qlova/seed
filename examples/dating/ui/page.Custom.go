package ui

import (
	"dating"

	"qlova.org/seed"
	"qlova.org/seed/new/feed"
	"qlova.org/seed/new/page"
	"qlova.org/seed/set"
	"qlova.org/seed/set/transition"
	"qlova.org/seed/use/css/units/vmin"
	"qlova.tech/rgb"
)

type CustomPage struct{}

func (p CustomPage) Page(r page.Router) seed.Seed {
	var holidays = feed.With(dating.GetCustom)

	return page.New(
		transition.Fade(),
		set.Scrollable(),

		set.If.Small().Portrait(
			set.Width(vmin.New(100)),
		),

		page.OnEnter(holidays.Refresh()),
		set.Color(rgb.Lavender),
		NewHolidays(holidays),
	)
}
