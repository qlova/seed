package main

import (
	"qlova.org/seed"
	"qlova.org/seed/new/feed"
	"qlova.org/seed/new/page"
	"qlova.org/seed/set"
	"qlova.org/seed/set/transition"
	"qlova.tech/rgb"
)

type CustomPage struct{}

func (p CustomPage) Page(r page.Router) seed.Seed {
	var holidays = feed.With(GetCustom)

	return page.New(
		transition.Fade(),
		set.Scrollable(),
		page.OnEnter(holidays.Refresh()),
		set.Color(rgb.Lavender),
		NewHolidays(holidays),
	)
}
