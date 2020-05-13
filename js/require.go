package js

import (
	"github.com/qlova/seed"
)

type data struct {
	seed.Data

	requires map[string]string
}

//Require js script.
func Require(path string, contents string) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		c.Use()

		var d data
		c.Read(&d)

		if d.requires == nil {
			d.requires = make(map[string]string)
		}

		d.requires[path] = contents

		c.Write(d)
	})
}

func scripts(c seed.Seed, fill map[string]string) {
	var data data
	c.Read(&data)

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
