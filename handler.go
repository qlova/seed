package seed

import (
	"fmt"
	"net/http"
	"path"
	"regexp"
	"strings"

	"github.com/NYTimes/gziphandler"
	"github.com/qlova/seed/inbed"
	"github.com/qlova/seed/script"
)

var intranet, _ = regexp.Compile(`(^192\.168\.([0-9]|[0-9][0-9]|[0-2][0-5][0-5])\.([0-9]|[0-9][0-9]|[0-2][0-5][0-5]):.*$)`)

func isLocal(r *http.Request) (local bool) {
	if !Production {
		local = strings.Contains(r.RemoteAddr, "[::1]") || strings.Contains(r.RemoteAddr, "127.0.0.1")
		if intranet.Match([]byte(r.RemoteAddr)) {
			local = true
		}
	}
	return
}

//Handler returns a http handler that serves this application.
func (runtime Runtime) Handler() http.Handler {

	var AssetsServer = inbed.FileSystem{
		Prefix: "assets",
	}

	minified, err := mini(runtime.app.Render(Default))
	if err != nil {
		//Panic?
	}

	var html = runtime.app.render(false, Default)

	var worker = runtime.app.Worker.Render()
	var manifest = runtime.app.Manifest.Render()

	var LocalClients = 0

	var router = http.NewServeMux()

	for pattern, handler := range runtime.app.Handlers {
		router.Handle(pattern, handler)
	}

	router.Handle("/Qlovaseed.png", gziphandler.GzipHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")
		icon, _ := fsByte(false, "/Qlovaseed.png")
		w.Write(icon)
		return
	})))

	router.Handle("/call/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		script.Handler(w, r, r.URL.Path[6:])
	}))

	router.Handle("/conn/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		script.ConnectionHandler(w, r, r.URL.Path[6:])
	}))

	router.Handle("/.well-known/assetlinks.json", gziphandler.GzipHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if runtime.app.pkg != "" {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`[{
	"relation": ["delegate_permission/common.handle_all_urls"],
	"target" : { "namespace": "android_app", "package_name": "` + runtime.app.pkg + `",
				"sha256_cert_fingerprints": [`))

			for i, hash := range runtime.app.hashes {
				w.Write([]byte("\"" + hash + "\""))
				if i < len(runtime.app.hashes)-1 {
					w.Write([]byte(`,`))
				}
			}

			w.Write([]byte(`] }
	}]`))
		}
	})))

	//Socket for qlovaseed app-development features.
	router.Handle("/socket", gziphandler.GzipHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isLocal(r) {
			LocalClients++
			singleLocalConnection = LocalClients == 1
			socket(w, r)
		}
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

	router.Handle("/browserconfig.xml", gziphandler.GzipHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "text.xml")

		var icon string
		if len(runtime.app.Manifest.Icons) > 0 {
			icon = runtime.app.Manifest.Icons[0].Source
		}

		fmt.Fprintf(w, `<?xml version="1.0" encoding="utf-8"?>
<browserconfig>
	<msapplication>
		<tile>
			<square70x70logo src="%[1]v"/>
			<square150x150logo src="%[1]v"/>
			<wide310x150logo src="%[1]v"/>
			<square310x310logo src="%[1]v"/>
			<TileColor>%[2]v</TileColor>
		</tile>
	</msapplication>
</browserconfig>
`,
			icon, runtime.app.Manifest.ThemeColor)
	})))

	router.Handle("/app.webmanifest", gziphandler.GzipHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		w.Write(manifest)
	})))

	router.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//Serve assets.
		if path.Ext(r.URL.Path) != "" {
			AssetsServer.ServeHTTP(w, r)
			return
		}

		if isLocal(r) {
			w.Write(html)
		} else {
			w.Write(minified)
		}
	}))

	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		if origin := request.Header.Get("Origin"); origin == "https://"+runtime.app.host && origin != "" {
			response.Header().Set("Access-Control-Allow-Origin", origin)
		} else {
			response.Header().Set("Access-Control-Allow-Origin", "file://")
		}
		response.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		response.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if request.Method == "OPTIONS" {
			return
		}

		//Is this an embedded resource? Imported libraries will add these.
		if embedded(response, request) {
			return
		}

		if runtime.app.rest != "" && (request.URL.Host == runtime.app.rest || request.Host == runtime.app.rest) {
			response.Write([]byte(string("This place is for computers")))
			return
		}

		//Identify platform.
		/*if os.Getenv("IGNORE_PLATFORM") == "" {
			device := ua.Parse(request.UserAgent())

			if device.DeviceType == ua.DeviceComputer {
				response.Write(desktop)
				return
			}
		}*/

		router.ServeHTTP(response, request)
	})
}
