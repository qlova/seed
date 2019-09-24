/*
Package seed can be used to write web applications without the need for HTML, CSS or Javascript.

	package main

	import "github.com/qlova/seed"

	func main() {
		seed.NewApp("Hello World", "Hello World").Launch()
	}

Applications are automatically served and managed by the package.

Getting Started

Before you write a line of code, sit down with a piece of paper and create some drawings of how you would like your application to look, feel & function. Think about different platforms (eg. mobile, tablet, desktop) and portait/landscape orientations. This is the design process and is the most important part of creating an app.
A rough design is fine. Now for each design, identify the outline of the application, recognise the types of widgets, the rows, columns and spacing. You can do this in your head if you like.
Now you can actually program the application!
Start by picking a page of the application and importing the widgets you need (a full list can be found at https://github.com/qlova/seeds).

To use one of these widgets, import it at the top of your .go file.

	package main

	import "github.com/qlova/seed"

	import (
		"github.com/qlova/seeds/row"
		"github.com/qlova/seeds/textbox"
	)

Now you can add widgets to your application and to other widgets like this.

	func main() {
		var App = seed.NewApp()

		var Row = row.AddTo(App)

		textbox.AddTo(Row)

		App.Launch()
	}

All widgets have a AddTo(parent) method for clearly creating and adding the widget to a parent in a single line.
They also have a New() method that returns the widget for more advanced purposes.

*/
package seed
