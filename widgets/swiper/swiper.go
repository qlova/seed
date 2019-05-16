package swiper

import "github.com/qlova/seed"
import "github.com/qlova/seed/script"
import qlova "github.com/qlova/script"

type Direction int

const Left Direction = -1
const Right Direction = 1

func init() {
	seed.Embed("/swiper.js", []byte(Javascript))
	seed.Embed("/swiper.css", []byte(CSS))
}

type Widget struct {
	seed.Seed
	wrapper seed.Seed
}

//Returns gallery that displays 'local' images (in the assets directory).
func New(images ...string) Widget {
	swiper := seed.New()

	swiper.Require("swiper.js")
	swiper.Require("swiper.css")

	wrapper := seed.AddTo(swiper)
	wrapper.SetClass("swiper-wrapper")

	pagination := seed.AddTo(swiper)
	pagination.SetClass("swiper-pagination")
	
	
	swiper.OnReady(func(q seed.Script) {
		q.Javascript(swiper.Script(q).Element()+`.swiper = new Swiper('#`+swiper.ID()+`', {pagination: {el: '#`+pagination.ID()+`'}});`)
	})
	
	return Widget{swiper, wrapper}
}

func AddTo(parent seed.Interface) Widget {
	var Swiper = New()
	parent.Root().Add(Swiper)
	return Swiper
}

func (widget *Widget) NewSlide() seed.Seed {
	var seed = seed.AddTo(widget.wrapper)
		seed.SetClass("swiper-slide")

	seed.Set("display", "flex")
	seed.Set("align-items", "center")
	seed.Set("justify-items", "center")
	seed.Set("text-align", "center")
	seed.Set("flex-direction", "column")

	return seed
}

type Script struct {
	script.Seed
}

func (w Widget) Script(q script.Script) Script {
	return Script{w.Seed.Script(q)}
}

func (s Script) Update() {
	s.Q.Javascript(s.Element()+".swiper.update();")
}

func (s Script) Reset() {
	s.Q.Javascript(s.Element()+".swiper.slideTo(0, 0);")
}

func (s Script) Swipe(direction Direction) {
	if direction == Left {
		s.Q.Javascript(s.Element()+".swiper.slidePrev();")
	}
	if direction == Right {
		s.Q.Javascript(s.Element()+".swiper.slideNext();")
	}
}

func (s Script) Left() qlova.Bool {
	return s.Q.Value(s.Element()+".swiper.isBeginning").Bool()
}

func (s Script) Right() qlova.Bool {
	return s.Q.Value(s.Element()+".swiper.isEnd").Bool()
}