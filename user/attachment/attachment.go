package attachment

import (
	"encoding/base64"
	"fmt"
	"math/big"

	"qlova.org/seed/js"
	"qlova.org/seed/script"
)

type Attachment struct {
	string
}

func (a Attachment) Set(v script.Value) func(script.Ctx) {
	return func(q script.Ctx) {
		q(`if (!window.attachments) window.attachments = {};`)
		q(fmt.Sprintf(`window.attachments["%v"] = %v;`, a.string, v))
	}
}

func (a Attachment) GetFile() script.File {
	return script.File{js.NewValue(`window.attachments[%v]`, js.NewString(a.string))}
}

func (a Attachment) GetValue() script.Value {
	return a.GetFile().Value
}

func (a Attachment) GetBool() script.Bool {
	return a.GetValue().GetBool()
}

var id int64

func New() Attachment {
	id++
	return Attachment{base64.RawURLEncoding.EncodeToString(big.NewInt(int64(id)).Bytes())}
}

type File struct{}
