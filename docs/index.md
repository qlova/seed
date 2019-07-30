![Logo](logo.svg)  
  
Welcome to the Qlovaseed documentation!

Qlovaseed is an app-development framework for Go, this means you can use it to develop:

* Desktop applications
* Mobile applications.
* Progresive webapps. 

Let's get started. To use Qlovaseed you will need Go installed.
You can get Go [here](https://golang.org/dl/)

<iframe width="560" height="315" src="https://www.youtube.com/embed/16ZOmBmqJ0s" frameborder="0" allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>

Run this in a commandline to update/install Qlovaseed:

```sh
	go get -u -v github.com/qlova/seed
```

Now you can build a HelloWorld application.
In Qlovaseed, seeds are used that will grow into the application.
More information on seeds will be covered later in the documentation.

Create a folder called HelloWorld and create a file inside called HelloWorld.go, paste in the following:
```go
package main

import (
	//This is the Qlovaseed framework.
	"github.com/qlova/seed"

	//text is a type of seed that will display text.
	"github.com/qlova/seed/text"

	//expander is a type of seed that will expand to take up space.
	//it is very useful for centering other seeds.
	"github.com/qlova/seed/expander"
	
)

func main() {
	//Create a new app called Hello World, keep in mind that an app is also a seed.
	var App = seed.NewApp("Hello World")

	//seeds can be added to other seeds with the seed.AddTo(other) pattern.
	expander.AddTo(App)

	//text has an optional second argument that sets the text.
	text.AddTo(App, "Hello World")

	expander.AddTo(App)

	//Launch the app, this will open your app in Google Chrome or your default browser.
	App.Launch()
}

```

On the commandline, run this to build the app:

```sh
	go build
```

There will be a resuting executable, this is your app. Try running it!