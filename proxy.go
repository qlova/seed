package seed

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
	"strings"
	"time"
)

var localClients = 0

func proxy(from, to string) {

	url, err := url.Parse("http://localhost" + from)
	if err != nil {
		panic("localhost" + from + " is an invalid url!")
	}

	intranet, err := regexp.Compile(`(^192\.168\.([0-9]|[0-9][0-9]|[0-2][0-5][0-5])\.([0-9]|[0-9][0-9]|[0-2][0-5][0-5]):.*$)`)
	if err != nil {
		panic("invalid regexp!")
	}

	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.ErrorHandler = func(response http.ResponseWriter, request *http.Request, err error) {
		time.Sleep(time.Second)

		var local = strings.Contains(request.RemoteAddr, "[::1]")

		if intranet.Match([]byte(request.RemoteAddr)) {
			local = true
		}

		//Editmode socket.
		if request.URL.Path == "/socket" && local {
			reloading = false

			localClients++
			println(localClients)
			singleLocalConnection = localClients == 1
			socket(response, request)
			return
		}

		proxy.ServeHTTP(response, request)
	}

	http.Handle("/", http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {

		var local = strings.Contains(request.RemoteAddr, "[::1]")
		//Editmode socket.
		if request.URL.Path == "/socket" && local {
			localClients++
			singleLocalConnection = localClients == 1
			socket(response, request)
			return
		}

		proxy.ServeHTTP(response, request)
	}))

	//TLS support?

	err = http.ListenAndServe(to, nil)
	if err != nil {
		println(err.Error())
	}
}
