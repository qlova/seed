package assets

import (
	"strings"

	"qlova.org/seed"
	"qlova.org/seed/assets/inbed"
	"qlova.org/seed/client"
	"qlova.org/seed/client/clientfmt"
	"qlova.org/seed/client/clientside"
	"qlova.org/seed/use/js"
)

func init() {
	inbed.PackageName = "assets"
	inbed.ImporterName = "assets.go"
}

func init() {
	client.RegisterRenderer(func(c seed.Seed) []byte {
		return []byte(`
			seed.asset = function(src) {
				if (!src.startsWith("/") && !src.startsWith("http")) {
					return "/assets/"+src;
				}
				return src;
			};
		`)
	})
}

//Path returns the correct path from the given base path.
func Path(src string) string {
	//New(src)
	if !strings.HasPrefix(src, "/") && !strings.HasPrefix(src, "http") {
		return "/assets/" + src
	}
	return src
}

//PathOf returns the correct path from the given base path.
func PathOf(src client.String) client.String {
	switch p := src.(type) {
	case clientfmt.String:
		return clientfmt.NewString(js.String{Value: js.Func("seed.asset").Call(p)}, p)

	case *clientside.String:
		return clientfmt.NewString(js.String{Value: js.Func("seed.asset").Call(p)}, p)

	case client.String:
		return js.String{Value: js.Func("seed.asset").Call(p)}
	}
	return nil
}
