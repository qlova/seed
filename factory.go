package seed

type factory struct {
	assets []Asset
}

func (f *factory) build(seed Seed) {
	if seed.assets != nil {
		f.assets = append(f.assets, seed.assets...)
	}
	for _, child := range seed.children {
		f.build(child.Root())
	}
}

//Build the application, finding all the nested assets, handlers, etc.
func (app *App) build() {
	var f = new(factory)
		f.build(app.Root())
		
	//Register assets.
	for _, asset := range f.assets {
		app.Assets[asset.path] = true
	}
}
