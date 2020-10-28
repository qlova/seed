package app

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/NYTimes/gziphandler"

	"qlova.org/seed/web/api"
	"qlova.org/seed/assets/inbed"
	"qlova.org/seed/client"
	"qlova.org/seed/web/css"
	"qlova.org/seed/web/js"
)

var intranet, _ = regexp.Compile(`(^192\.168\.([0-9]|[0-9][0-9]|[0-2][0-5][0-5])\.([0-9]|[0-9][0-9]|[0-2][0-5][0-5]):.*$)`)

func isLocal(r *http.Request) (local bool) {
	if r.Header.Get("X-Real-IP") != "" || r.Header.Get("X-Forwarded-For") != "" {
		return false
	}

	local = strings.Contains(r.RemoteAddr, "[::1]") || strings.Contains(r.RemoteAddr, "127.0.0.1")
	if local {
		return
	}
	if intranet.Match([]byte(r.RemoteAddr)) {
		local = true
	}

	split := strings.Split(r.Host, ":")
	if len(split) == 0 {
		local = false
	} else {
		if split[0] != "localhost" {
			local = false
		}
	}

	return
}

//Handler returns an http.Handler that serve's the app.
func (a App) Handler() http.Handler {

	//TODO implement basic auth for staging endpoints.
	//use crypto/subtle

	var app app
	a.Load(&app)

	var AssetsServer = inbed.FileSystem{}

	inbed.File("assets")

	if err := inbed.Done(); err != nil {
		log.Println(err)
	}

	router := http.NewServeMux()

	a.build()

	var rendered = app.document.Render()

	var document, err = mini(rendered)
	if err != nil {
		document = rendered
	}

	var scripts = js.Scripts(app.document.Seed)
	var stylesheets = css.Stylesheets(app.document.Seed)
	var imports = js.Imports()

	//Checksum is used for versioning, ensure deterministic renderers are used to prevent distributed versions from mismatching.
	//use deterministic ordered-maps instead of default maps or sort the keys before iteration.
	var checksum = md5.Sum(document)

	var version = hex.EncodeToString(checksum[:])

	app.worker.Version = version

	var worker = app.worker.Render()

	router.Handle("/Qlovaseed.png", gziphandler.GzipHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")
		icon, _ := fsByte(false, "/Qlovaseed.png")
		w.Write(icon)
		return
	})))

	router.Handle("/assets/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		AssetsServer.ServeHTTP(w, r)
	}))

	router.Handle("/go/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if version, err := r.Cookie("version"); err == nil && version.Value != app.worker.Version {

			http.SetCookie(w, &http.Cookie{
				Name:   "version",
				Value:  "",
				Path:   "/",
				MaxAge: -1,
			})

			w.Write([]byte(`if (!seed.production) {
				throw "";
			} else {
				if (document.body.onupdatefound) await document.body.onupdatefound();
				throw "";
			}`))
			return
		}

		client.Handler(w, r, r.URL.Path[4:])
	}))

	router.Handle("/seed.socket", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isLocal(r) && a.port == ":0" {
			localClients++
			singleLocalConnection = localClients == 1
			socket(w, r)
		} else {
			localClients = 99
			socket(w, r)
		}
	}))

	var manifest = app.manifest.Render()
	router.Handle("/app.webmanifest", gziphandler.GzipHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		w.Write(manifest)
	})))

	router.Handle("/robots.txt", gziphandler.GzipHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("\n"))
	})))

	router.Handle("/index.js", gziphandler.GzipHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isLocal(r) {
			//Don't use a web worker if we are running locally.
			w.Header().Set("content-type", "text/javascript")
			w.Write([]byte(`self.addEventListener('install', () => {self.skipWaiting();});`))
		} else {
			w.Header().Set("content-type", "text/javascript")
			w.Write(worker)
		}
	})))

	for route, handler := range api.Routes(app.document.Seed) {
		router.Handle(route, handler)
	}

	router.Handle("/", gziphandler.GzipHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if content, ok := scripts[r.URL.Path]; ok {
			w.Header().Set("Content-Type", "text/javascript")
			fmt.Fprint(w, content)
			return
		}

		if content, ok := stylesheets[r.URL.Path]; ok {
			w.Header().Set("Content-Type", "text/css")
			fmt.Fprint(w, content)
			return
		}

		if content, ok := imports[r.URL.Path]; ok {
			w.Header().Set("Content-Type", "text/javascript")
			fmt.Fprint(w, content)
			return
		}

		if version, err := r.Cookie("version"); err != nil || version.Value == app.worker.Version {
			http.SetCookie(w, &http.Cookie{
				Name:    "version",
				Value:   app.worker.Version,
				Path:    "/",
				Expires: time.Now().Add(24 * time.Hour),
			})
		}

		w.Write(document)
	})))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//PWA's require HTTPS to work.
		if r.Header.Get("X-Forwarded-Proto") == "http" && r.Host != "localhost" {
			http.Redirect(w, r, "https://"+r.Host+r.URL.String(), http.StatusSeeOther)
		}

		router.ServeHTTP(w, r)
	})
}
