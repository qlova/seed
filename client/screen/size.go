//Package screen provides client booleans for different screen sizes. Similar to a media query.
package screen

import (
	"fmt"

	"qlova.org/seed/client"
	"qlova.org/seed/use/js"
)

//SizeQuery is a general screen-size query.
type SizeQuery uint8

//Query types.
const (
	//Base screen sizes.
	Tiny SizeQuery = 1 << iota
	Small
	Medium
	Large
	Huge

	//Orientation is applied with the bitwise ^ operator.
	Orientation
	Portrait  = Orientation | 1<<(6)
	Landscape = Orientation | 1<<(7)

	//Screen size ranges.
	TinyToSmall  = Tiny | Small
	TinyToMedium = TinyToSmall | Medium
	TinyToLarge  = TinyToMedium | Large
	TinyToHuge   = TinyToLarge | Huge

	SmallToMedium = Small | Medium
	SmallToLarge  = SmallToMedium | Large
	SmallToHuge   = SmallToLarge | Huge

	MediumToLarge = Medium | Large
	MediumToHuge  = MediumToLarge | Huge

	LargeToHuge = Large | Huge

	//Inversions
	NotTiny   = ^Tiny
	NotSmall  = ^Small
	NotMedium = ^Medium
	NotLarge  = ^Large
	NotHuge   = ^Huge

	NotTinyToSmall  = ^TinyToSmall
	NotTinyToMedium = ^TinyToMedium
	NotTinyToLarge  = ^TinyToLarge
	NotTinyToHuge   = ^TinyToHuge

	NotSmallToMedium = ^SmallToMedium
	NotSmallToLarge  = ^SmallToLarge
	NotSmallToHuge   = ^SmallToHuge

	NotMediumToLarge = ^MediumToLarge
	NotMediumToHuge  = ^MediumToHuge

	NotLargeToHuge = ^LargeToHuge
)

//Not returns the inverse of the query.
func Not(q SizeQuery) SizeQuery { return ^q }

//Media returns the query as a CSS media query.
func (q SizeQuery) Media() string {

	fmt.Printf("p %v l %v o %v", q&Portrait, q&Landscape, q&Orientation)

	var s string

	var p, l bool
	if (q&Portrait != 0) == (q&Landscape != 0) {
		p = true
		l = true
	}

	if (q&Portrait != 0) == (q&Orientation != 0) {
		p = true
	}

	if (q&Landscape != 0) == (q&Orientation != 0) {
		l = true
	}

	fmt.Println(p, l)

	apply := func(portrait, landscape string) {

		if p {
			if len(s) > 0 {
				s += ","
			}
			s += "screen and " + portrait +
				" and (orientation:portrait)"
		}

		if l {
			if len(s) > 0 {
				s += ","
			}
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

	fmt.Println(s, q&Portrait, q&Landscape)

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
