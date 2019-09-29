package style

import "math/big"
import "github.com/qlova/seed/style/css"
import "encoding/base64"

//Font is a style definition for how text should be rendered.
type Font struct {
	name string
	css.FontFace
}

var fontID int64 = 1

//NewFont creates a new font based on the given font file path.
func NewFont(path string) Font {

	var id = base64.RawURLEncoding.EncodeToString(big.NewInt(fontID).Bytes())
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
	style.SetFontFamily(font.FontFace)
}
