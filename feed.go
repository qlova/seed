package seed

import "strconv"

func (seed Seed) AddFeed(template Seed, feed func(Client)) {
	var WrapperSeed = New()
	WrapperSeed.SetSize(100, 100)

	minified, err := mini(template.HTML())
	if err != nil {
		//Panic?
	}
	WrapperSeed.OnReady(func(q Script) {
		//TODO do this properly.
		q.Get(WrapperSeed).Javascript(q.Get(WrapperSeed).Element()+`.innerHTML = `+strconv.Quote(string(minified))+`.repeat(4);`)
	})

	seed.Add(WrapperSeed)	
}

func (seed Seed) SetTextFeeder(name string) {
	
}