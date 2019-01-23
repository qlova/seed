/*

This package can be used to write web applications without the need for HTML, CSS or Javascript.

	package main
	
	import "github.com/qlova/seed"
	
	func main() {
		var App = seed.NewApp()
		App.SetName("Hello World")
		App.SetText("Hello World")
		App.Launch()
	}

Applications can be local or remote and are automatically served and managed by the package.

*/
package seed