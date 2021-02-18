package set

import "qlova.org/seed"

type data struct {
	queries map[string]string
}

//Query applies the given styles only when the given css query is valid.
func Query(q string, styles ...Style) seed.Option {
	return seed.Mutate(func(data *data) {
		if data.queries == nil {
			data.queries = make(map[string]string)
		}

		var rules string

		for _, style := range styles {
			for _, rule := range style.Rules() {
				rules += string(rule)
			}
		}

		data.queries["@media "+q] = rules
	})
}
