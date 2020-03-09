package style

import "github.com/qlova/seed"

const TinyMin = `10rem`
const TinyMax = `20rem`
const SmallMin = `40rem`
const SmallMax = `60rem`
const MediumMin = `60rem`
const MediumMax = `80rem`
const LargeMin = `80rem`
const LargeMax = `100rem`

type data struct {
	queries map[string]string
}

var seeds = make(map[seed.Seed]data)

//Conditional backs the If variable and can apply conditional styles.
type Conditional struct{}

func (c Conditional) do(query string, styles ...Style) seed.Option {
	return seed.NewOption(func(c seed.Any) {
		data := seeds[c.Root()]
		if data.queries == nil {
			data.queries = make(map[string]string)
		}

		var rules string

		for _, style := range styles {
			for _, rule := range style.Rules() {
				rules += string(rule)
			}
		}

		data.queries[query] = rules

		seeds[c.Root()] = data
	}, nil, nil)
}

//Tiny applies the styles on tiny screens. ie. SmartWatches.
func (c Conditional) Tiny(styles ...Style) seed.Option {
	return c.do(`@media screen and (max-width: `+TinyMax+`) and (max-height: `+TinyMax+
		`), screen and (max-width: `+TinyMin+`), screen and (max-height: `+TinyMin+`)`, styles...)
}

func (c Conditional) Small(styles ...Style) seed.Option {
	return c.do(`@media screen and (`+
		`min-width: `+TinyMin+
		`) and (`+
		`min-height: `+TinyMax+
		`) and (`+
		`max-width: `+SmallMin+
		`) and (`+
		`max-height: `+SmallMax+

		`), screen and (`+
		`min-height: `+TinyMin+
		`) and (`+
		`min-width:  `+TinyMax+
		`) and (`+
		`max-height: `+SmallMin+
		`) and (`+
		`max-width: `+SmallMax+
		`)`, styles...)
}

func (c Conditional) Medium(styles ...Style) seed.Option {
	return c.do(`@media screen and (`+
		`min-width: `+SmallMin+
		`) and (`+
		`min-height: `+SmallMin+
		`) and (`+
		`max-width: `+MediumMin+
		`) and (`+
		`max-height: `+MediumMax+

		`), screen and (`+

		`min-width: `+TinyMin+
		`) and (`+
		`min-height: `+SmallMax+
		`) and (`+
		`max-width: `+MediumMin+
		`) and (`+
		`max-height: `+MediumMax+

		`), screen and (`+

		`min-height: `+TinyMin+
		`) and (`+
		`min-width: `+SmallMax+
		`) and (`+
		`max-width: `+MediumMax+
		`) and (`+
		`max-height: `+MediumMin+

		`), screen and (`+

		`min-width:  `+SmallMax+
		`) and (`+
		`min-height: `+SmallMin+
		`) and (`+
		`max-height: `+MediumMin+
		`) and (`+
		`max-width: `+MediumMax+
		`)`, styles...)
}

func (c Conditional) Large(styles ...Style) seed.Option {
	return c.do(`@media screen and (`+

		`min-width: `+TinyMin+
		`) and (`+
		`min-height: `+MediumMax+
		`) and (`+
		`max-width: `+LargeMin+
		`) and (`+
		`max-height: `+LargeMax+

		`), screen and (`+

		`min-height: `+TinyMin+
		`) and (`+
		`min-width:  `+MediumMax+``+
		`) and (`+
		`max-height: `+LargeMin+
		`) and (`+
		`max-width: `+LargeMax+
		`)`, styles...)
}

func (c Conditional) Huge(styles ...Style) seed.Option {
	return c.do(`@media screen and (`+

		`min-width: `+TinyMin+
		`) and (`+
		`min-height: `+LargeMax+

		`), screen and (`+

		`min-height: `+TinyMin+
		`) and (`+
		`min-width:  `+LargeMax+``+
		`)`, styles...)
}

func (c Conditional) Portrait(styles ...Style) seed.Option {
	return c.do(`@media screen and (orientation: portrait)`, styles...)
}

func (c Conditional) Landscape(styles ...Style) seed.Option {
	return c.do(`@media screen and (orientation: landscape)`, styles...)
}

//If allows conditional styles.
var If Conditional
