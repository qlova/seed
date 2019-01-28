package video

import "github.com/qlova/seed"

type Widget struct {
	seed.Seed
}

func New(path ...string) Widget {
	widget := seed.New()

	widget.SetTag("video")
	if len(path) > 0 {
		widget.SetAttributes("src='"+path[0]+"' playsinline preload='auto'")
		seed.NewAsset(path[0]).AddTo(widget)
	} else {
		widget.SetAttributes("playsinline preload='auto'")
	}

	return  Widget{widget}
}

//Create a new Text widget and add it to the provided parent.
func AddTo(parent seed.Interface, path ...string) Widget {
	var widget = New(path...)
	parent.Root().Add(widget)
	return widget
} 
