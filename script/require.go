package script

import (
	"github.com/qlova/seed"
)

//Require js script.
func Require(path string, contents string) seed.Option {
	return seed.Do(func(s seed.Seed) {
		s.Root().Use()
		data := seeds[s.Root()]
		if data.requires == nil {
			data.requires = make(map[string]string)
			seeds[s.Root()] = data
		}

		data.requires[path] = contents
	})
}

func scripts(c seed.Seed, fill map[string]string) {
	data := seeds[c]
	if data.requires != nil {
		for path, contents := range data.requires {
			fill[path] = contents
		}
	}

	for _, child := range c.Children() {
		scripts(child.Root(), fill)
	}
}

//Scripts returns the external scripts needed by this seed.
func Scripts(root seed.Any) map[string]string {
	result := make(map[string]string)
	scripts(root.Root(), result)
	return result
}
