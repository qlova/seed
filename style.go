package seed

import (
	"bytes"

	"github.com/qlova/seed/style"
)

//Shadow is an alias to the style.Shadow type.
type Shadow = style.Shadow

//Gradient is an alias to the style.Gradient type.
type Gradient = style.Gradient

type sheet struct {
	style.Sheet

	Tiny, Small, Medium, Large, Huge style.Sheet
}

func newSheet() sheet {
	return sheet{
		style.NewSheet(),
		style.NewSheet(),
		style.NewSheet(),
		style.NewSheet(),
		style.NewSheet(),
		style.NewSheet(),
	}
}

func (s *sheet) AddSeed(selector string, seed Seed) {
	s.AddGroup(selector, seed.Group)
	s.Tiny.AddGroup(selector, seed.Tiny)
	s.Small.AddGroup(selector, seed.Small)
	s.Medium.AddGroup(selector, seed.Medium)
	s.Large.AddGroup(selector, seed.Large)
	s.Huge.AddGroup(selector, seed.Huge)
}

func (s sheet) Bytes() []byte {
	var buffer bytes.Buffer

	buffer.Write(s.Sheet.Bytes())

	const TinyMin = `10rem`
	const TinyMax = `20rem`

	buffer.WriteString(`@media screen and (max-width: ` + TinyMax + `) and (max-height: ` + TinyMax +
		`), screen and (max-width: ` + TinyMin + `), screen and (max-height: ` + TinyMin + `) {`)
	buffer.Write(s.Tiny.Bytes())
	buffer.WriteString(`}`)

	const SmallMin = `40rem`
	const SmallMax = `60rem`

	buffer.WriteString(`@media screen and (min-width: ` + TinyMin + `) and (min-height: ` + TinyMax +
		`) and (max-width: ` + SmallMin + `) and (max-height: ` + SmallMax +
		`), screen and (min-height: ` + TinyMin + `) and (min-width:  ` + TinyMax +
		`) and (max-height: ` + SmallMin + `) and (max-width: ` + SmallMax + `) {`)
	buffer.Write(s.Small.Bytes())
	buffer.WriteString(`}`)

	const MediumMin = `60rem`
	const MediumMax = `80rem`

	buffer.WriteString(`@media screen and (` +
		`min-width: ` + SmallMin +
		`) and (` +
		`min-height: ` + SmallMin +
		`) and (` +
		`max-width: ` + MediumMin +
		`) and (` +
		`max-height: ` + MediumMax +

		`), screen and (` +

		`min-width: ` + TinyMin +
		`) and (` +
		`min-height: ` + SmallMax +
		`) and (` +
		`max-width: ` + MediumMin +
		`) and (` +
		`max-height: ` + MediumMax +

		`), screen and (` +

		`min-height: ` + TinyMin +
		`) and (` +
		`min-width: ` + SmallMax +
		`) and (` +
		`max-width: ` + MediumMax +
		`) and (` +
		`max-height: ` + MediumMin +

		`), screen and (` +

		`min-width:  ` + SmallMax +
		`) and (` +
		`min-height: ` + SmallMin +
		`) and (` +
		`max-height: ` + MediumMin +
		`) and (` +
		`max-width: ` + MediumMax +
		`) {`)
	buffer.Write(s.Medium.Bytes())
	buffer.WriteString(`}`)

	const LargeMin = `80rem`
	const LargeMax = `100rem`

	buffer.WriteString(`@media screen and (` +

		`min-width: ` + TinyMin +
		`) and (` +
		`min-height: ` + MediumMax +
		`) and (` +
		`max-width: ` + LargeMin +
		`) and (` +
		`max-height: ` + LargeMax +

		`), screen and (` +

		`min-height: ` + TinyMin +
		`) and (` +
		`min-width:  ` + MediumMax + `` +
		`) and (` +
		`max-height: ` + LargeMin +
		`) and (` +
		`max-width: ` + LargeMax +
		`) {`)
	buffer.Write(s.Large.Bytes())
	buffer.WriteString(`}`)

	buffer.WriteString(`@media screen and (` +

		`min-width: ` + TinyMin +
		`) and (` +
		`min-height: ` + LargeMax +

		`), screen and (` +

		`min-height: ` + TinyMin +
		`) and (` +
		`min-width:  ` + LargeMax + `` +
		`) {`)
	buffer.Write(s.Huge.Bytes())
	buffer.WriteString(`}`)

	/*buffer.WriteString(`@media screen and (max-width: 20rem) and (max-height: 20rem), screen and (max-width: 10rem), screen and (max-height: 10rem) {`)
	buffer.Write(s.Medium.Bytes())
	buffer.WriteString(`}`)

	buffer.WriteString(`@media screen and (max-width: 20rem) and (max-height: 20rem), screen and (max-width: 10rem), screen and (max-height: 10rem) {`)
	buffer.Write(s.Large.Bytes())
	buffer.WriteString(`}`)

	buffer.WriteString(`@media screen and (max-width: 20rem) and (max-height: 20rem), screen and (max-width: 10rem), screen and (max-height: 10rem) {`)
	buffer.Write(s.Huge.Bytes())
	buffer.WriteString(`}`)*/

	return buffer.Bytes()
}
