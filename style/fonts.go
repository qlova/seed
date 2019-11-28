package style

import (
	"encoding/base64"
	"math/big"
	"strings"

	"github.com/qlova/seed/style/css"
)

//Font is a style definition for how text should be rendered.
type Font struct {
	name string
	css.FontFace
}

var fontID int64 = 1

//NewFont creates a new font based on the given font file path.
func NewFont(path string) Font {

	var id = base64.RawURLEncoding.EncodeToString(big.NewInt(fontID).Bytes())

	if id[0] >= '0' && id[0] <= '9' {
		id = "_" + id
	}

	id = strings.Replace(id, "-", "__", -1)

	fontID++

	var font = Font{
		name:     id,
		FontFace: css.NewFontFace(id, path),
	}

	//Avoid invisisible text while webfonts are loading.
	//font.FontFace.FontDisplay = css.Swap

	return font
}

//SetFont sets the font used by this element.
func (style Style) SetFont(font Font) {
	style.CSS().SetFontFamily(font.FontFace)
}
