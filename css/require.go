package css

import (
	"qlova.org/seed"
)

//Require js script.
func Require(path string, contents string) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		var d data
		c.Read(&d)

		if d.requires == nil {
			d.requires = make(map[string]string)
		}

		d.requires[path] = contents

		c.Write(d)
	})
}

func styles(c seed.Seed, fill map[string]string) {
	var data data
	c.Read(&data)

	if data.requires != nil {
		for path, contents := range data.requires {
			fill[path] = contents
		}
	}

	for _, child := range c.Children() {
		styles(child, fill)
	}
}

//Stylesheets returns the external stylesheets needed by this seed.
func Stylesheets(root seed.Seed) map[string]string {
	result := make(map[string]string)
	styles(root, result)
	return result
}
