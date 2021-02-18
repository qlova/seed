//Package screen provides client booleans for different screen sizes. Similar to a media query.
package screen

import (
	"qlova.org/seed/client"
	"qlova.org/seed/use/js"
)

//SizeQuery is a general screen-size query.
type SizeQuery uint8

//Query types.
const (
	Tiny SizeQuery = 1 << iota
	Small
	Medium
	Large
	Huge

	Portrait
	Landscape

	Smaller = Tiny | Small
	Larger  = Large | Huge
)

//Not returns the inverse of the query.
func Not(q SizeQuery) SizeQuery {
	p := q & Portrait
	l := q & Landscape
	return ^q | p | l
}

//Media returns the query as a CSS media query.
func (q SizeQuery) Media() string {

	var s string

	var p, l bool
	if q&Portrait+q&Landscape == 0 {
		p = true
		l = true
	}

	if q&Portrait != 0 {
		p = true
	}

	if q&Landscape != 0 {
		l = true
	}

	apply := func(portrait, landscape string) {
		if len(s) > 0 {
			s += ","
		}

		if p {
			s += "screen and " + portrait +
				" and (orientation:portrait)"
		}

		if len(s) > 0 {
			s += ","
		}

		if l {
			s += "screen and " + landscape +
				" and (orientation:landscape)"
		}
	}

	if q&Tiny != 0 {
		apply(`(max-width: 9.9999rem)`, `(max-height: 9.9999rem)`)
	}
	if q&Small != 0 {
		apply(`(min-width: 10rem) and (max-width: 29.999rem)`,
			`(min-height: 10rem) and (max-height: 29.999rem)`)
	}
	if q&Medium != 0 {
		apply(`(min-width: 30rem) and (max-width: 49.999rem)`,
			`(min-height: 30rem) and (max-height: 49.999rem)`)
	}
	if q&Large != 0 {
		apply(`(min-width: 50rem) and (max-width: 64.999rem)`,
			`(min-height: 50rem) and (max-height: 64.999rem)`)
	}
	if q&Huge != 0 {
		apply(`(min-width: 65rem)`,
			`(min-height: 65rem)`)
	}

	if len(s) == 0 {
		if q&Portrait != 0 && q&Landscape == 0 {
			s += "(orientation:portrait)"
		}
		if q&Landscape != 0 && q&Portrait == 0 {
			s += "(orientation:landscape)"
		}
	}

	return s
}

//GetBool implements client.Bool
func (q SizeQuery) GetBool() js.Bool {
	return js.Func("window.matchMedia").Call(client.NewString(q.Media())).Get("matches").GetBool()
}

//GetValue implements client.Value
func (q SizeQuery) GetValue() js.Value {
	return q.GetBool().Value
}
