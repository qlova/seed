package maps

import "github.com/qlova/seed"
import "github.com/qlova/seed/script"

func init() {
	seed.Embed("/leaflet.js", []byte(Javascript))
	seed.Embed("/leaflet.css", []byte(Stylesheet))
}

//Returns a full-featured map.
func New() seed.Seed {	
	var m = seed.New()

	m.Require("leaflet.js")
	m.Require("leaflet.css")

	m.Stylable.Set("width", "100vw")
	m.Stylable.Set("height", "100vh")
	
	m.OnReady(func(q seed.Script) {
		q.Javascript(`let map = L.map("`+m.ID()+`", {center: [-36.2647, 174.7975], zoom: 13 }); L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', { maxZoom: 19, attribution: '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>' }).addTo(map); get("`+m.ID()+`").map = map;`)
	})
	
	return m
}

type Map script.Seed
