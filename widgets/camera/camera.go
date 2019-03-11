package camera

import "github.com/qlova/seed"

type Widget struct {
	seed.Seed
}

func New() Widget {
	widget := seed.New()

	widget.SetTag("video")
	widget.SetAttributes("autoplay")
	
	widget.OnReady(func(q seed.Script) {
		q.Javascript(`
			if(navigator.mediaDevices && navigator.mediaDevices.getUserMedia) {
				navigator.mediaDevices.getUserMedia({ video: true }).then(function(stream) {
					`+widget.Script(q).Element()+`.srcObject = stream;
					`+widget.Script(q).Element()+`.play();
				});
			}
		
		`)
	})

	return  Widget{widget}
}

//Create a new Text widget and add it to the provided parent.
func AddTo(parent seed.Interface) Widget {
	var widget = New()
	parent.Root().Add(widget)
	return widget
} 
