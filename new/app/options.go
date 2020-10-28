package app

import (
	"fmt"
	"image/color"
	"strconv"

	"qlova.org/seed"
	"qlova.org/seed/assets"
	"qlova.org/seed/client"
	"qlova.org/seed/use/css"
	"qlova.org/seed/new/page"
)

func OnUpdateFound(do client.Script) seed.Option {
	return client.On("updatefound", do)
}

//SetLoadingPage sets the loading page of this app.
func SetLoadingPage(p page.Page) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		switch mode, _ := client.Seed(c); mode {
		case client.AddTo, client.Undo:
			panic("app.SetLoadingPage must not be called on a client.Seed")
		}

		var app app
		c.Load(&app)
		app.loadingPage = p
		c.Save(app)
	})
}

//SetColor sets the color of the app.
func SetColor(col color.Color) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		var app app
		c.Load(&app)

		switch mode, q := client.Seed(c); mode {
		case client.AddTo:
			fmt.Fprintf(q, `document.querySelector("meta[name=theme-color]").setAttribute("content", %v);`, css.RGB{Color: col}.Rule())
		case client.Undo:
			fmt.Fprintf(q, `document.querySelector("meta[name=theme-color]").setAttribute("content", %v);`, css.RGB{Color: app.color}.Rule())
		default:
			app.manifest.SetThemeColor(col)
			app.color = col
		}

		c.Save(app)
	})
}

//SetIcon sets the icon of the app.
func SetIcon(icon string) seed.Option {
	icon = assets.Path(icon)

	return seed.NewOption(func(c seed.Seed) {
		var app app
		c.Load(&app)

		switch mode, q := client.Seed(c); mode {
		case client.AddTo:
			fmt.Fprintf(q, `
			{
				let head = document.head || document.getElementsByTagName('head')[0];

				let link = document.createElement('link'),
				let oldLink = document.getElementById('dynamic-favicon');
				link.id = 'dynamic-favicon';
				link.rel = 'shortcut icon';
				link.href = %v;
				if (oldLink) {
					head.removeChild(oldLink);
				}
				head.appendChild(link);
			}
			`, strconv.Quote(icon))
		case client.Undo:
			fmt.Fprintf(q, `document.getElementById('dynamic-favicon').removeChild(oldLink);`)
		default:
			app.manifest.SetIcon(icon)
		}

		c.Save(app)
	})
}
