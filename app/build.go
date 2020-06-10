package app

import (
	"fmt"

	"github.com/qlova/seed"
	"github.com/qlova/seed/app/manifest"
	"github.com/qlova/seed/asset"
	"github.com/qlova/seed/css"
	"github.com/qlova/seed/html"
	"github.com/qlova/seed/html/attr"
	"github.com/qlova/seed/js"
	"github.com/qlova/seed/js/window"
	"github.com/qlova/seed/page"
	"github.com/qlova/seed/popup"
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
func (a App) build() {
	var app app
	a.Seed.Read(&app)

	//We need to check if onerror is defined.
	var Script script.Data
	app.document.Body.Read(&Script)

	app.document.Body.With(
		seed.If(Script.On["error"] == nil,
			script.OnError(func(q script.Ctx, err script.Error) {
				q(window.Alert(err.String))
			}),
		),

		column.New(seed.NewOption(func(c seed.Seed) {
			if app.loadingPage != nil {
				loading_page := app.loadingPage.Page(page.RouterOf(c))
				app.document.Body.With(loading_page)

				c.With(script.OnReady(func(q script.Ctx) {
					fmt.Fprintf(q, `seed.LoadingPage = q.get("%v"); seed.CurrentPage = seed.LoadingPage;`, html.ID(loading_page))
				}))
			}
		})),
	)

	app.document.Body.With(
		page.Harvest(app.page),
		popup.Harvest(),
	)

	var onready = string(script.Render(a))
	var scripts = js.Scripts(a)

	app.worker.Assets = asset.Of(a)
	a.Seed.Write(app)

	app.document.Head.With(
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
				c.With(link.New(
					attr.Set("rel", "shortcut icon"),
					attr.Set("type", "image/png"),
					attr.Set("href", icon.Source),
				))
			}

			c.With(link.New(
				attr.Set("rel", "apple-touch-icon"),
				attr.Set("sizes", icon.Sizes),
				attr.Set("href", icon.Source),
			))
		})),

		style.New(html.Set(CSS+string(css.Render(a)))),

		//Add external scripts.
		repeater.New(scripts, repeater.Do(func(c repeater.Seed) {
			c.With(script_html.New(
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

											console.log("found update");

											if (!window.localStorage.getItem("updated")) {
												window.localStorage.setItem("updated", "true");
												return;
											}
											if (window.localStorage.getItem("updating")) {
												return;
											}

											console.log("updating");
											
											if (document.body.onupdatefound) document.body.onupdatefound();
										}
								}
							};
						}
					}, function(err) {
						throw err;
					});
				}

				function launch(url, title, w, h) {
					const y = window.top.outerHeight / 2 + window.top.screenY - ( h / 2);
					const x = window.top.outerWidth / 2 + window.top.screenX - ( w / 2);
					return window.open(url, title, `+"`"+`toolbar=no, location=no, directories=no, status=no, menubar=no, scrollbars=no, resizable=no, copyhistory=no, width=${w}, height=${h}, top=${y}, left=${x}`+"`"+`);
				}

				document.addEventListener("contextmenu", function (e) {
					e.preventDefault();
				}, false);

				if (!seed.production) {
					let url = new URL('/seed.socket', location.href);
					url.protocol = url.protocol.replace('http', 'ws');
					let Socket = new WebSocket(url.href);

					Socket.onclose = function() {
						///close();
					}
					Socket.onerror = function() {
						//close();
					}
					Socket.onmessage = function(event) {
						eval(event.data);
					}

					Socket.onopen = function(event) {
						let old_log = console.log;
						console.log = function() {
							let args = [];
							for (let arg of arguments) {
								args.push(arg);
							}
							Socket.send(JSON.stringify(args));

							if (seed.debug) old_log.apply(null, arguments);
						};
						let old = console.error;
						console.error = function() {
							old.apply(null, arguments);
							console.log.apply(null, arguments);
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
