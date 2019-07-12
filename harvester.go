package seed

import "bytes"
import "github.com/qlova/seed/style"
import "github.com/qlova/seed/style/css"
import "github.com/qlova/seed/script"
import "net/http"

//A harvester collects the extracts the necessary information of the seed tree in order to render the application.
type harvester struct {
	//Assets associated with a seed.
	assets []Asset

	//Fonts associated with a seed (need special handling compared to assets).
	fonts map[style.Font]struct{}

	//Animations and animation ids associated with a seed.
	animations     []Animation
	animationNames []string

	//Dynamic handlers
	dynamicHandlers map[string][]func(Script)

	customHandlers []func(response http.ResponseWriter, request *http.Request)

	stateHandlers map[State][]func(Script)

	screenSmallerThans map[Unit]style.Sheet

	onReadyHandlers []func(Script)

	dependencies script.Dependencies
}

func newHarvester() *harvester {
	return &harvester{
		fonts:              make(map[style.Font]struct{}),
		stateHandlers:      make(map[State][]func(Script)),
		screenSmallerThans: make(map[Unit]style.Sheet),
		dynamicHandlers: make(map[string][]func(Script)),

		dependencies: make(script.Dependencies),
	}
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

	//Harvest Dynamic Handlers.
	if seed.dynamicText.Variable != "" {
		var index = string(seed.dynamicText.Variable)
		h.dynamicHandlers[index] = append(h.dynamicHandlers[index], func(q Script) {
			seed.Script(q).SetText(seed.dynamicText.Script(q))
		})
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
	if seed.font.FontFace.FontFamily != "" {
		h.fonts[seed.font] = struct{}{}
	}

	//Harvest mediaQueries.
	if seed.screenSmallerThan != nil {
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
	}

	//Harvest onready handlers
	if seed.onready != nil && !seed.template {
		seed.ready = true
		h.onReadyHandlers = append(h.onReadyHandlers, seed.onready)
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

	buffer.WriteString(`document.addEventListener('DOMContentLoaded', function() { goto_ready = true;`)

	for _, handler := range h.onReadyHandlers {
		buffer.WriteByte('{')
		buffer.Write(script.ToJavascript(handler, h.dependencies))
		buffer.WriteByte('}')
	}

	//TODO move this elsewhere.
	buffer.WriteString(`
		if (window.localStorage) {
			if (!window.goto) return;
			let current_page = window.localStorage.getItem('*CurrentPage');
			if (current_page ) {
				goto(current_page);

				//clear history
				last_page = null;
				goto_history = [];

				if (get(current_page) && get(current_page).enterpage)
					get(current_page).enterpage();
			}
		}
	`)

	buffer.WriteString(`}, false);`)

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
			buffer.Write([]byte(script.ToJavascript(handler, h.dependencies)))
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
		if state.not {
			buffer.WriteString("function " + string(state.Variable) + "_unset() {")
			buffer.Write([]byte(script.ToJavascript(func(q Script) {
				q.Set(state, q.Bool(false))
			})))
		} else {
			buffer.WriteString("function " + string(state.Variable) + "_set() {")
			buffer.Write([]byte(script.ToJavascript(func(q Script) {
				q.Set(state, q.Bool(true))
			})))
		}

		for _, handler := range handlers {
			buffer.Write([]byte(script.ToJavascript(handler, h.dependencies)))
		}
		buffer.WriteByte('}')

		buffer.Write([]byte(script.ToJavascript(func(q Script) {
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
