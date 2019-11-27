package seed

import (
	"bytes"
	"net/http"

	"github.com/qlova/seed/internal"
	"github.com/qlova/seed/script"
	"github.com/qlova/seed/style"
	"github.com/qlova/seed/style/css"
)

//A harvester collects the extracts the necessary information of the seed tree in order to render the application.
type harvester struct {
	app *App

	Context internal.Context

	//Assets associated with a seed.
	assets []Asset

	//Fonts associated with a seed (need special handling compared to assets).
	//font-family CSS will be generated from this.
	fonts map[style.Font]struct{}

	//Animations and animation ids associated with a seed.
	animations     []Animation
	animationNames []string

	//Dynamic handlers
	dynamicHandlers map[string][]func(script.Ctx)

	customHandlers []func(response http.ResponseWriter, request *http.Request)

	stateHandlers map[State][]func(script.Ctx)

	screenSmallerThans map[Unit]style.Sheet

	onReadyHandlers [][]byte
}

func newHarvester() *harvester {
	return &harvester{
		fonts:              make(map[style.Font]struct{}),
		stateHandlers:      make(map[State][]func(script.Ctx)),
		screenSmallerThans: make(map[Unit]style.Sheet),
		dynamicHandlers:    make(map[string][]func(script.Ctx)),

		Context: internal.NewContext(),
	}
}

func (app *harvester) harvestOnReadyPage(seed Seed) []byte {
	var buffer bytes.Buffer

	buffer.WriteString("onready['")
	buffer.WriteString(seed.id)
	buffer.WriteString("'] = function() {")

	seed.page = false
	buffer.Write(app.harvestOnReady(seed))
	seed.page = true

	buffer.WriteString("};")

	return buffer.Bytes()
}

func (app *harvester) harvestOnReady(seed Seed) []byte {
	if seed.page && !seed.splash {
		return nil
	}

	var h = app

	var buffer bytes.Buffer

	//TODO sort?
	for event, handler := range seed.on {
		buffer.Write(script.ToJavascript(func(q script.Ctx) {
			q.Javascript(seed.Ctx(q).Element() + ".on" + event + " = function() {")
			handler(q)
			q.Javascript("};")
		}, h.Context))
	}

	//Harvest onready handlers
	if (seed.onready != nil) && !seed.Template {
		buffer.WriteString("get('")
		buffer.WriteString(seed.id)
		buffer.WriteString("').onready = function() {")
		buffer.Write(script.ToJavascript(seed.onready, h.Context))
		buffer.WriteString("};")
	}

	for _, child := range seed.children {
		buffer.Write(h.harvestOnReady(child.Root()))
	}

	if (seed.onready != nil) && !seed.Template {
		buffer.WriteString("get('")
		buffer.WriteString(seed.id)
		buffer.WriteString("').onready();")
	}

	return buffer.Bytes()
}

//Do the harvesting.
func (app *harvester) harvest(seed Seed) {

	var h = app

	//Harvest Animations.
	if seed.animation != nil {
		h.animations = append(h.animations, seed.animation)
		h.animationNames = append(h.animationNames, seed.id)
	}

	//Harvest Assets.
	if seed.assets != nil {
		h.assets = append(h.assets, seed.assets...)
	}

	for reference, handlers := range seed.dynamic.Handlers {
		h.dynamicHandlers[reference] = append(h.dynamicHandlers[reference], handlers...)
	}

	//Harvest State Handlers.
	if seed.states != nil {
		for state, handler := range seed.states {
			h.stateHandlers[state] = append(h.stateHandlers[state], handler)
		}
	}

	//Harvest Dynamic Handlers.
	if seed.handlers != nil {
		h.customHandlers = append(h.customHandlers, seed.handlers...)
	}

	//Harvest Fonts.
	if seed.font != "" {
		var path = string(seed.font)
		if font, ok := app.Context.FontCache[path]; ok {
			seed.Style.SetFont(font)
		} else {

			h.assets = append(h.assets, NewAsset(path))

			var font = style.NewFont(path)
			app.Context.FontCache[path] = font

			h.fonts[font] = struct{}{}

			seed.Style.SetFont(font)
		}
	}

	//Harvest mediaQueries.
	/*if seed.screenSmallerThan != nil {
		for unit, stile := range seed.screenSmallerThan {
			if h.screenSmallerThans[unit] == nil {
				h.screenSmallerThans[unit] = make(style.Sheet)
			}
			if seed.template {
				h.screenSmallerThans[unit].Add("."+seed.id, stile)
			} else {
				h.screenSmallerThans[unit].Add("#"+seed.id, stile)
			}
		}
	}*/

	//Recursively harvest children.
	for _, child := range seed.children {
		h.harvest(child.Root())
	}
}

//Harvest and combine the results with the application.
func (app *App) build() {

	var done = make(map[string]bool)

	//Recursively harvest children.
	for _, child := range app.Root().children {
		app.harvester.harvestOnReady(child.Root())
	}

	for {
		var pages = app.Context.Pages
		app.Context.ClearPages()

		if len(pages) == 0 {
			break
		}

		for id, page := range pages {
			if done[id] {
				continue
			}
			page := page.(Page)
			app.Add(page.Root())
			app.harvester.harvestOnReadyPage(page.Root())
			done[id] = true
		}
	}

	app.harvester.harvest(app.Root())
	app.harvester.app = app

	//Index assets for the application.
	for _, asset := range app.assets {
		app.Assets[asset.path] = true
	}

	app.built = true
}

//Return rendered fonts.
func (app *harvester) Fonts() []byte {
	var h = app

	var buffer bytes.Buffer

	for font := range h.fonts {
		buffer.WriteString("@font-face {")
		buffer.Write(font.Bytes())
		buffer.WriteByte('}')
	}

	return buffer.Bytes()
}

func (app *harvester) OnReadyHandler() []byte {
	var h = app
	var buffer bytes.Buffer

	buffer.Write(h.harvestOnReady(app.app.Root()))

	return buffer.Bytes()
}

//Return rendered animations.
func (app *harvester) Animations() []byte {
	var h = app

	var buffer bytes.Buffer

	for i, animation := range h.animations {
		buffer.WriteString("@keyframes " + h.animationNames[i] + " {")
		buffer.Write(animation.Bytes())
		buffer.WriteByte('}')
	}

	return buffer.Bytes()
}

func (app *harvester) DynamicHandlers() []byte {
	var h = app

	var buffer bytes.Buffer

	buffer.WriteString("var dynamic = {")
	for dynamic, handlers := range h.dynamicHandlers {
		buffer.WriteString("\"" + dynamic + "\": function() {")
		for _, handler := range handlers {
			buffer.Write([]byte(script.ToJavascript(handler, h.Context)))
		}
		buffer.WriteString("},")
	}
	buffer.WriteString("};")
	for dynamic := range h.dynamicHandlers {
		buffer.WriteString("dynamic[\"" + dynamic + "\"]();")
	}

	return buffer.Bytes()
}

func (app *harvester) MediaQueries() []byte {
	var h = app

	var buffer bytes.Buffer

	for unit, sheet := range h.screenSmallerThans {
		buffer.WriteString("@media (max-width: " + string(css.Decode(unit)) + ") {")
		buffer.Write(sheet.Bytes())
		buffer.WriteString("}")
	}

	return buffer.Bytes()
}

func (app *harvester) StateHandlers() []byte {
	var h = app

	var buffer bytes.Buffer

	for state, handlers := range h.stateHandlers {
		var reference = state.Bool.Ref()
		if state.not {
			buffer.WriteString("window." + reference + "_unset = function() {")
			buffer.Write([]byte(script.ToJavascript(func(q script.Ctx) {
				state.Bool.Set(q, q.Bool(false))
			})))
		} else {
			buffer.WriteString("window." + reference + "_set = function () {")
			buffer.Write([]byte(script.ToJavascript(func(q script.Ctx) {
				state.Bool.Set(q, q.Bool(true))
			})))
		}

		for _, handler := range handlers {
			buffer.Write([]byte(script.ToJavascript(handler, h.Context)))
		}
		buffer.WriteByte('}')
		buffer.WriteByte(';')

	}

	for state := range h.stateHandlers {
		buffer.Write([]byte(script.ToJavascript(func(q script.Ctx) {
			q.If(state.Get(q), func() {
				state.Set(q)
			})
		})))
	}

	return buffer.Bytes()
}

func (app *harvester) CustomHandler() func(w http.ResponseWriter, r *http.Request) {
	var h = app

	if len(h.customHandlers) == 0 {
		return nil
	}

	return func(w http.ResponseWriter, r *http.Request) {
		for _, handler := range h.customHandlers {
			handler(w, r)
		}
	}
}
