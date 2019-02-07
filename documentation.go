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

*/
package seed
