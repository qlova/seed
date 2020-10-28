package font

import (
	"encoding/base64"
	"math/big"
	"strings"

	"qlova.org/seed"
	"qlova.org/seed/asset"
	"qlova.org/seed/asset/assets"
	"qlova.org/seed/css"
)

//Font is a type of font.
type Font struct {
	name, path string
	css.FontFace
}

//AddTo impliments seed.Option
func (f Font) AddTo(c seed.Seed) {
	var data data
	c.Load(&data)

	data.fonts = append(data.fonts, f)

	c.Save(data)

	css.SetFontFamily(f).And(assets.New(f.path)).AddTo(c)
}

//And impliments seed.Option
func (f Font) And(more ...seed.Option) seed.Option {
	return seed.And(f, more...)
}

type data struct {
	fonts []Font
}

var id int64

//New returns a new font.
func New(path string) Font {
	path = asset.Path(path)

	id++

	var name = base64.RawURLEncoding.EncodeToString(big.NewInt(id).Bytes())

	if name[0] >= '0' && name[0] <= '9' {
		name = "_" + name
	}

	name = strings.Replace(name, "-", "__", -1)

	return Font{
		name:     name,
		path:     path,
		FontFace: css.NewFontFace(name, path),
	}
}
