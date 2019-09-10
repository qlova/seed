package seed

import (
	"bytes"
	"fmt"
)

import "github.com/qlova/seed/style"

type Platform int

const (
	Default Platform = iota

	Desktop
	Mobile
	Tablet
	Watch
	Tv
	Playstation
	Xbox
)

func (seed Seed) ShortCircuit(platform Platform) Seed {
	if platform == Desktop && seed.desktop.seed != nil {
		return seed.desktop
	}
	return Seed{}
}

func (seed Seed) buildStyleSheet(platform Platform, sheet *style.Sheet) {
	if short := seed.ShortCircuit(platform); short.seed != nil {
		short.buildStyleSheet(platform, sheet)
		return
	}

	//seed.postProduction()
	if data := seed.Style.Bytes(); data != nil {
		seed.styled = true
		if seed.template {
			sheet.Add("."+seed.id, seed.Style)
		} else {
			sheet.Add("#"+seed.id, seed.Style)
		}
	}
	for _, child := range seed.children {
		child.Root().buildStyleSheet(platform, sheet)
	}
}

func (seed Seed) BuildStyleSheet(platform Platform) style.Sheet {
	var stylesheet = make(style.Sheet)
	seed.buildStyleSheet(platform, &stylesheet)
	return stylesheet
}

func (seed Seed) buildStyleSheetForLandscape(platform Platform, sheet *style.Sheet) {
	if short := seed.ShortCircuit(platform); short.seed != nil {
		short.buildStyleSheetForLandscape(platform, sheet)
		return
	}

	//seed.postProduction()
	if data := seed.Landscape.Bytes(); data != nil {
		seed.styled = true
		if seed.template {
			sheet.Add("."+seed.id, seed.Landscape)
		} else {
			sheet.Add("#"+seed.id, seed.Landscape)
		}
	}
	for _, child := range seed.children {
		child.Root().buildStyleSheetForLandscape(platform, sheet)
	}
}

func (seed Seed) BuildStyleSheetForLandscape(platform Platform) style.Sheet {
	var stylesheet = make(style.Sheet)
	seed.buildStyleSheetForLandscape(platform, &stylesheet)
	return stylesheet
}

func (seed Seed) BuildStyleSheetForPortrait(platform Platform) style.Sheet {
	var stylesheet = make(style.Sheet)
	seed.buildStyleSheetForPortrait(platform, &stylesheet)
	return stylesheet
}

func (seed Seed) buildStyleSheetForPortrait(platform Platform, sheet *style.Sheet) {
	if short := seed.ShortCircuit(platform); short.seed != nil {
		short.buildStyleSheetForPortrait(platform, sheet)
		return
	}

	//seed.postProduction()
	if data := seed.Portrait.Bytes(); data != nil {
		seed.styled = true
		if seed.template {
			sheet.Add("."+seed.id, seed.Portrait)
		} else {
			sheet.Add("#"+seed.id, seed.Portrait)
		}
	}
	for _, child := range seed.children {
		child.Root().buildStyleSheetForPortrait(platform, sheet)
	}
}

//Replace this seed with its desktop version.

func (seed Seed) HTML(platform Platform) []byte {
	if short := seed.ShortCircuit(platform); short.seed != nil {
		return short.HTML(platform)
	}
	if seed.template {

		for _, child := range seed.children {
			child.Root().Render(platform)
		}

		return nil
	}

	//seed.postProduction()

	var html bytes.Buffer

	if seed.page && !seed.splash {
		html.WriteString("<template id='")
		html.WriteString(fmt.Sprint(seed.id))
		html.WriteString(":template'>")
	}

	html.WriteByte('<')
	html.WriteString(seed.tag)
	html.WriteByte(' ')
	if seed.attr != "" {
		html.WriteString(seed.attr)
		html.WriteByte(' ')
	}

	for attribute, value := range seed.Element.Attributes {
		html.WriteString(string(attribute))
		html.WriteByte('=')
		html.WriteByte('\'')
		html.WriteString(value)
		html.WriteByte('\'')
	}

	html.WriteString("id='")
	html.WriteString(fmt.Sprint(seed.id))
	html.WriteByte('\'')

	if seed.class != "" {
		html.WriteString("class='")
		html.WriteString(seed.class)

		for tag := range seed.tags {
			html.WriteByte(' ')
			html.WriteString(tag)
		}

		html.WriteByte('\'')
	}

	if !seed.styled {
		if data := seed.Style.Bytes(); data != nil {
			html.WriteString(" style='")
			html.Write(data)
			html.WriteByte('\'')
		}
	}

	/*if seed.onclick != nil && seed.parent == nil {
		html.WriteString(" onclick='")
		html.Write(script.ToJavascript(seed.onclick))
		html.WriteByte('\'')
	}

	if seed.onchange != nil {
		html.WriteString(" onchange='")
		html.Write(script.ToJavascript(seed.onchange))
		html.WriteByte('\'')
	}*/

	html.WriteByte('>')

	if seed.content != nil {
		html.Write(seed.content)
	}

	if seed.Element.HTML != nil {
		html.Write(seed.Element.HTML)
	}

	for _, child := range seed.children {
		html.Write(child.Root().Render(platform))
	}

	switch seed.tag {
	case "input", "img", "br", "hr", "meta", "area", "base", "col", "embed", "link", "param", "source", "track", "wbr":

	default:
		html.WriteString("</")
		html.WriteString(seed.tag)
		html.WriteByte('>')
	}

	if seed.page && !seed.splash {
		var onready = seed.app.harvestOnReadyPage(seed)

		html.WriteString("<script>")
		html.Write(onready)
		html.WriteString("</script>")

		html.WriteString("</template>")
	}

	return html.Bytes()
}

func (seed Seed) Render(platform Platform) []byte {
	return seed.HTML(platform)
}

func (seed Seed) getScripts(platform Platform) []string {
	if short := seed.ShortCircuit(platform); short.seed != nil {
		return short.getScripts(platform)
	}

	if seed.template {
		return nil
	}

	var scripts = seed.scripts

	for _, child := range seed.children {
		scripts = append(scripts, child.Root().getScripts(platform)...)
	}

	return scripts
}

func (seed Seed) Scripts(platform Platform) map[string]struct{} {
	var scripts = seed.getScripts(platform)
	var uniques = make(map[string]struct{})

	for _, script := range scripts {
		uniques[script] = struct{}{}
	}

	return uniques
}

func (app App) Render(platform Platform) []byte {
	if !app.built {
		app.build()
	}
	return app.render(true, platform)
}

//Return a fully fully rendered application in HTML for the seed.
func (app App) render(production bool, platform Platform) []byte {
	if production {
		app.production = true
	}
	app.platform = platform

	return app.HTML()
}

var tail string

func Tail(t string) {
	tail += t
}
