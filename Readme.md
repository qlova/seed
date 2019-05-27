# ![logo](media/logo.png) Qlovaseed 
[![Godoc](https://godoc.org/github.com/qlova/seed)](https://godoc.org/github.com/qlova/seed)  

Qlovaseed is a cross-platform Go framework for building apps.

**Please Note:** Qlovaseed is currently very unstable and in heavy development.

An app written in Qlovaseed spins up a http server to host the app, 
the executable can be dropped on a public facing web server in order to serve it as a progressive webapp.
Alternatively, the executable can be distributed to users directly and run
 as a local application (users must have a compatiable web-browser such as Chrome).

## Installing

You can install Qlovaseed using go get.
```
	go get -u -v github.com/qlova/seed
```

## Getting started

Create a file called HelloWorld.go and paste in the following contents:

```
	package main

	import "github.com/qlova/seed"

	func main() {
		//Create a new app with Hello World as both the title and the content.
		seed.NewApp("Hello World", "Hello World").Launch()
	}
```

In the same folder, run go build to create an executable of the app that you can run to see the app in action!