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

	seed.Javascript(seed.Element()+`.style.`+property+` = "`+value+`";`)
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

func (f File) Type() qlova.String {
	return f.seed.Qlovascript.Wrap(Javascript.String(f.expression+`.type`)).(qlova.String)
}

func (seed Seed) SetText(s qlova.String) {
	seed.Javascript(seed.Element()+`.textContent = `+s.Raw()+`;`)
}

func (seed Seed) SetLeft(s qlova.String) {
	seed.Javascript(seed.Element()+`.style.left = `+s.Raw()+`;`)
}

func (seed Seed) SetDisplay(s qlova.String) {
	seed.Javascript(seed.Element()+`.style.display = `+s.Raw()+`;`)
}

func (seed Seed) SetVisible() {
	seed.Javascript(seed.Element()+`.style.display = "inline-flex";`)
}

func (seed Seed) SetHidden() {
	seed.Javascript(seed.Element()+`.style.display = "none";`)
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

func (seed Seed) Left() qlova.String {
	return seed.Qlovascript.Wrap(Javascript.String(seed.Element()+`.style.left`)).(qlova.String)
}

func (seed Seed) Width() qlova.String {
	return seed.Qlovascript.Wrap(Javascript.String(`getComputedStyle(get("`+seed.ID+`")).width`)).(qlova.String)
}

func (seed Seed) SetValue(value qlova.String) {
	seed.Javascript(seed.Element()+`.value = `+value.Raw()+`;`)
}

func (seed Seed) SetClass(value qlova.String) {
	seed.Javascript(seed.Element()+`.className = `+value.Raw()+`;`)
}

func (seed Seed) Value() qlova.String {
	return seed.Qlovascript.Wrap(Javascript.String(seed.Element()+`.value`)).(qlova.String)
}

func (seed Seed) Text() qlova.String {
	return seed.Qlovascript.Wrap(Javascript.String(seed.Element()+`.innerText`)).(qlova.String)
}

func (seed Seed) File() File {
	return File{seed: seed, expression:seed.Element()+`.files[0]`}
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