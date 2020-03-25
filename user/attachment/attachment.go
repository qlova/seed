package attachment

import (
	"encoding/base64"
	"math/big"

	"github.com/qlova/seed/script"
)

type Attachment struct {
	string
}

func (a Attachment) Set(v script.Value) func(script.Ctx) {
	return func(q script.Ctx) {
		q.Javascript(`if (!window.attachments) window.attachments = {};`)
		q.Javascript(`window.attachments["%v"] = %v;`, a.string, q.Raw(v))
	}
}

func (a Attachment) ValueFromCtx(q script.AnyCtx) script.Value {
	return script.CtxFrom(q).Value(`window.attachments["%v"]`, a.string).File()
}

var id int64

func New() Attachment {
	id++
	return Attachment{base64.RawURLEncoding.EncodeToString(big.NewInt(int64(id)).Bytes())}
}

type File struct{}
