package main

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/script"
	"github.com/qlova/seeds/maps"
)

func main() {
	var App = seed.NewApp("GeoLocation")

	var Maps = maps.AddTo(App, maps.Options{
		Center: maps.Location{-36.8485, 174.7633}, //Default to Auckland, NZ
		Zoom:   15,
	})

	App.OnReady(func(q script.Ctx) {
		var Maps = Maps.Ctx(q)

		q.RequestGeoLocation().Then(func() {
			Maps.FlyTo(q.GeoLocation())
		}).Catch(func(err script.Error) {
			q.Alert(q.String("Could not determine location"))
		})
	})

	App.Launch()
}
