package seed

import (
	"github.com/qlova/seed/script"
	"github.com/qlova/seed/style"
	"github.com/qlova/seed/style/css"
	"github.com/qlova/seed/unit"
	"github.com/qlova/seed/user"
	"github.com/russross/blackfriday/v2"

	"encoding/base64"
	"html"
	"math/big"
	"os"
	"path/filepath"
	"strings"

	HTML "github.com/qlova/seed/html"
)

//Interface is anything that has a Root() seed method.
type Interface interface {
	Root() Seed
}

//Dir is the working directory of the seed.
var Dir = filepath.Dir(os.Args[0])

//User is an alias to the user.User type.
type User = user.Ctx

//Arial is a default arial font.
var Arial = style.Font{
	FontFace: css.FontFace{
		FontFamily: "Arial",
	},
}

//Font is a font type.
type Font string

//NewFont registers the font and creates it.
func NewFont(path string) Font {
	return Font("/" + path)
}

//SetFont sets the font of the specified seed.
func (seed Seed) SetFont(font Font) {
	seed.font = font
}

//Seed is a component of an app.
//It can contain children seeds and have styles, scripts attached to it.
type Seed struct {
	*seed
}

type seed struct {
	HTML.Element

	//Default, Portrait, Landscape.
	style.Group
	//Watch, Phone, Tablet, Laptop, Desktop
	Tiny, Small, Medium, Large, Huge style.Group

	dynamic

	id               string
	tag, attr, class string
	children         []Interface

	attributes map[string]string

	styled bool
	ready  bool

	font      Font
	animation Animation

	content []byte
	page    bool
	splash  bool

	setup func()

	onclick func(script.Ctx)
	onready func(script.Ctx)

	on map[string]func(script.Ctx)

	Template, TemplateRoot bool

	parent Interface

	//This is a list of scripts that are needed by this seed.
	//eg. []string{"jquery.js"}
	scripts []string

	//Special styles.
	query *query

	screenSmallerThan, screenGreaterThan map[Unit]style.Style

	desktop, mobile, tablet, watch, tv, native Seed

	app *App

	assets []Asset

	states map[State]func(script.Ctx)

	tags map[string]bool
}

//On runs a script callback when the specified event is fired.
func (seed Seed) On(event string, callback func(script.Ctx)) {
	if seed.on == nil {
		seed.on = make(map[string]func(script.Ctx))
	}

	if original, ok := seed.on[event]; ok {
		seed.on[event] = func(q script.Ctx) {
			original(q)
			callback(q)
		}
		return
	}
	seed.on[event] = callback
}

//Ready returns the ready event handler.
func (seed Seed) Ready() func(script.Ctx) {
	return func(q script.Ctx) {
		if seed.onready != nil {
			seed.onready(q)
		}

		for event, handler := range seed.on {
			q.Javascript(seed.Ctx(q).Element() + ".on" + event + " = async function() {")
			handler(q)
			q.Javascript("};")
		}
	}
}

//Is returns true if a and b are the same seed.
func (seed Seed) Is(b Interface) bool {
	return seed.seed == b.Root().seed
}

//Unit is a measurement for styles.
type Unit = complex128

//ScreenSmallerThan is a conditional query that returns a style that effects that condition.
func (seed Seed) ScreenSmallerThan(u unit.Unit) style.Style {
	if seed.screenSmallerThan == nil {
		seed.screenSmallerThan = make(map[unit.Unit]style.Style)
	}

	if s, ok := seed.screenSmallerThan[u]; ok {
		return s
	}
	var s = style.New()
	seed.screenSmallerThan[u] = s
	return s
}

//Null returns true if the seed is null.
func (seed Seed) Null() bool {
	return seed.seed == nil
}

//MarshalText marshals this seed as a HTML id.
func (seed Seed) MarshalText() ([]byte, error) {
	return []byte("#" + seed.ID()), nil
}

//ID returns the id of this seed.
func (seed Seed) ID() string {
	if seed.seed == nil {
		return ""
	}
	return seed.id
}

//Children returns a slice of children of this seed.
func (seed Seed) Children() []Interface {
	return seed.children
}

//SetClass sets the HTML class of this seed [DEPRECIATED].
func (seed Seed) SetClass(class string) {
	seed.class = class
}

//Tag returns the HTML tag of this seed [DEPRECIATED].
func (seed Seed) Tag() string {
	return seed.tag
}

//SetTag sets the HTML tag for this seed [DEPRECIATED].
func (seed Seed) SetTag(tag string) {
	seed.tag = tag
}

//SetAttributes sets the HTML attributes for this seed [DEPRECIATED].
func (seed Seed) SetAttributes(attr string) {
	seed.attr = attr
}

//Attributes returns the HTML attributes for this seed [DEPRECIATED].
func (seed Seed) Attributes() string {
	return seed.attr
}

//AddTo adds this seed to the specified parent.
func (seed Seed) AddTo(parent Interface) Seed {
	parent.Root().Add(seed)
	return seed
}

func (seed Seed) clone() Seed {
	var clone = New()
	clone.id = seed.id
	clone.tag = seed.tag
	clone.attr = seed.attr
	clone.class = seed.class
	clone.content = seed.content
	clone.parent = seed.parent

	return clone
}

//Desktop returns the seed that should replace this seed when on the Desktop.
func (seed Seed) Desktop() Seed {
	if seed.desktop.seed == nil {
		seed.desktop = seed.clone()
	}
	return seed.desktop
}

//ReactNative returns the seed that should replace this seed when on ReactNative [EXPERIMENTAL].
func (seed Seed) ReactNative() Seed {
	if seed.native.seed == nil {
		seed.native = seed.clone()
	}
	return seed.native
}

//Root returns the seed itself, when embedded in a struct, this is good way to retrieve the original seed.
func (seed Seed) Root() Seed {
	return seed
}

//Parent returns the parent seed.
func (seed Seed) Parent() Seed {
	return seed.parent.Root()
}

//Copy creates a copy of this seed.
func (seed Seed) Copy() Seed {
	var another = *seed.seed

	another.Style = another.Style.Copy()
	another.Portrait = another.Portrait.Copy()
	another.Landscape = another.Landscape.Copy()

	//TODO copy media queries?

	another.id = base64.RawURLEncoding.EncodeToString(big.NewInt(id).Bytes())
	id++
	return Seed{seed: &another}
}

//All seeds have a unique id.
var id int64 = 1

var allSeeds = make(map[string]*seed)

//New create and return a new seed.
func New() Seed {
	s := new(seed)

	//Seed identification is compressed to base64.
	s.id = base64.RawURLEncoding.EncodeToString(big.NewInt(id).Bytes())

	if s.id[0] >= '0' && s.id[0] <= '9' {
		s.id = "_" + s.id
	}

	s.id = strings.Replace(s.id, "-", "__", -1)

	id++

	s.Group.Init()
	s.Tiny.Init()
	s.Small.Init()
	s.Medium.Init()
	s.Large.Init()
	s.Huge.Init()

	s.tag = "div"

	allSeeds[s.id] = s

	//Intial style.
	//seed.SetSize(100, 100)

	return Seed{seed: s}
}

//AddTo creates and returns a seed attached to the specified parent.
func AddTo(parent Interface) Seed {
	var seed = New()
	parent.Root().Add(seed)
	return seed
}

//OnPress is the JS required for OnPress handling.
const OnPress = `
	function op(element, func, propagate) {
		let handler = async function(event) {
			await func(event);
		};
		
		let moved = false;
		let point = [0, 0];
		
		element.ontouchstart = function(e) {
			var changedTouch = event.changedTouches[0];
				point[0]  = changedTouch.clientX;
				point[1]  = changedTouch.clientY;
		};
		
		element.ontouchmove = function(event) {
			var changedTouch = event.changedTouches[0];
			var elem = document.elementFromPoint(changedTouch.clientX, changedTouch.clientY);
						
			if (elem != event.target) moved = true;
						
			let a = changedTouch.clientX - point[0];
			let b = changedTouch.clientY - point[1];
			if ((a*a + b*b) > 50*50) moved = true;
		};
		
		element.ontouchend = async function(ev) {
			if (ev.stopPropagation && !propagate) ev.stopPropagation(); 
			ev.preventDefault(); 
			if (moved) {
				moved = false; 
				return; 
			}
			ev = ev.changedTouches[0];
			await handler(ev);
		};

		element.onclick = handler;
	}
`

//OnClick runs a script when this seed is clicked.
func (seed Seed) OnClick(f func(script.Ctx)) {
	seed.onclick = f
	seed.OnReady(func(q script.Ctx) {
		q.Require(OnPress)
		q.Javascript("op(" + seed.Ctx(q).Element() + ", async function(event) {")
		f(q)
		q.Javascript("});")
	})
}

//OnClickThrough runs a script when this seed is clicked, allows click to propagate to other scripts.
func (seed Seed) OnClickThrough(f func(script.Ctx)) {
	seed.onclick = f
	seed.OnReady(func(q script.Ctx) {
		q.Require(OnPress)
		q.Javascript("op('" + seed.id + "', async function(event) {")
		f(q)
		q.Javascript("}, true);")
	})
}

//OnClickGoto is shorthand for seed.OnClick(func(q script.Ctx){ page.Ctx(q).Goto() })
func (seed Seed) OnClickGoto(page Page) {
	seed.OnClick(func(q script.Ctx) {
		page.Ctx(q).Goto()
	})
}

//OnReady runs a script when this seed is ready/loaded/onload/init.
func (seed Seed) OnReady(f func(script.Ctx)) {
	if seed.onready == nil {
		seed.onready = f
	} else {
		var old = seed.onready
		seed.onready = func(q script.Ctx) {
			old(q)
			f(q)
		}
	}
}

//OnChange runs a script when this seed's value is changed by the user.
func (seed Seed) OnChange(f func(script.Ctx)) {
	seed.On("change", func(q script.Ctx) {
		f(q)
	})
}

//Page returns true if this seed is a page.
func (seed Seed) Page() bool {
	return seed.page
}

//Require requires an external script needed by this seed.
func (seed Seed) Require(script string) {
	seed.scripts = append(seed.scripts, script)

	if len(script) > 0 && script[0] == '/' {
		NewAsset(script).AddTo(seed)
	}
}

//Add a child seed to this seed.
func (seed Seed) Add(child Interface) {
	seed.children = append(seed.children, child)
	child.Root().parent = seed
	child.Root().app = seed.app
	if seed.Template || seed.TemplateRoot {
		child.Root().Template = true
	}
}

//SetContent adds text, html or whatever!
func (seed Seed) SetContent(data string) {
	seed.content = []byte(data)
}

//SetHTML sets the html content of the seed.
func (seed Seed) SetHTML(data string) {
	seed.content = []byte(data)
}

//HTML returns the html content of the seed.
func (seed Seed) HTML() []byte {
	return seed.content
}

//Text returns the text content of the seed.
func (seed Seed) Text() string {
	return string(seed.content)
}

//SetText sets the text content of this seed.
func (seed Seed) SetText(data string) {
	data = html.EscapeString(data)
	data = strings.Replace(data, "\n", "<br>", -1)
	data = strings.Replace(data, "  ", "&nbsp;&nbsp;", -1)
	data = strings.Replace(data, "\t", "&emsp;", -1)
	seed.content = []byte(data)
}

//SetMarkdown sets the content of the seed in markdown format.
func (seed Seed) SetMarkdown(data string) {
	seed.content = blackfriday.Run([]byte(data))
}

//OnSwipeLeft runs a script when swiped left.
func (seed Seed) OnSwipeLeft(f func(script.Ctx)) {
	seed.Require("hammer.js")
	seed.OnReady(func(q script.Ctx) {
		q.Javascript("{")
		q.Javascript("let hammertime = new Hammer(" + seed.Ctx(q).Element() + ");")
		q.Javascript(`hammertime.on("swipeleft", async function() {`)
		f(q)
		q.Javascript("});")
		q.Javascript("}")
	})
}

//OnSwipeRight runs a script when swiped right.
func (seed Seed) OnSwipeRight(f func(script.Ctx)) {
	seed.Require("hammer.js")
	seed.OnReady(func(q script.Ctx) {
		q.Javascript("{")
		q.Javascript("let hammertime = new Hammer(" + seed.Ctx(q).Element() + ");")
		q.Javascript(`hammertime.on("swiperight", async function() {`)
		f(q)
		q.Javascript("});")
		q.Javascript("}")
	})
}

//OnFocus run a script when this seed is focused.
func (seed Seed) OnFocus(f func(script.Ctx)) {
	seed.On("focus", func(q script.Ctx) {
		f(q)
	})
}

//OnInput runs a script when this seed has input.
func (seed Seed) OnInput(f func(script.Ctx)) {
	seed.On("input", func(q script.Ctx) {
		f(q)
	})
}

//OnEnter runs a script when this seed has enter input.
func (seed Seed) OnEnter(f func(script.Ctx)) {
	seed.OnReady(func(q script.Ctx) {
		q.Javascript("{")
		q.Javascript(`let onenter = async function(ev) {if (ev.keyCode == 13 || ev.which == 13){ `)
		f(q)
		q.Javascript(`}};`)
		q.Javascript(seed.Ctx(q).Element() + `.onkeypress = onenter;`)
		q.Javascript("}")
	})
}

//OnFocusLost runs a script when this seed has focus lost.
func (seed Seed) OnFocusLost(f func(script.Ctx)) {
	seed.On("focusout", func(q script.Ctx) {
		f(q)
	})
}

//LongPress is the JS required for handling longpreses.
const LongPress = `
!function(t,e){"use strict";function n(){this.dispatchEvent(new CustomEvent("long-press",{bubbles:!0,cancelable:!0})),clearTimeout(o)}var o=null,u="ontouchstart"in t||navigator.MaxTouchPoints>0||navigator.msMaxTouchPoints>0,s=u?"touchstart":"mousedown",i=u?"touchcancel":"mouseout",a=u?"touchend":"mouseup",c=u?"touchmove":"mousemove";"initCustomEvent"in e.createEvent("CustomEvent")&&(t.CustomEvent=function(t,n){n=n||{bubbles:!1,cancelable:!1,detail:void 0};var o=e.createEvent("CustomEvent");return o.initCustomEvent(t,n.bubbles,n.cancelable,n.detail),o},t.CustomEvent.prototype=t.Event.prototype),e.addEventListener(s,function(t){var e=t.target,u=parseInt(e.getAttribute("data-long-press-delay")||"500",10);o=setTimeout(n.bind(e),u)}),e.addEventListener(a,function(t){clearTimeout(o)}),e.addEventListener(i,function(t){clearTimeout(o)}),e.addEventListener(c,function(t){clearTimeout(o)})}(this,document);
`

//OnLongPress runs a script when this seed is long-clicked.
func (seed Seed) OnLongPress(f func(script.Ctx)) {
	seed.OnReady(func(q script.Ctx) {
		q.Require(LongPress)
		q.Javascript("{")
		q.Javascript(`let onlongpress = async function(ev) {ev.preventDefault();`)
		f(q)
		q.Javascript(`};`)
		q.Javascript(seed.Ctx(q).Element() + `.addEventListener('long-press', onlongpress);`)
		q.Javascript("}")
	})
}
