package style

import "github.com/qlova/seed"

type data struct {
	seed.Data

	queries map[string]string
}

var seeds = make(map[seed.Seed]data)

//Condition backs the If variable and can apply conditional styles.
type Condition struct {
	query, portrait, landscape string
	styles                     []Style
}

func (con Condition) AddTo(c seed.Seed) {
	var data data
	c.Read(&data)

	if data.queries == nil {
		data.queries = make(map[string]string)
	}

	var rules string

	for _, style := range con.styles {
		for _, rule := range style.Rules() {
			rules += string(rule)
		}
	}

	var q = `@media`
	if con.portrait != "" {
		q += ` screen and ` + con.portrait + ` and (orientation:portrait)`
		if con.landscape != "" {
			q += ","
		}
	}
	if con.landscape != "" {
		q += ` screen and ` + con.landscape + ` and (orientation:landscape)`
	}
	if con.portrait == "" && con.landscape == "" {
		q += ` screen and ` + con.query
	}

	data.queries[q] = rules

	c.Write(data)
}

func (con Condition) And(options ...seed.Option) seed.Option {
	return seed.And(con, options...)
}

//Tiny applies the styles on tiny screens. ie. SmartWatches.
func (con Condition) Tiny(styles ...Style) Condition {
	return Condition{
		portrait:  `(max-width: 9.9999rem)`,
		landscape: `(max-height: 9.9999rem)`,
		styles:    styles,
	}
}

func (con Condition) Small(styles ...Style) Condition {
	return Condition{
		portrait:  `(min-width: 10rem) and (max-width: 29.999rem)`,
		landscape: `(min-height: 10rem) and (max-height: 29.999rem)`,
		styles:    styles,
	}
}

func (con Condition) Medium(styles ...Style) Condition {
	return Condition{
		portrait:  `(min-width: 30rem) and (max-width: 49.999rem)`,
		landscape: `(min-height: 30rem) and (max-height: 49.999rem)`,
		styles:    styles,
	}
}

func (con Condition) Large(styles ...Style) Condition {
	return Condition{
		portrait:  `(min-width: 50rem) and (max-width: 64.999rem)`,
		landscape: `(min-height: 50rem) and (max-height: 64.999rem)`,
		styles:    styles,
	}
}

func (con Condition) Huge(styles ...Style) Condition {
	return Condition{
		portrait:  `(min-width: 65rem)`,
		landscape: `(min-height: 65rem)`,
		styles:    styles,
	}
}

func (con Condition) Portrait(styles ...Style) Condition {
	con.query = `(orientation: portrait)`
	con.landscape = ""
	con.styles = append(con.styles, styles...)
	return con
}

func (con Condition) Landscape(styles ...Style) Condition {
	con.query = `(orientation: landscape)`
	con.portrait = ""
	con.styles = append(con.styles, styles...)
	return con
}

//If allows conditional styles.
var If Condition
