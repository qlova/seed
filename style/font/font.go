package font

import (
	"encoding/base64"
	"image/color"
	"math/big"
	"strings"

	"github.com/qlova/seed"
	"github.com/qlova/seed/asset"
	"github.com/qlova/seed/css"
	"github.com/qlova/seed/style"
)

//Font is a type of font.
type Font struct {
	name, path string
	css.FontFace
}

//AddTo impliments seed.Option
func (f Font) AddTo(c seed.Seed) {
	var data data
	c.Read(&data)

	data.fonts = append(data.fonts, f)

	c.Write(data)

	css.SetFontFamily(f).And(asset.New(f.path)).AddTo(c)
}

//And impliments seed.Option
func (f Font) And(more ...seed.Option) seed.Option {
	return seed.And(f, more...)
}

type data struct {
	seed.Data

	fonts []Font
}

var seeds = make(map[seed.Seed]data)

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

//SetColor sets the color of the text.
func SetColor(c color.Color) seed.Option {
	return style.SetTextColor(c)
}

//SetSize sets the size of the text.
func SetSize(u style.Unit) css.Rule {
	return css.SetFontSize(u.Unit())
}
