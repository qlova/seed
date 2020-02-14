package script

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/qlova/script/language"
	Javascript "github.com/qlova/script/language/javascript"
	"github.com/qlova/seed/style/css"
)

func dashes2camels(s string) string {
	var camel string
	var parts = strings.Split(s, "-")
	for i, part := range parts {
		if i == 0 {
			camel += part
		} else {
			camel += strings.Title(part)
		}
	}
	return camel
}

//Seed is the script interface to a seed.Seed
type Seed struct {
	css.Style

	ID, Native string
	Q          Ctx
}

//Seed creates a new seed dynamically.
func (ctx Ctx) Seed(options ...string) Seed {
	var tag = "div"
	if len(options) > 0 {
		tag = options[0]
	}
	var variable = ctx.Value(`document.createElement(` + strconv.Quote(tag) + `)`).Native().Var()
	var seed = Seed{
		Native: variable.LanguageType().Raw(),
		Q:      ctx,
	}
	seed.Stylable = seed
	return seed
}

//Set is the required JS code for setting styles.
const Set = `
function set(element, property, value) {
	element.style[property] = value;
}
`

//CSS implements css.Stylable
func (seed Seed) CSS() css.Style {
	return css.Style{
		Stylable: seed,
	}
}

//Set sets the CSS style property to value.
func (seed Seed) Set(property, value string) {
	seed.Q.Require(Set)

	property = dashes2camels(property)

	seed.Javascript(`set(` + seed.Element() + `, "` + property + `", "` + value + `");`)
}

//SetDynamic sets the CSS style property to dynamic value.
func (seed Seed) SetDynamic(property, value string) {
	seed.Q.Require(Set)

	property = dashes2camels(property)

	seed.Javascript(`set(` + seed.Element() + `, "` + property + `", ` + value + `);`)
}

//Get returns the CSS property of this seed.
func (seed Seed) Get(property string) string {

	property = dashes2camels(property)

	return string(`getComputedStyle(` + seed.Element() + `).` + property)
}

//Data returns the data associated with this seed.
func (seed Seed) Data(property String) String {
	return seed.Q.Value(fmt.Sprint(seed.Element(), ".dataset[", property.LanguageType().Raw(), "]")).String()
}

//SetData sets string Data associated with this seed.
func (seed Seed) SetData(property, data String) {
	seed.Q.Javascript(fmt.Sprint(seed.Element(), ".dataset[", property.LanguageType().Raw(), "] = ", data.LanguageType().Raw(), ";"))
}

//Bytes TODO this is unimplemented.
func (seed Seed) Bytes() []byte {
	panic("unimplemented")
}

//Get is the required JS code for getting seeds.
const Get = `
	let get_cache = {};
	function get(id) {
		if (id in get_cache) {
			return get_cache[id];
		}

		let element = document.getElementById(id);
		
		if (element) {
			get_cache[id] = element;
			return element;
		}
		
		//Check the templates
		let templates = document.querySelectorAll('template');
		for (let template of templates) {
			element = template.content.getElementById(id);
			if (element) {
				get_cache[id] = element;
				return element;
			}
		}

		return null;
	}
`

//Element returns the seed as an HTML element.
func (seed Seed) Element() string {
	if seed.Native != "" {
		return seed.Native
	}

	seed.Q.Require(Get)

	return `get("` + seed.ID + `")`
}

//Javascript is shorthand for seed.Q.Javascript
func (seed Seed) Javascript(js string) {
	seed.Q.Javascript(language.Statement(js))
}

//File is a script interface to a file type.
type File struct {
	Q Ctx
	Native
}

//Type returns the type of the file.
func (f File) Type() String {
	return f.Q.Value(f.LanguageType().Raw() + `.type`).String()
}

//Name returns the name of the file.
func (f File) Name() String {
	return f.Q.Value(f.LanguageType().Raw() + `.name`).String()
}

//URL returns a url to the file.
func (f File) URL() Promise {
	return f.Q.Value(`new Promise(function (resolve, reject) {
		const reader = new FileReader();
		reader.onload = function() {
			resolve(reader.result);
		}
		reader.readAsDataURL(` + f.LanguageType().Raw() + `)
	});
	`).Promise()
}

func (seed Seed) wrap(s string) String {
	return seed.Q.StringFromLanguageType(Javascript.String{
		Expression: language.Statement(s),
	})
}

//Format is required to format text.
const Format = `function formatText(s) {
	let div = document.createElement('div');
	div.innerText = s;
	s = div.innerHTML;

	s.replace("\n", "<br>");
	s.replace("  ", "&nbsp;&nbsp;");
	s.replace("\t", "&emsp;");
	return s;
}`

//SetText sets the text of the string.
func (seed Seed) SetText(s String) {
	seed.Q.Require(Format)
	seed.Javascript(seed.Element() + `.innerHTML = formatText(` + raw(s) + `);`)
}

//SetPath sets the resource source/path of the seed.
func (seed Seed) SetPath(s String) {
	seed.Javascript(seed.Element() + `.src = "/"+` + raw(s) + `;`)
}

//SetSource sets the resource source/path of the seed.
func (seed Seed) SetSource(s String) {
	seed.Javascript(seed.Element() + `.src = "/"+` + raw(s) + `;`)
}

//Source returns the source of the seed.
func (seed Seed) Source() String {
	return seed.Q.Value(seed.Element() + `.src`).String()
}

//SetHTML sets the HTML of the seed.
func (seed Seed) SetHTML(s String) {
	seed.Javascript(seed.Element() + `.innerHTML = ` + raw(s) + `;`)
}

//SetHidden sets the seed to be hidden.
func (seed Seed) SetHidden() {
	seed.Javascript(`set(` + seed.Element() + `, "display", "none");`)
}

//Click simulates a click on this seed.
func (seed Seed) Click() {
	seed.Javascript(seed.Element() + `.click();`)
}

var unique int

//Unique returns a unique string suitable for variable names.
func Unique() string {
	unique++
	return fmt.Sprint("unique_", unique)
}

//Play calls play on the seed and returns the resulting promise.
func (seed Seed) Play() Promise {
	var variable = Unique()
	seed.Javascript(`let ` + variable + ` = ` + seed.Element() + `.play();`)
	return Promise{q: seed.Q, Native: seed.Q.Value(variable).Native()}
}

//Pause calls pause on the seed and returns the resulting promise.
func (seed Seed) Pause() {
	seed.Javascript(seed.Element() + `.pause();`)
}

//Focus focuses the seed for input.
func (seed Seed) Focus() {
	seed.Javascript(seed.Element() + `.focus();`)
}

//Blur removes focus from the seed.
func (seed Seed) Blur() {
	seed.Javascript(seed.Element() + `.blur();`)
}

//Restart calls the load method on the seed.
func (seed Seed) Restart() {
	seed.Javascript(seed.Element() + `.load();`)
}

//Width returns the width of the seed.
func (seed Seed) Width() Unit {
	return seed.Q.Value(`getComputedStyle(get("` + seed.ID + `")).width`).Unit()
}

//SetValue sets the input value of the seed.
func (seed Seed) SetValue(value String) {
	seed.Javascript(seed.Element() + `.value = ` + raw(value) + `;`)
}

//SetPlaceholder sets the input placeholder of the seed.
func (seed Seed) SetPlaceholder(value String) {
	seed.Javascript(seed.Element() + `.placeholder = ` + raw(value) + `;`)
}

//SetClass sets the class name of this seed.
func (seed Seed) SetClass(value String) {
	seed.Javascript(seed.Element() + `.className = ` + raw(value) + `;`)
}

//Value returns the input value of this seed.
func (seed Seed) Value() String {
	return seed.wrap(seed.Element() + `.value`)
}

//Text returns the text of this seed.
func (seed Seed) Text() String {
	return seed.wrap(seed.Element() + `.innerText`)
}

//URL returns the url/link of this seed.
func (seed Seed) URL() String {
	return seed.wrap(seed.Element() + `.href`)
}

//HTML returns the HTML of this seed.
func (seed Seed) HTML() String {
	return seed.wrap(seed.Element() + `.innerHTML`)
}

//File returns the first file of this seed.
func (seed Seed) File() File {
	return File{seed.Q, seed.Q.NativeFromLanguageType(Javascript.Native{
		Expression: seed.Element() + `.files[0]`}).Var()}
}

//Load sets the resource location of this seed to the specified file.
func (seed Seed) Load(f File) {
	seed.Javascript(seed.Element() + `.src = window.URL.createObjectURL(` + f.LanguageType().Raw() + `);`)
}

//Add a child seed to this seed.
func (seed Seed) Add(child Seed) {
	seed.Javascript(seed.Element() + `.appendChild(` + child.Element() + `);`)
}

//OnClick sets the onclick event of this seed.
func (seed Seed) OnClick(f func()) {
	seed.Javascript(seed.Element() + `.onclick = async function() {`)
	f()
	seed.Javascript(`};`)
}

//Filter runs a function on each child of this seed.
func (seed Seed) Filter(f func(child Seed)) {
	seed.Q.Javascript(`for (let i = 0; i < ` + seed.Element() + `.children.length; i++) {`)
	f(Seed{
		Native: seed.Element() + `.children[i]`,
		Q:      seed.Q,
	})
	seed.Q.Javascript(`}`)
}
