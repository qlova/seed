package asset

import (
	"reflect"
	"strings"

	"qlova.org/seed"
	"qlova.org/seed/asset/assets"
	"qlova.org/seed/client"
	"qlova.org/seed/client/clientfmt"
	"qlova.org/seed/client/clientside"
	"qlova.org/seed/js"
	"qlova.org/seed/script"
	"qlova.org/seed/sum"
)

func init() {
	script.RegisterRenderer(func(c seed.Seed) []byte {
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
func Path(src sum.String) sum.String {
	switch p := src.(type) {
	case string:
		assets.New(p)
		if !strings.HasPrefix(p, "/") && !strings.HasPrefix(p, "http") {
			return "/assets/" + p
		}
		return p

	case *clientside.String:
		return clientfmt.NewString(js.String{Value: js.Func("seed.asset").Call(p)}, p)

	case client.Value:
		return js.String{Value: js.Func("seed.asset").Call(p)}

	default:
		panic("asset.Path: invalid argument " + reflect.TypeOf(src).String())
	}
}
