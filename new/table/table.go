package table

import (
	"fmt"

	"qlova.org/seed"
	"qlova.org/seed/use/css"
	"qlova.org/seed/new/html/table"
	"qlova.org/seed/new/html/table/td"
	"qlova.org/seed/new/html/table/th"
	"qlova.org/seed/new/html/table/tr"
	"qlova.org/seed/new/text"
	"qlova.org/seed/set"
	"qlova.org/seed/use/css/units"
)

//New returns a new Table.
func New(options ...seed.Option) seed.Seed {
	return table.New(options...)
}

//SetColumns sets the columns of the table, each argument is converted into a client.String and used as a column heading.
func SetColumns(labels ...interface{}) seed.Seed {
	var row = tr.New()
	for _, label := range labels {
		th.New(text.Set(fmt.Sprint(label))).AddTo(row)
	}
	return row
}

//Row returns a table row, each argument is converted into a subsequent Cell in the row.
func Row(labels ...interface{}) seed.Seed {
	var row = tr.New()
	for _, label := range labels {
		td.New(text.Set(fmt.Sprint(label))).AddTo(row)
	}
	return row
}

//SetBorder sets the border of this table.
func SetBorder(style set.BorderStyle) seed.Option {
	return seed.Options{
		css.Select(" th",
			set.Border(style).Rules()[0],
			css.SetBorderCollapse(css.Collapse),
		),
		css.Select(" td",
			set.Border(style).Rules()[0],
			css.SetBorderCollapse(css.Collapse),
		),
	}
}

//SetCellPadding sets the padding of the cells in this table.
func SetCellPadding(padding units.Unit, more ...units.Unit) seed.Option {
	return seed.Options{
		css.Select(" th",
			set.Padding(padding, more...),
		),
		css.Select(" td",
			set.Padding(padding, more...),
		),
	}
}
