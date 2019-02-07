package seed

import "bytes"
import "github.com/qlova/seed/style"

//A harvester collects the extracts the necessary information of the seed tree in order to render the application. 
type harvester struct {
	//Assets associated with a seed.
	assets []Asset

	//Fonts associated with a seed (need special handling compared to assets).
	fonts map[style.Font]struct{}
}

func newHarvester() *harvester {
	return &harvester{
		fonts: make(map[style.Font]struct{}),
	}
}

//Do the harvesting.
func (h *harvester) harvest(seed Seed) {
	
	//Add assets to the factory.
	if seed.assets != nil {
		h.assets = append(h.assets, seed.assets...)
	}
	
	//Add fonts to the factory.
	if seed.font.FontFace.FontFamily != "" {
		h.fonts[seed.font] = struct{}{}
	}
	
	//Recursively harvest children.
	for _, child := range seed.children {
		h.harvest(child.Root())
	}
}

//Harvest and combine the results with the application.
func (app *App) build() {
	app.harvester.harvest(app.Root())
	
	//Index assets for the application.
	for _, asset := range app.assets {
		app.Assets[asset.path] = true
	}
}

//Return rendered fonts.
func (h *harvester) Fonts() []byte {

	var buffer bytes.Buffer

	for font := range h.fonts {
		buffer.WriteString("@font-face {")
		buffer.Write(font.Bytes())
		buffer.WriteByte('}')
	}

	return buffer.Bytes()
}
