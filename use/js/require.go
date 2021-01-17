package js

import (
	"qlova.org/seed"
)

type data struct {
	requires map[string]string
}

//Require js script.
func Require(path string, contents string) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		var d data
		c.Load(&d)

		if d.requires == nil {
			d.requires = make(map[string]string)
		}

		d.requires[path] = contents

		c.Save(d)
	})
}

func scripts(c seed.Seed, fill map[string]string) {
	var data data
	c.Load(&data)

	if data.requires != nil {
		for path, contents := range data.requires {
			fill[path] = contents
		}
	}

	for _, child := range c.Children() {
		scripts(child, fill)
	}
}

//Scripts returns the external scripts needed by this seed.
func Scripts(root seed.Seed) map[string]string {
	result := make(map[string]string)
	scripts(root, result)
	return result
}
