package style

import "github.com/qlova/seed/style/css"
import "encoding/base64"

//A font, a style definition for how text should be rendered.
type Font struct {
	name string
	css.FontFace
}

var font_id int64 = 1;

//Create a new font based on the given font file path.
func NewFont(path string) Font {
	
	var id = base64.RawURLEncoding.EncodeToString(big.NewInt(font_id).Bytes())
	font_id++
	
	var font = Font{
		name: id,
		FontFace: css.NewFontFace(id, path),
	}

	//Avoid invisisible text while webfonts are loading.
	//font.FontFace.FontDisplay = css.Swap

	return font
}

//Set the font used by this element.
func (style Style) SetFont(font Font) {
	style.SetFontFamily(font.FontFace)
}
