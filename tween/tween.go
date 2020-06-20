package tween

import (
	"qlova.org/seed"
	"qlova.org/seed/css"
	"qlova.org/seed/html"
	"qlova.org/seed/html/attr"
	"qlova.org/seed/js"
	"qlova.org/seed/script"
)

//Auto tween.
func Auto() seed.Option {
	return css.Set("transition", "all 0.4s linear")
}

//This adds tweening to this seed which can be enabled with the Tween() function.
func This() seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		c.With(
			js.Require("/flipping.js", flipping),

			attr.Set("data-flip-key", html.ID(c)),
		)
	})
}

//DisableScaling stops scale being applied during the tween.
func DisableScaling() seed.Option {
	return attr.Set("data-flip-no-scale", "")
}

//Track tracks this seed with a key, only one seed of any given key should be visible at a time.
func Track(key string) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		c.With(
			js.Require("/flipping.js", flipping),

			attr.Set("data-flip-key", key),
		)
	})
}

//Tween attempts to tween any elements with This() options that have changed position, scale or rotation.
func Tween(s script.Script) script.Script {
	return func(q script.Ctx) {
		q(`try { flipping.read(); } catch(error) { seed.report(error) }`)
		s(q)
		q(`try { flipping.flip(); } catch(error) {seed.report(error) }`)
	}
}
