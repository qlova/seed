package script

import (
	"fmt"
	"strings"
)

import qlova "github.com/qlova/script"
import "github.com/qlova/script/language"
import "github.com/qlova/script/language/javascript"

import "github.com/qlova/seed/style/css"

type String = qlova.String
type Object string

type Expression struct {
	seed Seed
	expression string
}



type Promise Expression


func (p Promise) Raw() string {
	return p.expression
}

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

type Seed struct {
	css.Style

	ID, Native string
	Qlovascript qlova.Script
}

func (seed Seed) Set(property, value string) {
	property = dashes2camels(property)

	seed.Javascript(`set(`+seed.Element()+`, "`+property+`", "`+value+`");`)
}

func (seed Seed) Get(property string) string {

	property = dashes2camels(property)

	return string(`getComputedStyle(`+seed.Element()+`).`+property)
}

//TODO
func (seed Seed) Bytes() []byte {
	return nil
}

func (seed Seed) Element() string {
	if seed.Native != "" {
		return seed.Native
	}
	return `get("`+seed.ID+`")`
}

func (seed Seed) Javascript(js string) {
	seed.Qlovascript.Raw("Javascript", language.Statement(js))
}

type File Expression

func (f File) Type() String {
	return f.seed.wrap(f.expression+`.type`)
}

func (f File) Name() String {
	return f.seed.wrap(f.expression+`.name`)
}

func raw(s String) string {
	return string(s.LanguageType().(Javascript.String).Expression)
}

func (seed Seed) wrap(s string) String {
	return seed.Qlovascript.StringFromLanguageType(Javascript.String{
		Expression: language.Statement(s),
	})
}

func (seed Seed) SetText(s String) {	
	seed.Javascript(seed.Element()+`.textContent = `+raw(s)+`;`)
}

func (seed Seed) SetPath(s String) {
	seed.Javascript(seed.Element()+`.src = `+raw(s)+`;`)
}
func (seed Seed) SetSource(s String) {
	seed.Javascript(seed.Element()+`.src = `+raw(s)+`;`)
}

func (seed Seed) SetHTML(s String) {
	seed.Javascript(seed.Element()+`.innerHTML = `+raw(s)+`;`)
}

func (seed Seed) SetLeft(s String) {
	seed.Javascript(`set(`+seed.Element()+`, "left", `+raw(s)+`);`)
}

func (seed Seed) SetDisplay(s String) {
	seed.Javascript(`set(`+seed.Element()+`, "display", `+raw(s)+`);`)
}

func (seed Seed) SetVisible() {
	seed.Javascript(`set(`+seed.Element()+`, "display", "inline-flex");`)
}

func (seed Seed) SetHidden() {
	seed.Javascript(`set(`+seed.Element()+`, "display", "none");`)
}

func (seed Seed) Click() {
	seed.Javascript(seed.Element()+`.click();`)
}

var unique int
func Unique() string {
	unique++
	return fmt.Sprint("unique_", unique)
}

func (seed Seed) Play() Promise {
	var variable = Unique() 
	seed.Javascript(`let `+variable+` = `+seed.Element()+`.play();`)
	return Promise{seed:seed, expression: variable}
}

func (seed Seed) Pause() {
	seed.Javascript(seed.Element()+`.pause();`)
}

func (seed Seed) Focus() {
	seed.Javascript(seed.Element()+`.focus();`)
}

func (seed Seed) Restart() {
	seed.Javascript(seed.Element()+`.load();`)
}

func (seed Seed) Left() String {
	return seed.wrap(seed.Element()+`.style.left`)
}

func (seed Seed) Width() String {
	return seed.wrap(`getComputedStyle(get("`+seed.ID+`")).width`)
}

func (seed Seed) SetValue(value String) {
	seed.Javascript(seed.Element()+`.value = `+raw(value)+`;`)
}

func (seed Seed) SetPlaceholder(value String) {
	seed.Javascript(seed.Element()+`.placeholder = `+raw(value)+`;`)
}

func (seed Seed) SetClass(value String) {
	seed.Javascript(seed.Element()+`.className = `+raw(value)+`;`)
}

func (seed Seed) Value() String {
	return seed.wrap(seed.Element()+`.value`)
}

func (seed Seed) Text() String {
	return seed.wrap(seed.Element()+`.innerText`)
}

func (seed Seed) Location() String {
	return seed.wrap(seed.Element()+`.href`)
}

func (seed Seed) Data(key string) String {
	return seed.wrap(seed.Element()+`.data["`+key+`"]`)
}

func (seed Seed) HTML() String {
	return seed.wrap(seed.Element()+`.innerHTML`)
}

func (seed Seed) File() File {
	return File{seed: seed, expression:seed.Element()+`.files[0]`}
}

func (seed Seed) Display() String {
	return seed.wrap(seed.Element()+`.style.display`)
}

//Temporary method DEPRECIATED
func (f File) Raw() string {
	return f.expression
}

func (seed Seed) Load(f File) {
	seed.Javascript(seed.Element()+`.src = window.URL.createObjectURL(`+f.expression+`);`)
}

//Add a child seed to this seed.
func (seed Seed) Add(child Seed) {
	seed.Javascript(seed.Element()+`.appendChild(`+child.Element()+`);`)
}

func (seed Seed) OnClick(f func()) {
	seed.Javascript(seed.Element()+`.onclick = function() {`)
	f()
	seed.Javascript(`};`)
}

//Animations
func (seed Seed) SlideInFrom(direction complex128) {

	if direction == 1i {
	
		seed.Javascript(`set(get(last_page), "display", "inline-flex");`)
		seed.Javascript(`set(`+seed.Element()+`, "z-index", "50");`)
		seed.Javascript(seed.Element()+`.style.transform = "translateY(100vh)";`)
		seed.Javascript(seed.Element()+`.style.transition = "transform 0.5s";`)
		
		seed.Javascript(`window.requestAnimationFrame(function() {window.requestAnimationFrame(function() {`)
			seed.Javascript(seed.Element()+`.style.transform = "translateY(0vh)";`)
			seed.Javascript(`setTimeout(function() { set(get(last_page), "display", "none"); set(`+seed.Element()+`, "z-index", "initial"); }, 500);`)
		seed.Javascript(`})})`)
	}

	if direction == 1 {
		seed.Javascript(`set(get(last_page), "display", "inline-flex");`)
		seed.Javascript(`set(`+seed.Element()+`, "z-index", "50");`)
		seed.Javascript(seed.Element()+`.style.transform = "translateX(100vw)";`)
		seed.Javascript(seed.Element()+`.style.transition = "transform 0.5s";`)
		
		seed.Javascript(`window.requestAnimationFrame(function() {window.requestAnimationFrame(function() {`)
			seed.Javascript(seed.Element()+`.style.transform = "translateX(0vw)";`)
			seed.Javascript(`setTimeout(function() { set(get(last_page), "display", "none"); set(`+seed.Element()+`, "z-index", ""); }, 500);`)
		seed.Javascript(`})})`)
	}

	if direction == -1 {
		seed.Javascript(`set(get(last_page), "display", "inline-flex");`)
		seed.Javascript(`set(`+seed.Element()+`, "z-index", "50");`)
		seed.Javascript(seed.Element()+`.style.transform = "translateX(-100vw)";`)
		seed.Javascript(seed.Element()+`.style.transition = "transform 0.5s";`)
		
		seed.Javascript(`window.requestAnimationFrame(function() {window.requestAnimationFrame(function() {`)
			seed.Javascript(seed.Element()+`.style.transform = "translateX(0vw)";`)
			seed.Javascript(`setTimeout(function() { set(get(last_page), "display", "none"); set(`+seed.Element()+`, "z-index", ""); }, 500);`)
		seed.Javascript(`})})`)
	}
}

//Animations
func (seed Seed) SlideOutFrom(direction complex128) {
	/*seed.Javascript(`set(`+seed.Element()+`, "display", "inline-flex");`)
	seed.Javascript(`set(`+seed.Element()+`, "z-index", "50");`)
	seed.Javascript(`set(`+seed.Element()+`, "position", "fixed");`)
	seed.Javascript(`set(`+seed.Element()+`, "top", "0");`)
	seed.Javascript(`set(`+seed.Element()+`, "left", "0");`)
	seed.Javascript(`set(`+seed.Element()+`, "transition", "top 0.5s");`)
	seed.Javascript(`setTimeout(function() { set(`+seed.Element()+`, "top", "100vh"); }, 30);`)
	seed.Javascript(`setTimeout(function() { set(`+seed.Element()+`, "display", "none"); set(`+seed.Element()+`, "z-index", "initial"); }, 500);`)*/

	if direction == 1i {
		seed.Javascript(`set(`+seed.Element()+`, "display", "inline-flex");`)
		seed.Javascript(`set(`+seed.Element()+`, "z-index", "50");`)
		seed.Javascript(seed.Element()+`.style.transform = "translateY(0vh)";`)
		seed.Javascript(seed.Element()+`.style.transition = "transform 0.5s";`)
		
		seed.Javascript(`window.requestAnimationFrame(function() {window.requestAnimationFrame(function() {`)
			seed.Javascript(seed.Element()+`.style.transform = "translateY(100vh)";`)
			seed.Javascript(`setTimeout(function() { set(`+seed.Element()+`, "display", "none"); set(`+seed.Element()+`, "z-index", ""); }, 500);`)
		seed.Javascript(`})})`)
	}

	if direction == 1 {
		seed.Javascript(`set(`+seed.Element()+`, "display", "inline-flex");`)
		seed.Javascript(`set(`+seed.Element()+`, "z-index", "50");`)
		seed.Javascript(seed.Element()+`.style.transform = "translateX(0vw)";`)
		seed.Javascript(seed.Element()+`.style.transition = "transform 0.5s";`)
		
		seed.Javascript(`window.requestAnimationFrame(function() {window.requestAnimationFrame(function() {`)
			seed.Javascript(seed.Element()+`.style.transform = "translateX(100vw)";`)
			seed.Javascript(`setTimeout(function() { set(`+seed.Element()+`, "display", "none"); set(`+seed.Element()+`, "z-index", ""); }, 500);`)
		seed.Javascript(`})})`)
	}
}
