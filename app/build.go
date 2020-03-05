package app

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/css"
	"github.com/qlova/seed/html"
	"github.com/qlova/seed/script"

	"github.com/qlova/seed/s/html/link"
	"github.com/qlova/seed/s/html/meta"
	"github.com/qlova/seed/s/html/style"
	"github.com/qlova/seed/s/html/title"

	script_html "github.com/qlova/seed/s/html/script"
)

//Build builds the app.
func (app App) build() {
	var scripts = string(script.Render(app))

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

		meta.Key("mobile-web-app-capable", "yes"),
		meta.Key("apple-mobile-web-app-capable", "yes"),
		meta.Key("apple-mobile-web-app-status-bar-style", "black"),
		meta.Key("apple-mobile-web-app-title", app.name),
		meta.Key("twitter:card", "app"),

		title.New(app.name),

		link.Manifest("/app.webmanifest"),

		style.New(html.Set(string(css.Render(app)))),

		script_html.New(
			html.Set(`document.addEventListener('DOMContentLoaded', async function() { `+scripts+`});`),
		),
	)
}
