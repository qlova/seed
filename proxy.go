package seed

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"
)

var LocalClients = 0

func proxy(from, to string) {
	
	url, err := url.Parse("http://localhost"+from)
	if err != nil {
		panic("localhost"+from+" is an invalid url!")
	}
	
	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.ErrorHandler = func(response http.ResponseWriter, request *http.Request, err error) {
		time.Sleep(time.Second)
		
		var local = strings.Contains(request.RemoteAddr, "[::1]")
		//Editmode socket.
		if request.URL.Path == "/socket" && local {
			RELOADING = false
			
			LocalClients++
			println(LocalClients)
			singleLocalConnection = LocalClients == 1
			socket(response, request)
			return
		}
		
		proxy.ServeHTTP(response, request)
	}
	
	
	
	http.Handle("/", http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		
		
		var local = strings.Contains(request.RemoteAddr, "[::1]")
		//Editmode socket.
		if request.URL.Path == "/socket" && local {
			LocalClients++
			singleLocalConnection = LocalClients == 1
			socket(response, request)
			return
		}
		
		proxy.ServeHTTP(response, request)
	}))
	

	err = http.ListenAndServe(to, nil)
	if err != nil {
		println(err.Error())
	}
}
