package css

import (
	"fmt"
	"strings"

	"qlova.org/seed"
	"qlova.org/seed/client"
)

func init() {
	client.ClickCursor = SetCursor(Pointer)
}

type data struct {
	selector string

	rules *OrderedMap

	prefix, suffix *OrderedMap

	queries map[string]Rules

	sheets map[string]struct{}

	requires map[string]string
}

type ruleable interface {
	Rule() Rule
}

type Style interface {
	Rules() Rules
}

//Rule is a single css Rule.
type Rule string

func (r Rule) Rules() Rules {
	return Rules{r}
}

func (r Rule) Split() (property, value string) {
	split := strings.Split(string(r), ":")
	return split[0], split[1][:len(split[1])-1]
}

func (r Rule) Property() string {
	return strings.Split(string(r), ":")[0]
}

func (r Rule) Value() string {
	value := strings.Split(string(r), ":")[1]
	return value[:len(value)-1]
}

func (r Rule) AddTo(c seed.Seed) {
	var d data
	c.Load(&d)

	switch mode, q := client.Seed(c); mode {
	case client.AddTo:
		property, value := r.Split()
		if strings.HasPrefix(property, "-") {
			fmt.Fprintf(q, `%v.style.setProperty("%v", "%v");`, client.Element(c), property, value)
		} else {
			fmt.Fprintf(q, `%v.style.%v = "%v";`, client.Element(c), dashes2camels(property), value)
		}

	case client.Undo:
		property := r.Property()
		if strings.HasPrefix(property, "-") {
			fmt.Fprintf(q, `%v.style.setProperty("%v", "");`, client.Element(c), property)
		} else {
			fmt.Fprintf(q, `%v.style.%v = "";`, client.Element(c), dashes2camels(property))
		}
	default:
		if d.rules == nil {
			d.rules = NewOrderedMap()
		}
		property, value := r.Split()
		d.rules.Set(property, value)
	}

	c.Save(d)
}

func (r Rule) And(options ...seed.Option) seed.Option {
	return seed.And(r, options...)
}

//Important returns an important version of this rule.
func (r Rule) Important() Rule {
	return r[:len(r)-1] + " !important;"
}

//Selector returns the css selector of this seed.
func Selector(c seed.Seed) string {
	c.Use()

	var d data
	c.Load(&d)

	if d.selector != "" {
		return d.selector
	}

	return "#" + client.ID(c)
}

func dashes2camels(s string) string {
	var camel string
	var parts = strings.Split(s, "-")
	for i, part := range parts {
		if i == 0 {
			camel += part
		} else {
			camel += strings.Title(part)
		}
	}
	return camel
}

//SetSelector sets the CSS selector of this seed.
func SetSelector(selector string) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		var d data
		c.Load(&d)
		d.selector = selector
		c.Save(d)
	})
}

//Set sets the CSS property to be set to the given value.
func Set(property, value string) Rule {
	return Rule(property + ":" + value + ";")
}

//Add adds a stylesheet associated with this seed.
func Add(stylesheet string) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		switch mode, _ := client.Seed(c); mode {
		case client.AddTo, client.Undo:
			panic("css.Add must not be called on a client.Seed")
		}

		var d data
		c.Load(&d)

		if d.sheets == nil {
			d.sheets = make(map[string]struct{})
			c.Save(d)
		}

		d.sheets[stylesheet] = struct{}{}

	})
}
