package seed

import (
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/NYTimes/gziphandler"
	"github.com/qlova/seed/script"
)

//Handler returns a http handler that serves this application.
func (runtime Runtime) Handler() http.Handler {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	minified, err := mini(runtime.app.Render(Default))
	if err != nil {
		//Panic?
	}

	var html = runtime.app.render(false, Default)

	var worker = runtime.app.Worker.Render()
	var manifest = runtime.app.Manifest.Render()

	var custom = runtime.app.CustomHandler()

	var LocalClients = 0

	intranet, err := regexp.Compile(`(^192\.168\.([0-9]|[0-9][0-9]|[0-2][0-5][0-5])\.([0-9]|[0-9][0-9]|[0-2][0-5][0-5]):.*$)`)
	if err != nil {
		panic("invalid regexp!")
	}

	return gziphandler.GzipHandler(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
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

		var local bool

		if !Production {
			local = strings.Contains(request.RemoteAddr, "[::1]") || strings.Contains(request.RemoteAddr, "127.0.0.1")
			if intranet.Match([]byte(request.RemoteAddr)) {
				local = true
			}

			//Editmode socket.
			if request.URL.Path == "/socket" && local {
				LocalClients++
				singleLocalConnection = LocalClients == 1
				socket(response, request)
				return
			}
		}

		if request.URL.Path == "/Qlovaseed.png" {
			response.Header().Set("Content-Type", "image/png")
			icon, _ := fsByte(false, "/Qlovaseed.png")
			response.Write(icon)
			return
		}

		//Is this an embedded resource? Imported libraries will add these.
		if embedded(response, request) {
			return
		}

		//Remote procedure calls.
		if len(request.URL.Path) > 5 && request.URL.Path[:6] == "/call/" {
			script.Handler(response, request, request.URL.Path[6:])
			return
		}

		//Remote procedure calls.
		if len(request.URL.Path) > 6 && request.URL.Path[:7] == "/feeds/" {
			feedHandler(response, request, request.URL.Path[7:])
			return
		}

		//Run custom handlers.
		if request.URL.Path != "/" {
			if custom != nil {
				custom(response, request)
			}

			if path.Ext(request.URL.Path) == "" {
				return
			}
		}

		if runtime.app.rest != "" && (request.URL.Host == runtime.app.rest || request.Host == runtime.app.rest) {
			response.Write([]byte(string("This place is for computers")))
			return

		}

		if request.URL.Path == "/.well-known/assetlinks.json" && runtime.app.pkg != "" {
			response.Header().Set("Content-Type", "application/json")
			response.Write([]byte(`[{
  "relation": ["delegate_permission/common.handle_all_urls"],
  "target" : { "namespace": "android_app", "package_name": "` + runtime.app.pkg + `",
               "sha256_cert_fingerprints": [`))

			for i, hash := range runtime.app.hashes {
				response.Write([]byte("\"" + hash + "\""))
				if i < len(runtime.app.hashes)-1 {
					response.Write([]byte(`,`))
				}
			}

			response.Write([]byte(`] }
}]`))
			return
		}

		//Serve service worker.
		if request.URL.Path == "/index.js" {
			if local {
				//Don't use a web worker if we are running locally.
				response.Header().Set("content-type", "text/javascript")
				response.Write([]byte(`self.addEventListener('install', () => {self.skipWaiting();});`))
			} else {
				response.Header().Set("content-type", "text/javascript")
				response.Write(worker)
			}
			return
		}

		//Serve web manifest.
		if request.URL.Path == "/app.webmanifest" {
			response.Header().Set("content-type", "application/json")
			response.Write(manifest)
			return
		}

		//Serve assets.
		if path.Ext(request.URL.Path) != "" {
			http.ServeFile(response, request, dir+"/assets"+request.URL.Path)
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

		//Anything else? Serve application.
		if local {
			response.Write(html)
		} else {
			response.Write(minified)
		}
	}))
}
