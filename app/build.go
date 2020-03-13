package app

import (
	"fmt"

	"github.com/qlova/seed"
	"github.com/qlova/seed/app/manifest"
	"github.com/qlova/seed/css"
	"github.com/qlova/seed/html"
	"github.com/qlova/seed/html/attr"
	"github.com/qlova/seed/page"
	"github.com/qlova/seed/script"

	"github.com/qlova/seed/s/column"
	"github.com/qlova/seed/s/html/link"
	"github.com/qlova/seed/s/html/meta"
	"github.com/qlova/seed/s/html/style"
	"github.com/qlova/seed/s/html/title"
	"github.com/qlova/seed/s/repeater"

	script_html "github.com/qlova/seed/s/html/script"
)

//Build builds the app.
func (app App) build() {
	app.document.Body.Add(
		column.New(seed.Do(func(c seed.Seed) {
			app.loadingPage.Page(page.Scope{Seed: c})

			c.Add(script.OnReady(func(q script.Ctx) {
				fmt.Fprintf(q, `seed.LoadingPage = seed.get("%v"); seed.CurrentPage = seed.LoadingPage;`, html.ID(c))
			}))
		})),
	)

	app.document.Body.Add(
		page.Harvest(app.page),
	)

	var onready = string(script.Render(app))
	var scripts = script.Scripts(app)

	app.document.Head.Add(
		meta.Charset("utf-8"),

		meta.New(meta.Viewport{
			Width: meta.DeviceWidth,

			InitialScale: 1,
			MinimumScale: 1,
			MaximumScale: 5,
		}),

		seed.If(app.description != "",
			meta.Description(app.description),
		),

		//Meta values.
		meta.Key("mobile-web-app-capable", "yes"),
		meta.Key("apple-mobile-web-app-capable", "yes"),
		meta.Key("apple-mobile-web-app-status-bar-style", "black"),
		meta.Key("apple-mobile-web-app-title", app.name),
		meta.Key("twitter:card", "app"),

		meta.Key(`theme-color`, app.manifest.ThemeColor),

		title.New(app.name),

		link.Manifest("/app.webmanifest"),

		//Add icons to app.
		repeater.New(app.manifest.Icons, repeater.Do(func(c repeater.Seed) {
			var icon = c.Data.Interface().(manifest.Icon)

			//The first icon can be the Favicon. TODO better heuristic? allow other file types.
			if c.Data.Index().Int() == 0 {
				c.Add(link.New(
					attr.Set("rel", "shortcut icon"),
					attr.Set("type", "image/png"),
					attr.Set("href", icon.Source),
				))
			}

			c.Add(link.New(
				attr.Set("rel", "apple-touch-icon"),
				attr.Set("sizes", icon.Sizes),
				attr.Set("href", icon.Source),
			))
		})),

		style.New(html.Set(CSS+string(css.Render(app)))),

		//Add external scripts.
		repeater.New(scripts, repeater.Do(func(c repeater.Seed) {
			c.Add(script_html.New(
				attr.Set("src", c.Data.Index().String()),
				attr.Set("defer", ""),
			))
		})),

		//Add js.
		script_html.New(
			html.Set(`document.addEventListener('DOMContentLoaded', async function() { 

				`+onready+`


				if ('serviceWorker' in navigator) {
					navigator.serviceWorker.register('/index.js').then(function(registration) {
						ServiceWorker_Registration = registration;
						registration.onupdatefound = function() {

							registration.installing.onstatechange = function(event) {
								switch (event.target.state) {
									case 'installed':
										if (navigator.serviceWorker.controller) {

											if (!window.localStorage.getItem("updated")) {
												window.localStorage.setItem("updated", "true");
												return;
											}
											if (window.localStorage.getItem("updating")) {
												return;
											}

											console.log("updating");
											
											//Clear all unnamed variables because they could have changed!
											//Unamed variables have a 'g_' prefix.
											for (let i in localStorage) {
												let item = window.localStorage[i];
												if (item.substring && item.substring(0, 3) == "g_") {
													window.localStorage.removeItem(i);
												}
											}
										}
								}
							};
						}
					}, function(err) {
						
					});
				}
				
				

				document.addEventListener("contextmenu", function (e) {
					e.preventDefault();
				}, false);

				if (!seed.production) {
					let url = new URL('/seed.socket', location.href);
					url.protocol = url.protocol.replace('http', 'ws');
					let Socket = new WebSocket(url.href);

					Socket.onclose = function() {
						close();
					}
					Socket.onerror = function() {
						close();
					}
					Socket.onmessage = function(event) {
						eval(event.data);
					}

					Socket.onopen = function(event) {
						console.log = function() {
							let args = [];
							for (let arg of arguments) {
								args.push(arg);
							}
							Socket.send(JSON.stringify(args));
						};
					}

					//Disable refresh on chrome because otherwise the app will close.
					document.onkeydown = function() {    
						switch (event.keyCode) { 
							case 116 : //F5 button
								event.returnValue = false;
								event.keyCode = 0;
								return false; 
							case 82 : //R button
								if (event.ctrlKey) { 
									event.returnValue = false; 
									event.keyCode = 0;  
									return false; 
								} 
						}
					}
				}
				});`),
		),
	)
}
