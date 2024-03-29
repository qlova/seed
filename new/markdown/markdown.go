package markdown

import (
	"regexp"

	"qlova.org/seed"
	"qlova.org/seed/new/html/div"
	"qlova.org/seed/use/html"

	"github.com/gomarkdown/markdown"
	htmlmd "github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/microcosm-cc/bluemonday"
)

var policy = bluemonday.UGCPolicy()

func init() {
	policy.AllowAttrs("style").OnElements("span", "p")
	policy.AllowStyles("color").Matching(regexp.MustCompile("(?i)^#([0-9a-f]{3,4}|[0-9a-f]{6}|[0-9a-f]{8})$")).Globally()
}

//New returns a new markdown container.
func New(options ...seed.Option) seed.Seed {
	return div.New(
		seed.Options(options),
	)
}

//Set sets the inner HTML of the seed to rendered and sanitized markdown.
func Set(md string) seed.Option {

	extensions := parser.Tables
	p := parser.NewWithExtensions(extensions)

	rendered := markdown.ToHTML([]byte(md), p, htmlmd.NewRenderer(htmlmd.RendererOptions{
		Flags: htmlmd.CommonFlags,
	}))
	rendered = policy.SanitizeBytes(rendered)
	return html.Set(string(rendered))
}
