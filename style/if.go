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

//Condition backs the If variable and can apply conditional styles.
type Condition struct {
	query, portrait, landscape string
	styles                     []Style
}

func (c Condition) AddTo(s seed.Any) {
	data := seeds[s.Root()]
	if data.queries == nil {
		data.queries = make(map[string]string)
	}

	var rules string

	for _, style := range c.styles {
		for _, rule := range style.Rules() {
			rules += string(rule)
		}
	}

	var q = `@media`
	if c.portrait != "" {
		q += ` screen and ` + c.portrait + ` and (orientation:portrait)`
		if c.landscape != "" {
			q += ","
		}
	}
	if c.landscape != "" {
		q += ` screen and ` + c.landscape + ` and (orientation:landscape)`
	}
	if c.portrait == "" && c.landscape == "" {
		q += ` screen and ` + c.query
	}

	data.queries[q] = rules

	seeds[s.Root()] = data
}

func (c Condition) Apply(seed.Ctx) {
	panic("cannot apply a style.Condition")
}

func (c Condition) Reset(seed.Ctx) {
	panic("cannot reset a style.Condition")
}

func (c Condition) And(options ...seed.Option) seed.Option {
	return seed.And(c, options...)
}

//Tiny applies the styles on tiny screens. ie. SmartWatches.
func (c Condition) Tiny(styles ...Style) Condition {
	return Condition{
		portrait:  `(max-width: 9.9999rem)`,
		landscape: `(max-height: 9.9999rem)`,
		styles:    styles,
	}
}

func (c Condition) Small(styles ...Style) Condition {
	return Condition{
		portrait:  `(min-width: 10rem) and (max-width: 24.999rem)`,
		landscape: `(min-height: 10rem) and (max-height: 24.999rem)`,
		styles:    styles,
	}
}

func (c Condition) Medium(styles ...Style) Condition {
	return Condition{
		portrait:  `(min-width: 25rem) and (max-width: 49.999rem)`,
		landscape: `(min-height: 25rem) and (max-height: 49.999rem)`,
		styles:    styles,
	}
}

func (c Condition) Large(styles ...Style) Condition {
	return Condition{
		portrait:  `(min-width: 50rem) and (max-width: 64.999rem)`,
		landscape: `(min-height: 50rem) and (max-height: 64.999rem)`,
		styles:    styles,
	}
}

func (c Condition) Huge(styles ...Style) Condition {
	return Condition{
		portrait:  `(min-width: 65rem)`,
		landscape: `(min-height: 65rem)`,
		styles:    styles,
	}
}

func (c Condition) Portrait(styles ...Style) Condition {
	c.query = `(orientation: portrait)`
	c.landscape = ""
	c.styles = append(c.styles, styles...)
	return c
}

func (c Condition) Landscape(styles ...Style) Condition {
	c.query = `(orientation: landscape)`
	c.portrait = ""
	c.styles = append(c.styles, styles...)
	return c
}

//If allows conditional styles.
var If Condition
