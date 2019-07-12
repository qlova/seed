# ![logo](media/logo.png) Qlovaseed [![Godoc](https://godoc.org/github.com/qlova/seed?status.svg)](https://godoc.org/github.com/qlova/seed) [![Go Report Card](https://goreportcard.com/badge/github.com/qlova/seed)](https://goreportcard.com/report/github.com/qlova/seed) [![Build Status](https://travis-ci.org/qlova/seed.svg?branch=master)](https://travis-ci.org/qlova/seed)

The cross-platform Go framework for building apps.

## Features

* Qlovaseed apps are optimised for low-data usage, fast loading times and smooth animations.
* No need to install, just visit the URL and use app straight away. Save it to the homescreen for offline use.
* Cross platform, develop, build and run these apps on almost any device. All with the same code.
* Develop the frontend and backend of an app simultanously, the only 'Full Stack' framework of its kind. 

![showcase](media/showcase.jpg)

## Installing

You can get Qlovaseed using go get.
```sh
go get -u -v github.com/qlova/seed
```

## Getting started

Create HelloWorld.go file and paste in the following contents:

```go
package main

import "github.com/qlova/seed"
import "github.com/qlova/seeds/text"
import "github.com/qlova/seeds/expander"

func main() {
	var App = seed.NewApp("Hello World")

	expander.AddTo(App)
	text.AddTo(App, "Hello World")
	expander.AddTo(App)

	App.Launch()
}

```

In the same folder, run 'go build' to create an executable for the app, run this to see the app in action!

## Widgets and Logic

Create a file called MyApp.go and paste in the following contents:

```go
package main

import "github.com/qlova/seed"

//Import a seed to use it, a list of seeds can be found [here](https://github.com/qlova/seeds).
import "github.com/qlova/seeds/button"

func main() {
	var App = seed.NewApp("My App")

	//In order to add a widget to your app, or container, use the package's AddTo method.
	var ClientPowered = button.AddTo(App, "My callback runs on the client")
	
		ClientPowered.OnClick(func(q seed.Script) {
			ClientPowered.Script(q).SetText(q.String("You clicked me!"))
		})
	
	
	var ServerPowered = button.AddTo(App, "My callback runs on the server")
	
		//You can style widgets with methods of the style package.
		ServerPowered.SetColor(seed.RGB(100, 100, 0))
	
		ServerPowered.OnClick(seed.Go(func(user seed.User) {
			ServerPowered.For(user).SetText("You clicked me!")
		}))

	App.Launch()
}
```

This example shows a quick glimpse on how powerful Qlovaseed can be. You can find more widgets in the [seeds repository](https://github.com/qlova/seeds).

## Themes

A work in progress material theme can be found [here](https://github.com/qlova/theme/tree/master/material).
Check the examples folder to learn how to use it.

**Please remember**, this framework is in development, it does not have a stable API and features are currently implemented as required.
