package app

import (
	"fmt"

	"qlova.org/seed"
	"qlova.org/seed/app/manifest"
	"qlova.org/seed/asset/assets"
	"qlova.org/seed/client"
	"qlova.org/seed/client/clientside"
	"qlova.org/seed/css"
	"qlova.org/seed/html"
	"qlova.org/seed/html/attr"
	"qlova.org/seed/js"
	"qlova.org/seed/js/window"
	"qlova.org/seed/page"
	"qlova.org/seed/popup"
	"qlova.org/seed/script"

	"qlova.org/seed/s/column"
	"qlova.org/seed/s/html/link"
	"qlova.org/seed/s/html/meta"
	"qlova.org/seed/s/html/style"
	"qlova.org/seed/s/html/title"
	"qlova.org/seed/s/repeater"

	script_html "qlova.org/seed/s/html/script"
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

		client.OnLoad(clientside.Render(app.document.Body)),
	)

	app.document.Body.With(
		page.Harvest(app.page),
		popup.Harvest(),
	)

	var onready = string(script.Render(a))
	var scripts = js.Scripts(a)
	var stylesheets = css.Stylesheets(a)

	app.worker.Assets = assets.Of(a)
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

		repeater.New(stylesheets, repeater.Do(func(c repeater.Seed) {
			c.With(link.New(
				attr.Set("href", c.Data.Index().String()),
				attr.Set("rel", "stylesheet"),
			))
		})),

		//Add js.
		script_html.New(
			html.Set(`
			if ((window.matchMedia('(display-mode: standalone)').matches) || (window.navigator.standalone)) {
				window.name = "";
			}
			
			document.addEventListener('DOMContentLoaded', async function() { 
				`+onready+`

				
				window.addEventListener('resize', function() {
					q.setvar("app.standalone", "", (document.fullscreen || (window.matchMedia('(display-mode: standalone)').matches) || (window.navigator.standalone) || window.name == 'installed'));
				});

				let AddToHomeScreenEvent = null;
				window.addEventListener('beforeinstallprompt', (e) => {
					// Prevent Chrome 76 and later from showing the mini-infobar
					e.preventDefault();
					// Stash the event so it can be triggered later.
					AddToHomeScreenEvent = e;
					q.setvar("app.installable", "", true);
				});

				function install() {
					if (AddToHomeScreenEvent) {
						AddToHomeScreenEvent.prompt();
						return;
					}
					//Provide instructions.
					let instructions = document.createElement("div");
					instructions.id = "install_instructions"
					let text = document.createElement("span");
					text.style.padding = "2em";
					//IOS Safari.
					if (navigator.vendor && navigator.vendor.indexOf('Apple') > -1 &&
						navigator.userAgent &&
						navigator.userAgent.match(/iPhone|iPad|iPod/i) &&
						navigator.userAgent.indexOf('CriOS') == -1 &&
						navigator.userAgent.indexOf('FxiOS') == -1) {
						text.innerText = "Click on the share button\n swipe the bottom row to the left\n tap on 'Add to Home Screen'";
						let iosShareButton = document.createElement("div");
						iosShareButton.innerHTML = '<svg xmlns="http://www.w3.org/2000/svg" width="51.518" height="75.01" viewBox="0 0 51.518 75.01"><g id="noun_Share_1504278" transform="translate(-25 -12.2)"><path id="Path_2055" data-name="Path 2055" d="M52.307,14.055,50.452,12.2,48.6,14.055h-.206v.206L35.1,27.552l2.885,2.885L48.392,20.031V65.573h4.121V20.031L62.92,30.437,65.8,27.552,52.513,14.261v-.206Z" transform="translate(0.307)" fill="rgba(2,0,0,0.8)"/><path id="Path_2056" data-name="Path 2056" d="M25,86.427H76.518V38H62.093v4.121H72.4V82.306H29.121V42.121h10.3V38H25Z" transform="translate(0 0.783)" fill="rgba(2,0,0,0.8)"/></g></svg>'
						instructions.appendChild(iosShareButton);
						instructions.appendChild(text);
					} else if (navigator.userAgent.match(/iPhone|iPad|iPod/i)) {
						text.innerText = "Open this app in Safari to install it";
						instructions.appendChild(text);
					} else {       
						alert("Sorry, it looks like we cannot install this app to your device, you can continue using it in the browser.");
						return;
					}
					document.body.appendChild(instructions);
				}


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

				if (seed.goto) await seed.goto.ready();

				});`),
		),
	)
}
