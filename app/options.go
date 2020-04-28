package app

import (
	"image/color"
	"strconv"

	"github.com/qlova/seed"
	"github.com/qlova/seed/asset"
	"github.com/qlova/seed/css"
	"github.com/qlova/seed/page"
	"github.com/qlova/seed/script"
)

func OnUpdateFound(do script.Script) seed.Option {
	return script.On("updatefound", do)
}

//SetPage sets the starting page of this app.
func SetPage(p page.Page) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		switch c.(type) {
		case script.Seed, script.Undo:
			panic("app.SetPage must not be called on a script.Seed")
		}

		var app app
		c.Read(&app)
		app.page = p
		c.Write(app)
	})
}

//SetLoadingPage sets the loading page of this app.
func SetLoadingPage(p page.Page) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		switch c.(type) {
		case script.Seed, script.Undo:
			panic("app.SetLoadingPage must not be called on a script.Seed")
		}

		var app app
		c.Read(&app)
		app.loadingPage = p
		c.Write(app)
	})
}

//SetColor sets the color of the app.
func SetColor(col color.Color) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		var app app
		c.Read(&app)

		switch q := c.(type) {
		case script.Seed:
			q.Javascript(`document.querySelector("meta[name=theme-color]").setAttribute("content", %v);`, css.RGB{Color: col}.Rule())
		case script.Undo:
			q.Javascript(`document.querySelector("meta[name=theme-color]").setAttribute("content", %v);`, css.RGB{Color: app.color}.Rule())
		default:
			app.manifest.SetThemeColor(col)
			app.color = col
		}

		c.Write(app)
	})
}

//SetIcon sets the icon of the app.
func SetIcon(icon string) seed.Option {
	icon = asset.Path(icon)

	return seed.NewOption(func(c seed.Seed) {
		var app app
		c.Read(&app)

		switch q := c.(type) {
		case script.Seed:
			q.Javascript(`
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
		case script.Undo:
			q.Javascript(`document.getElementById('dynamic-favicon').removeChild(oldLink);`)
		default:
			app.manifest.SetIcon(icon)
		}

		c.Write(app)
	})
}
