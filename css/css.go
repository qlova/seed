package css

import (
	"encoding/base64"
	"fmt"
	"math/big"
	"strings"

	"github.com/qlova/seed"
)

type AnyRule interface {
	Rule() Rule
}

//Rule is a single css Rule.
type Rule string

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

func (r Rule) AddTo(c seed.Any) {
	data := seeds[c.Root()]
	if data.rules == nil {
		data.rules = make(rules)
	}
	property, value := r.Split()
	data.rules[property] = value
	seeds[c.Root()] = data
}

func (r Rule) Apply(c seed.Ctx) {
	property, value := r.Split()
	fmt.Fprintf(c.Ctx, `%v.style.%v = "%v";`, c.Element(), dashes2camels(property), value)
}

func (r Rule) Reset(c seed.Ctx) {
	property := r.Property()
	fmt.Fprintf(c.Ctx, `%v.style.%v = "";`, c.Element(), dashes2camels(property))
}

func (r Rule) And(options ...seed.Option) seed.Option {
	return seed.And(r, options...)
}

//Important returns an important version of this rule.
func (r Rule) Important() Rule {
	return r[:len(r)-1] + " !important;"
}

type data struct {
	selector string

	rules

	prefix, suffix rules

	queries map[string]*data
}

var seeds = make(map[seed.Seed]data)

type rules map[string]string

//Selector returns the css selector of this seed.
func Selector(c seed.Any) string {
	c.Root().Use()
	data := seeds[c.Root()]
	if data.selector != "" {
		return data.selector
	}

	return "#" + base64.RawURLEncoding.EncodeToString(big.NewInt(int64(c.Root())).Bytes())
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
	return seed.NewOption(func(s seed.Any) {
		data := seeds[s.Root()]
		data.selector = selector
		seeds[s.Root()] = data
	}, func(s seed.Ctx) {
		panic("css.SetSelector must be called at buildtime")
	}, func(s seed.Ctx) {
		panic("css.SetSelector must be called at buildtime")
	})
}

//Set sets the CSS property to be set to the given value.
func Set(property, value string) Rule {
	return Rule(property + ":" + value + ";")
}
