package assets

import (
	"qlova.org/seed"
	"qlova.org/seed/asset/inbed"
)

func init() {
	inbed.PackageName = "assets"
	inbed.ImporterName = "assets.go"
}

type data struct {
	

	assets []string
}

//New creates a new asset and adds it to the seed.
func New(src string) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		if src != "" {
			var data data
			c.Load(&data)
			data.assets = append(data.assets, src)
			c.Save(data)
		}
	})
}

func of(c seed.Seed, addto map[string]bool) {

	var data data
	c.Load(&data)

	for _, asset := range data.assets {
		addto[asset] = true
	}

	for _, child := range c.Children() {
		of(child, addto)
	}

	return
}

//Of returns the assets of the given seed.
func Of(c seed.Seed) map[string]bool {
	var result = make(map[string]bool)

	of(c, result)

	return result
}
