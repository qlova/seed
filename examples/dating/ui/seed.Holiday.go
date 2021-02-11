package ui

import (
	"dating"
	"dating/ui/style"

	"qlova.org/seed"
	"qlova.org/seed/client/clientfmt"
	"qlova.org/seed/new/feed"
	"qlova.org/seed/new/image"
	"qlova.org/seed/new/row"
	"qlova.org/seed/new/text"
	"qlova.org/seed/set"
	"qlova.org/seed/use/css/units/rem"
	"qlova.tech/rgb"
)

func NewHolidays(f *feed.Feed) seed.Seed {

	var holiday dating.Holiday
	f.Into(&holiday)

	return f.New(
		row.New(style.Border,
			set.Height(rem.New(10.0)),
			set.Margin(rem.One, rem.One/2),
			set.Color(rgb.White),

			image.New(
				set.Width(rem.New(10.0)),
				set.If.Small(
					set.Width(rem.New(5.0)),
				),

				image.Crop(),

				image.SetTo(f.String(holiday.Image)),
			),

			text.New(style.Text,
				text.SetSize(rem.New(2.0)),
				set.If.Small(
					text.SetSize(rem.New(1.5)),
				),
				set.Padding(rem.New(2.0), rem.New(1.0)),

				text.SetStringTo(clientfmt.Sprintf("%v until %v",
					f.String(holiday.Distance),
					f.String(holiday.Name))),
			),
		),
	)
}
