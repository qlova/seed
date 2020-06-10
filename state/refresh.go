package state

import (
	"fmt"
	"io/ioutil"
	"reflect"

	"github.com/qlova/seed"
	"github.com/qlova/seed/asset"
	"github.com/qlova/seed/html"
	"github.com/qlova/seed/html/attr"
	"github.com/qlova/seed/js"
	"github.com/qlova/seed/script"
)

func If(condition js.AnyBool, options ...seed.Option) seed.Option {
	if condition == nil {
		return seed.Options(options)
	}

	return seed.NewOption(func(c seed.Seed) {

		//Add any children seeds to the parent seed.
		//Hacky fix.
		for _, o := range options {
			if o == nil {
				continue
			}
			switch child := o.(type) {
			case seed.Seed:
				c.With(child)
			}
		}

		c.With(OnRefresh(func(q script.Ctx) {
			q.If(condition, func(q script.Ctx) {
				for _, option := range options {
					if option == nil {
						continue
					}
					if other, ok := option.(seed.Seed); ok {
						script.Scope(other, q).AddTo(script.Scope(c, q))
					} else {
						option.AddTo(script.Scope(c, q))
					}
				}
			}).Else(func(q script.Ctx) {
				for _, option := range options {
					if option == nil {
						continue
					}
					if other, ok := option.(seed.Seed); ok {
						script.Scope(c, q).Undo(script.Scope(other, q))
					} else {
						script.Scope(c, q).Undo(option)
					}
				}
			})
		}))
	})
}

//Refresh triggers a state refresh of the seed and any of its children.
func Refresh(c seed.Seed) script.Script {
	var d data
	c.Read(&d)
	d.refresh = true
	c.Write(d)

	return js.Func("await c.r").Run(script.Q, js.NewString(script.ID(c)))
}

func AdoptRefresh(c seed.Seed) script.Script {
	var d data
	c.Read(&d)

	var refresh script.Script

	if d.onrefresh != nil {
		refresh = refresh.Append(js.Run(js.Func("c.or"), js.NewValue("q"), js.NewString(script.ID(c)), js.NewFunction(d.onrefresh)))
	}

	d.onrefresh = nil
	d.refresh = false
	c.Write(d)

	for _, child := range c.Children() {
		refresh = refresh.Append(AdoptRefresh(child))
	}

	return refresh
}

func AdoptRefreshOfChildren(c seed.Seed) script.Script {
	var refresh script.Script

	for _, child := range c.Children() {
		refresh = refresh.Append(AdoptRefresh(child))
	}

	return refresh
}

//OnRefresh is called whenever this seed has its state refreshed.
func OnRefresh(do script.Script) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		do(js.NewCtx(ioutil.Discard, c)) //Catch errors and harvest pages.

		var d data
		c.Read(&d)
		d.refresh = true
		d.onrefresh = d.onrefresh.Append(do)
		c.Write(d)
	})
}

//SetText sets the text of the seed based on the argument provided.
func SetText(text AnyString) seed.Option {
	switch t := text.(type) {
	case string:
		return html.SetInnerText(t)

	case String:
		return t.SetText()

	case js.AnyString:
		return seed.NewOption(func(c seed.Seed) {
			c.With(OnRefresh(func(q script.Ctx) {
				q(fmt.Sprintf(`%v.innerText = %v || "";`,
					script.Scope(c, q).Element(), t.GetString().String()))
			}))
		})
	case js.AnyValue:
		return seed.NewOption(func(c seed.Seed) {
			c.With(OnRefresh(func(q script.Ctx) {
				q(fmt.Sprintf(`%v.innerText = %v || "";`,
					script.Scope(c, q).Element(), t.GetValue().String()))
			}))
		})
	default:
		panic("unsupported AnyString argument " + reflect.TypeOf(t).String())
	}

	return seed.NewOption(func(c seed.Seed) {})
}

//SetSource sets the source of the seed based on the argument provided.
func SetSource(src AnyString) seed.Option {
	switch t := src.(type) {
	case string:

		t = asset.Path(t)

		return seed.Options{
			attr.Set("src", t),
			attr.Set("alt", t),

			asset.New(t),
		}

	case String:
		return t.SetSource()

	case js.AnyString:
		return seed.NewOption(func(c seed.Seed) {
			c.With(OnRefresh(func(q script.Ctx) {
				q(fmt.Sprintf(`%v.src = %v || "";`,
					script.Scope(c, q).Element(), t.GetString().String()))
			}))
		})
	case js.AnyValue:
		return seed.NewOption(func(c seed.Seed) {
			c.With(OnRefresh(func(q script.Ctx) {
				q(fmt.Sprintf(`%v.src = %v || "";`,
					script.Scope(c, q).Element(), t.GetValue().String()))
			}))
		})
	default:
		panic("unsupported AnyString argument " + reflect.TypeOf(t).String())
	}

	return seed.NewOption(func(c seed.Seed) {})
}
