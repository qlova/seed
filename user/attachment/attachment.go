package attachment

import (
	"encoding/base64"
	"fmt"
	"math/big"

	"qlova.org/seed/js"
)

type Attachment struct {
	string
}

func (a Attachment) Set(v js.Value) func(js.Ctx) {
	return func(q js.Ctx) {
		q(`if (!window.attachments) window.attachments = {};`)
		q(fmt.Sprintf(`window.attachments["%v"] = %v;`, a.string, v))
	}
}

func (a Attachment) GetFile() js.Value {
	return js.NewValue(`window.attachments[%v]`, js.NewString(a.string))
}

func (a Attachment) GetValue() js.Value {
	return a.GetFile()
}

func (a Attachment) GetBool() js.Bool {
	return a.GetValue().GetBool()
}

var id int64

func New() Attachment {
	id++
	return Attachment{base64.RawURLEncoding.EncodeToString(big.NewInt(int64(id)).Bytes())}
}

type File struct{}
