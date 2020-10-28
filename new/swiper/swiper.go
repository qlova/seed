package swiper

import (
	"qlova.org/seed"
	"qlova.org/seed/client/clientrender"
	"qlova.org/seed/web/html"
	"qlova.org/seed/web/js"
	"qlova.org/seed/new/column"
	"qlova.org/seed/new/html/div"
	"qlova.org/seed/new/row"
	"qlova.org/seed/script"
)

type Slide struct {
	seed.Seed
}

func New(options ...seed.Option) seed.Seed {

	var Container = div.New()

	var slides seed.Options
	var others seed.Options
	var config Options
	config.Observer = true
	config.ObserveParents = true

	var process func(seed.Options)
	process = func(options seed.Options) {
		for _, option := range options {
			switch o := option.(type) {
			case CoverflowEffect:
				config.Effect = "coverflow"
				config.CoverflowEffect = o
			case seed.Options:
				process(o)
			case Slide:
				slides = append(slides, option)
			case Pagination:
				config.Pagination.Element = ".swiper-pagination"
				others = append(others, option)
			default:
				others = append(others, option)
			}
		}
	}

	process(options)

	return Container.With(
		html.AddClass("swiper-container"),

		row.New(
			js.Require("/swiper.js", javascript),
			html.AddClass("swiper-wrapper"),

			clientrender.On(js.Script(func(q js.Ctx) {
				q(`if (!` + html.Element(Container).String() + `.swiper)`)
				q(html.Element(Container).Set("swiper",
					js.NewValue("new Swiper(%v, %v)",
						html.Element(Container), js.ValueOf(config))))
				q(html.Element(Container).Run("swiper.update"))
			})),

			script.OnReady(js.Script(func(q js.Ctx) {
				q(`
				window.addEventListener("resize", function() {
					setTimeout(function() {
						if (` + html.Element(Container).String() + `.swiper)
						` + html.Element(Container).String() + `.swiper.update();
					}, 250);
				}, false);window.addEventListener("orientationchange", function() {
					setTimeout(function() {
						if (` + html.Element(Container).String() + `.swiper)
						` + html.Element(Container).String() + `.swiper.update();
					}, 250);
				}, false);`)
			})),

			seed.Options(slides),
		),

		seed.Options(others),
	)
}

func NewSlide(options ...seed.Option) Slide {
	return Slide{div.New(html.AddClass("swiper-slide"),
		column.New(seed.Options(options)),
	)}
}

type Pagination struct {
	seed.Seed
}

func NewPagination(options ...seed.Option) Pagination {
	return Pagination{div.New(
		html.AddClass("swiper-pagination"),
	)}
}
