package seed

import (
	"bytes"

	"github.com/qlova/seed/internal"
	"github.com/qlova/seed/style"
	"github.com/qlova/seed/style/css"
)

//Shadow is an alias to the style.Shadow type.
type Shadow = style.Shadow

//Gradient is an alias to the style.Gradient type.
type Gradient = style.Gradient

//Frame is an animation frame.
type Frame = internal.Frame

//Animation is a change in styles across frames.
type Animation = internal.Animation

//SetAnimation sets the animation of this seed to be looping and 1 second long.
func (seed Seed) SetAnimation(animation Animation) {
	seed.animation = animation
	seed.CSS().SetAnimationName(css.AnimationName(seed.id))
	seed.CSS().SetAnimationDuration(css.Time(1))
	seed.CSS().SetAnimationIterationCount(css.Infinite)
}

type sheet struct {
	style.Sheet

	Tiny, Small, Medium, Large, Huge style.Sheet

	smallUp, mediumUp, largeUp,
	smallDown, mediumDown, largeDown,
	smallToMedium, mediumToLarge, smallToLarge *style.Sheet
}

type query struct {
	smallUp, mediumUp, largeUp,
	smallDown, mediumDown, largeDown,
	smallToMedium, mediumToLarge, smallToLarge *style.Group
}

func (s Seed) TinyToSmall() style.Group {
	if s.query == nil {
		s.query = new(query)
	}
	var sheet = style.NewGroup()
	s.query.smallDown = sheet
	return *sheet
}

func (s Seed) MediumToHuge() style.Group {
	if s.query == nil {
		s.query = new(query)
	}
	var sheet = style.NewGroup()
	s.query.mediumUp = sheet
	return *sheet
}

func newSheet() sheet {
	return sheet{
		Sheet:  style.NewSheet(),
		Tiny:   style.NewSheet(),
		Small:  style.NewSheet(),
		Medium: style.NewSheet(),
		Large:  style.NewSheet(),
		Huge:   style.NewSheet(),
	}
}

func (s *sheet) AddSeed(selector string, seed Seed) {
	s.AddGroup(selector, seed.Group)
	s.Tiny.AddGroup(selector, seed.Tiny)
	s.Small.AddGroup(selector, seed.Small)
	s.Medium.AddGroup(selector, seed.Medium)
	s.Large.AddGroup(selector, seed.Large)
	s.Huge.AddGroup(selector, seed.Huge)
	if seed.query != nil {
		var q = seed.query

		if q.mediumUp != nil {
			if s.mediumUp == nil {
				var sheet = style.NewSheet()
				s.mediumUp = &sheet
			}
			s.mediumUp.AddGroup(selector, *q.mediumUp)
		}

		if q.smallDown != nil {
			if s.smallDown == nil {
				var sheet = style.NewSheet()
				s.smallDown = &sheet
			}
			s.smallDown.AddGroup(selector, *q.smallDown)
		}

	}
}

const TinyMin = `10rem`
const TinyMax = `20rem`
const SmallMin = `40rem`
const SmallMax = `60rem`
const MediumMin = `60rem`
const MediumMax = `80rem`
const LargeMin = `80rem`
const LargeMax = `100rem`

func (s sheet) writeQueryTo(buffer *bytes.Buffer) {
	if s.mediumUp != nil {
		buffer.WriteString(`@media screen and (` +
			`min-width: ` + SmallMin +
			`) and (` +
			`min-height: ` + SmallMin +

			`), screen and (` +

			`min-width: ` + TinyMin +
			`) and (` +
			`min-height: ` + SmallMax +
			`) and (` +
			`max-width: ` + MediumMin +

			`), screen and (` +

			`min-height: ` + TinyMin +
			`) and (` +
			`min-width: ` + SmallMax +
			`) and (` +
			`max-height: ` + MediumMin +

			`), screen and (` +

			`min-width:  ` + SmallMax +
			`) and (` +
			`min-height: ` + SmallMin +
			`) {`)
		buffer.Write(s.mediumUp.Bytes())
		buffer.WriteString(`}`)
	}

	if s.smallDown != nil {
		buffer.WriteString(`@media screen and (` +
			`min-height: ` + TinyMax +
			`) and (` +
			`max-width: ` + SmallMin +
			`) and (` +
			`max-height: ` + SmallMax +

			`), screen and (` +
			`min-width:  ` + TinyMax +
			`) and (` +
			`max-height: ` + SmallMin +
			`) and (` +
			`max-width: ` + SmallMax +
			`) {`)
		buffer.Write(s.smallDown.Bytes())
		buffer.WriteString(`}`)
	}
}

func (s sheet) Bytes() []byte {
	var buffer bytes.Buffer

	buffer.Write(s.Sheet.Bytes())

	s.writeQueryTo(&buffer)

	buffer.WriteString(`@media screen and (max-width: ` + TinyMax + `) and (max-height: ` + TinyMax +
		`), screen and (max-width: ` + TinyMin + `), screen and (max-height: ` + TinyMin + `) {`)
	buffer.Write(s.Tiny.Bytes())
	buffer.WriteString(`}`)

	buffer.WriteString(`@media screen and (` +
		`min-width: ` + TinyMin +
		`) and (` +
		`min-height: ` + TinyMax +
		`) and (` +
		`max-width: ` + SmallMin +
		`) and (` +
		`max-height: ` + SmallMax +

		`), screen and (` +
		`min-height: ` + TinyMin +
		`) and (` +
		`min-width:  ` + TinyMax +
		`) and (` +
		`max-height: ` + SmallMin +
		`) and (` +
		`max-width: ` + SmallMax +
		`) {`)
	buffer.Write(s.Small.Bytes())
	buffer.WriteString(`}`)

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
