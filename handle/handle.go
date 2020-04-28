package handle

import (
	"github.com/qlova/seed"
)

//Data stores associations between seeds and event handlers.
type Data struct {
	seed.Data

	On map[string]func()
}

func On(event string, f func()) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		var data Data
		c.Read(&data)

		if data.On == nil {
			data.On = make(map[string]func())
			c.Write(data)
		}

		data.On[event] = f
	})
}

func OnClick(f func()) seed.Option {
	return On("click", f)
}
