package css

import (
	"bytes"
	"fmt"
)

const From = 0
const To = 100

type Keyframes map[float64]Style

func (k Keyframes) Bytes() []byte {
	var buffer bytes.Buffer

	var keys = make([]float64, len(k))
	var i = 0
	for key := range k {
		keys[i] = key
		i++
	}

	for _, key := range keys {
		switch key {
		case From:
			buffer.WriteString(`from {`)
		case To:
			buffer.WriteString(`to {`)
		default:
			buffer.WriteString(fmt.Sprint(key, `% {`))
		}

		buffer.Write(k[key].Bytes())
		buffer.WriteByte('}')
	}

	return buffer.Bytes()
}
