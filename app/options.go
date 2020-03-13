package app

import (
	"fmt"
	"image/color"
	"strconv"

	"github.com/qlova/seed"
	"github.com/qlova/seed/asset"
	"github.com/qlova/seed/css"
	"github.com/qlova/seed/page"
)

//SetPage sets the starting page of this app.
func SetPage(p page.Page) seed.Option {
	return seed.NewOption(func(s seed.Any) {
		if app, ok := s.(App); ok {
			app.page = p
		}
	}, func(c seed.Ctx) {
		panic("cannot conditionally apply app.SetPage")
	}, func(c seed.Ctx) {
		panic("cannot conditionally apply app.SetPage")
	})
}

//SetLoadingPage sets the loading page of this app.
func SetLoadingPage(p page.Page) seed.Option {
	return seed.NewOption(func(s seed.Any) {
		if app, ok := s.(App); ok {
			app.loadingPage = p
		}
	}, func(c seed.Ctx) {
		panic("cannot conditionally apply app.SetPage")
	}, func(c seed.Ctx) {
		panic("cannot conditionally apply app.SetPage")
	})
}

//SetColor sets the color of the app.
func SetColor(col color.Color) seed.Option {
	var backup color.Color
	return seed.NewOption(func(s seed.Any) {
		if app, ok := s.(App); ok {
			backup = col
			app.manifest.SetThemeColor(col)
		}
	}, func(c seed.Ctx) {
		fmt.Fprintf(c.Ctx, `document.querySelector("meta[name=theme-color]").setAttribute("content", %v);`, css.RGB{Color: col}.Rule())
	}, func(c seed.Ctx) {
		fmt.Fprintf(c.Ctx, `document.querySelector("meta[name=theme-color]").setAttribute("content", %v);`, css.RGB{Color: backup}.Rule())
	})
}

//SetIcon sets the icon of the app.
func SetIcon(icon string) seed.Option {
	icon = asset.Path(icon)

	return seed.NewOption(func(s seed.Any) {
		if app, ok := s.(App); ok {
			app.manifest.SetIcon(icon)
		}
	}, func(c seed.Ctx) {
		fmt.Fprintf(c.Ctx, `
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
	}, func(c seed.Ctx) {
		fmt.Fprintf(c.Ctx, `document.getElementById('dynamic-favicon').removeChild(oldLink);`)
	})
}
