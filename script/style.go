package script

import qlova "github.com/qlova/script"

func (seed Seed) Hidden() qlova.Bool {
	return seed.Q.Value(`(getComputedStyle(` + seed.Element() + `, null).display == "none")`).Bool()
}

