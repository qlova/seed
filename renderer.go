package seed

import (
	"bytes"
	"fmt"
)

//Platform is a platform type.
type Platform int

//List of potential platforms.
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

//ShortCircuit returns the actual seed for the given platform.
func (seed Seed) ShortCircuit(platform Platform) Seed {
	if platform == Desktop && seed.desktop.seed != nil {
		return seed.desktop
	}
	return Seed{}
}

func (seed Seed) buildStyleSheet(platform Platform, sheet *sheet) {
	if short := seed.ShortCircuit(platform); short.seed != nil {
		short.buildStyleSheet(platform, sheet)
		return
	}

	seed.styled = true
	var selector = "#" + seed.id
	if seed.Template {
		selector = "." + seed.id
	}

	sheet.AddSeed(selector, seed)

	for _, child := range seed.children {
		child.Root().buildStyleSheet(platform, sheet)
	}
}

//BuildStyleSheet builds the seeds stylesheet.
func (seed Seed) BuildStyleSheet(platform Platform) sheet {
	var stylesheet = newSheet()
	seed.buildStyleSheet(platform, &stylesheet)
	return stylesheet
}

//Render returns this seed rendered as HTML.
func (seed Seed) Render(platform Platform) []byte {
	if short := seed.ShortCircuit(platform); short.seed != nil {
		return short.Render(platform)
	}
	if seed.Template && !seed.TemplateRoot {
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

func (seed Seed) getScripts(platform Platform) []string {
	if short := seed.ShortCircuit(platform); short.seed != nil {
		return short.getScripts(platform)
	}

	if seed.Template {
		return nil
	}

	var scripts = seed.scripts

	for _, child := range seed.children {
		scripts = append(scripts, child.Root().getScripts(platform)...)
	}

	return scripts
}

//Scripts returns the scripts of this seed.
func (seed Seed) Scripts(platform Platform) map[string]struct{} {
	var scripts = seed.getScripts(platform)
	var uniques = make(map[string]struct{})

	for _, script := range scripts {
		uniques[script] = struct{}{}
	}

	return uniques
}

//Render renders the app to bytes.
func (app App) Render(platform Platform) []byte {
	app.build()
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
