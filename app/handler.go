package app

import (
	"net/http"

	"github.com/NYTimes/gziphandler"
	"github.com/qlova/seed/script"
)

//Handler returns an http.Handler that serve's the app.
func (app App) Handler() http.Handler {
	router := http.NewServeMux()

	app.build()

	var document = app.document.Render()

	router.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(document)
	}))

	router.Handle("/call/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		script.Handler(w, r, r.URL.Path[6:])
	}))

	var manifest = app.manifest.Render()
	router.Handle("/app.webmanifest", gziphandler.GzipHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		w.Write(manifest)
	})))

	return router
}
