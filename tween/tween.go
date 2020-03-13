package tween

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/html"
	"github.com/qlova/seed/html/attr"
	"github.com/qlova/seed/script"
)

//This adds tweening to this seed which can be enabled with the Tween() function.
func This() seed.Option {
	return seed.Do(func(c seed.Seed) {
		c.Add(
			script.Require("/flipping.js", js),

			attr.Set("data-flip-key", html.ID(c)),
		)
	})
}

//Track tracks this seed with a key, only one seed of any given key should be visible at a time.
func Track(key string) seed.Option {
	return seed.Do(func(c seed.Seed) {
		c.Add(
			script.Require("/flipping.js", js),

			attr.Set("data-flip-key", key),
		)
	})
}

//Tween attempts to tween any elements with This() options that have changed position, scale or rotation.
func Tween(s script.Script) script.Script {
	return func(q script.Ctx) {
		q.Javascript(`try { flipping.read(); } catch(error) {}`)
		s(q)
		q.Javascript(`try { flipping.flip(); } catch(error) {}`)
	}
}
