package seed

import "bytes"
import "github.com/qlova/seed/style"

import "net/http"
type dynamicHandler struct {
	id string
	handler func(User)
}

//A harvester collects the extracts the necessary information of the seed tree in order to render the application. 
type harvester struct {
	//Assets associated with a seed.
	assets []Asset

	//Fonts associated with a seed (need special handling compared to assets).
	fonts map[style.Font]struct{}
	
	//Animations and animation ids associated with a seed.
	animations []Animation
	animationNames []string
	
	//Dynamic handlers
	dynamicHandlers []dynamicHandler
}

func newHarvester() *harvester {
	return &harvester{
		fonts: make(map[style.Font]struct{}),
	}
}

//Do the harvesting.
func (h *harvester) harvest(seed Seed) {
	
	//Harvest Animations.
	if seed.animation != nil {
		h.animations = append(h.animations, seed.animation)
		h.animationNames = append(h.animationNames, seed.ID())
	}

	//Harvest Assets.
	if seed.assets != nil {
		h.assets = append(h.assets, seed.assets...)
	}
	
	//Harvest Dynamic Handlers.
	if seed.dynamicText != nil {
		h.dynamicHandlers = append(h.dynamicHandlers, dynamicHandler{
			id: seed.id,
			handler: seed.dynamicText,
		})
	}
	
	//Harvest Fonts.
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

//Return rendered animations.
func (h *harvester) Animations() []byte {
	var buffer bytes.Buffer

	for i, animation := range h.animations {
		buffer.WriteString("@keyframes "+h.animationNames[i]+" {")
		buffer.Write(animation.Bytes())
		buffer.WriteByte('}')
	}

	return buffer.Bytes()
}

func (h *harvester) DynamicHandler() (func(w http.ResponseWriter, r *http.Request)) {
	if len(h.dynamicHandlers) == 0 {
		return nil
	}

	return func(w http.ResponseWriter, r *http.Request) {
		for _, handler := range h.dynamicHandlers {
			w.Write([]byte(`"`))
			w.Write([]byte(handler.id))
			w.Write([]byte(`":"`))
			handler.handler(User{}.FromHandler(w, r))
			w.Write([]byte(`"`))
		}
	}
}
