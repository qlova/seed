/*

This package can be used to write web applications without the need for HTML, CSS or Javascript.

	package main
	
	import "github.com/qlova/seed"
	
	func main() {
		seed.NewApp("Hello World", "Hello World").Launch()
	}

Applications are automatically served and managed by the package.

Deploying (TODO)

Before you deploy, you will need a server and a domain name pointing at your server.

In order to use an application on a mobile device. Your seed needs to be deployed. In order to deploy you need to run SetHost and SetEmail on your application.
Then you can pass the deploy argument to your app. Make sure you have ssh installed on both your client and server.

	./App -deploy

  Android and IOS apps
  
You can export an .apk or .ipa of your application. These will need to be signed before being uploaded to an appstore.

	./App -apk
	./App -ipa


Environmentally Friendly

Qlovaseed uses less CPU usage on client devices compared to Javascript/CSS frameworks.
Since applications are cached on clients, they do not need to communicate with your server as much, meaning applications with Qlovaseed cater for more returning users than standard web technologies.
You will save money by deploying your application with Qlovaseed.

Workflow

Before you write a line of code, sit down with a piece of paper and create some drawings of how you would like your application to look, feel & function. Think about different platforms (eg. mobile, tablet, desktop) and portait/landscape orientations. This is the design process and is the most important part of creating an app.
A rough design is fine. Now for each design, identify the outline of the application, recognise the types of widgets, the rows, columns and spacing. You can do this in your head if you like.
Now you can actually program the application!
Start by picking a page of the application and importing the widgets you need, you can currently select from:

	* button
	* column
	* datebox
	* document
	* expander
	* filepicker
	* header
	* image
	* line
	* link
	* listbox
	* passwordbox
	* popup
	* row
	* spacer
	* text
	* textarea
	* textbox
	* toolbar
	* video

To use one of these widgets, import it at the top of your .go file.

	package main
	
	import "github.com/qlova/seed"

	import (
		"github.com/qlova/seed/widgets/row"
		"github.com/qlova/seed/widgets/textbox"
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
