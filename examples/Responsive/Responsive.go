package main

import (
	"github.com/qlova/seed/app"
	"github.com/qlova/seed/style"

	"github.com/qlova/seed/s/text"
)

func main() {
	app.New("Responsive",
		text.New("Tiny", style.SetHidden(), style.If.Tiny(style.SetVisible())),
		text.New("Small", style.SetHidden(), style.If.Small(style.SetVisible())),
		text.New("Medium", style.SetHidden(), style.If.Medium(style.SetVisible())),
		text.New("Large", style.SetHidden(), style.If.Large(style.SetVisible())),
		text.New("Huge", style.SetHidden(), style.If.Huge(style.SetVisible())),

		text.New("Portrait", style.SetHidden(), style.If.Portrait(style.SetVisible())),
		text.New("Landscape", style.SetHidden(), style.If.Landscape(style.SetVisible())),
	).Launch()
}
