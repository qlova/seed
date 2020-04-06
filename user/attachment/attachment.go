package attachment

import (
	"encoding/base64"
	"fmt"
	"math/big"

	"github.com/qlova/seed/js"
	"github.com/qlova/seed/script"
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

func (a Attachment) GetValue() script.Value {
	return js.NewValue(`window.attachments[%v]`, js.NewString(a.string))
}

var id int64

func New() Attachment {
	id++
	return Attachment{base64.RawURLEncoding.EncodeToString(big.NewInt(int64(id)).Bytes())}
}

type File struct{}
