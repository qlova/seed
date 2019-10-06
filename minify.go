package seed

import (
	"regexp"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/html"
	"github.com/tdewolff/minify/js"
	"github.com/tdewolff/minify/json"
	"github.com/tdewolff/minify/svg"
	"github.com/tdewolff/minify/xml"
)

func mini(data []byte) ([]byte, error) {
	minifier := minify.New()
	minifier.Add("text/html", &html.Minifier{
		KeepDocumentTags: true,
	})
	minifier.AddFunc("text/css", css.Minify)
	minifier.AddFunc("image/svg+xml", svg.Minify)
	minifier.AddFuncRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), js.Minify)
	minifier.AddFuncRegexp(regexp.MustCompile("[/+]json$"), json.Minify)
	minifier.AddFuncRegexp(regexp.MustCompile("[/+]xml$"), xml.Minify)

	minified, err := minifier.Bytes("text/html", data)
	if err != nil {
		return nil, err
	}

	return minified, nil
}
