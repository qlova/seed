//Package poll provides ways to repeatedly call a function at a given interval.
package poll

import (
	"time"

	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/use/html"
	"qlova.org/seed/use/js"
	"qlova.org/seed/use/js/window"
)

//When the condition is true, poll is called at the given interval.
func When(condition client.Bool, interval time.Duration, poll interface{}) seed.Option {
	return client.OnLoad(window.SetInterval(client.If(condition, client.Go(poll)).GetScript(), js.NewNumber(float64(interval.Milliseconds()))))
}

//Every calls pol at every interval as long as the seed is visible.
//If you want pol to be called regardless if the seed is visible, use When(client.True, interval, poll)
func Every(interval time.Duration, poll interface{}) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		When(html.Element(c).Get("offsetParent"), interval, poll).AddTo(c)
	})
}
