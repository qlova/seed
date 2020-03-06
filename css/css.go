package css

import (
	"encoding/base64"
	"fmt"
	"math/big"
	"strings"

	"github.com/qlova/seed"
)

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
func Set(property, value string) seed.Option {
	return seed.NewOption(func(s seed.Any) {
		data := seeds[s.Root()]
		if data.rules == nil {
			data.rules = make(rules)
		}
		data.rules[property] = value
		seeds[s.Root()] = data
	}, func(s seed.Ctx) {
		fmt.Fprintf(s.Ctx, `%v.style.%v = "%v";`, s.Element(), dashes2camels(property), value)
	}, func(s seed.Ctx) {
		fmt.Fprintf(s.Ctx, `%v.style.%v = "";`, s.Element(), dashes2camels(property))
	})
}
