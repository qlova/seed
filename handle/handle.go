package handle

import (
	"qlova.org/seed"
)

//Data stores associations between seeds and event handlers.
type Data struct {
	

	On map[string]func()
}

func On(event string, f func()) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		var data Data
		c.Load(&data)

		if data.On == nil {
			data.On = make(map[string]func())
			c.Save(data)
		}

		data.On[event] = f
	})
}

func OnClick(f func()) seed.Option {
	return On("click", f)
}
