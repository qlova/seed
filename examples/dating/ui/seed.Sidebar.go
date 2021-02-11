package ui

import (
	"dating/ui/style"

	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/new/column"
	"qlova.org/seed/new/expander"
	"qlova.org/seed/new/image"
	"qlova.org/seed/new/page"
	"qlova.org/seed/new/row"
	"qlova.org/seed/new/spacer"
	"qlova.org/seed/new/text"
	"qlova.org/seed/set"
	"qlova.org/seed/set/align"
	"qlova.org/seed/use/css/units/percentage/of"
	"qlova.org/seed/use/css/units/rem"
	"qlova.tech/rgb"
)

func NewSidebar() seed.Seed {
	var col = column.New()

	return col.With(
		set.If.Small().Portrait(
			row.Set(),
			set.Width(100%of.Parent),
			set.Height(rem.New(5.0)),
		),
		set.Color(rgb.Black),
		set.Width(rem.New(20.0)),
		set.Height(100%of.Parent),

		text.New(style.Text,
			text.Set("DatingApp"),
			text.SetColor(rgb.White),
			text.SetSize(rem.New(2.0)),
			text.Center(),
			align.Center(),
		),

		spacer.New(rem.New(2.0)),

		text.New(style.Text,
			text.Set("Popular"),
			text.SetColor(rgb.White),
			text.SetSize(rem.New(1.0)),
			text.Center(),
			align.Center(),

			client.OnClick(page.RouterOf(col).Goto(PopularPage{})),
		),

		spacer.New(rem.New(1.5)),

		text.New(style.Text,
			text.Set("Custom"),
			text.SetColor(rgb.White),
			text.SetSize(rem.New(1.0)),
			text.Center(),
			align.Center(),

			client.OnClick(page.RouterOf(col).Goto(CustomPage{})),
		),

		expander.New(),

		image.New(
			set.Width(rem.New(6.0)),
			set.If.Small(
				set.Width(rem.New(3.0)),
			),
			align.Center(),
			set.Margin(nil, rem.One),

			image.Set("heart-plus.svg"),
			client.OnClick(page.RouterOf(col).Goto(AddPage{})),
		),
	)
}
