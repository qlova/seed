package main

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/script"
	"github.com/qlova/seeds/anime"
	"github.com/qlova/seeds/row"
	"github.com/qlova/seeds/text"
)

func main() {
	var App = seed.NewApp()

	var Row = row.AddTo(App)
	anime.Animate(text.AddTo(Row, "A"))
	anime.Animate(text.AddTo(Row, "B"))
	anime.Animate(text.AddTo(Row, "C"))

	App.OnClick(func(q script.Ctx) {
		anime.Push(q)
		{
			q.If(Row.Ctx(q).Row(), func() {
				Row.Ctx(q).SetColumn()
			}, q.Else(func() {
				Row.Ctx(q).SetRow()
			}))
		}
		anime.Pop(q)
	})

	App.Launch()
}
