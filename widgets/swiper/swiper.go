package swiper

import "fmt"
import "encoding/json"

import "github.com/qlova/seed"
import "github.com/qlova/seed/script"
import qlova "github.com/qlova/script"

type Direction int

const Left Direction = -1
const Right Direction = 1

type Slide struct {
	index int

	seed.Seed
}

func init() {
	seed.Embed("/swiper.js", []byte(Javascript))
	seed.Embed("/swiper.css", []byte(CSS))
}

type Widget struct {
	slides int //The number of slides

	seed.Seed
	wrapper seed.Seed
}

//Returns gallery that displays 'local' images (in the assets directory).
func New(config ...Options) Widget {
	swiper := seed.New()

	swiper.Require("swiper.js")
	swiper.Require("swiper.css")

	wrapper := seed.AddTo(swiper)
	wrapper.SetClass("swiper-wrapper")

	pagination := seed.AddTo(swiper)
	pagination.SetClass("swiper-pagination")
	
	var options string
	if len(config) > 0 {
		
		config[0].Pagination = PaginationOptions{
			Element: pagination,
		}
		
		var JSON, err = json.Marshal(config[0])
		if err == nil {
			options = string(JSON)
		}
	}

	
	swiper.OnReady(func(q seed.Script) {
		q.Javascript(swiper.Script(q).Element() + `.swiper = new Swiper('#` + swiper.ID() + `', `+options+`);`)
	})

	return Widget{0, swiper, wrapper}
}

func AddTo(parent seed.Interface, config ...Options) Widget {
	var Swiper = New(config...)
	parent.Root().Add(Swiper)
	return Swiper
}

func (widget *Widget) NewSlide() Slide {
	var seed = seed.AddTo(widget.wrapper)
	seed.SetClass("swiper-slide")

	seed.Set("display", "flex")
	seed.Set("align-items", "center")
	seed.Set("justify-items", "center")
	seed.Set("text-align", "center")
	seed.Set("flex-direction", "column")

	widget.slides++

	return Slide{widget.slides - 1, seed}
}

type Script struct {
	script.Seed
}

func (w Widget) Script(q script.Script) Script {
	return Script{w.Seed.Script(q)}
}

func (s Script) Update() {
	s.Q.Javascript(s.Element() + ".swiper.update();")
}

func (s Script) Reset() {
	s.Q.Javascript(s.Element() + ".swiper.slideTo(0, 0);")
}

func (s Script) Goto(slide Slide) {
	s.Q.Javascript(s.Element() + ".swiper.slideTo(" + fmt.Sprint(slide.index) + ", 1000);")
}

func (s Script) Swipe(direction Direction) {
	if direction == Left {
		s.Q.Javascript(s.Element() + ".swiper.slidePrev();")
	}
	if direction == Right {
		s.Q.Javascript(s.Element() + ".swiper.slideNext();")
	}
}

func (s Script) Left() qlova.Bool {
	return s.Q.Value(s.Element() + ".swiper.isBeginning").Bool()
}

func (s Script) Right() qlova.Bool {
	return s.Q.Value(s.Element() + ".swiper.isEnd").Bool()
}
