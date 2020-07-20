package js

import (
	"bufio"
	"bytes"
	"io"
	"strconv"
	"text/scanner"

	"qlova.org/seed"
	"qlova.org/seed/asset/assets"
)

type macroWriter struct {
	w *bufio.Writer

	seeds []seed.Seed
}

func newMacroWriter(w io.Writer, seeds ...seed.Seed) macroWriter {
	return macroWriter{bufio.NewWriter(w), seeds}
}

func (m macroWriter) Flush() {
	m.w.Flush()
}

func (m macroWriter) Write(data []byte) (int, error) {
	for i := 0; i < len(data); i++ {
		b := data[i]

		if i+4 < len(data) && bytes.Equal(data[i:i+4], []byte(`"'#(`)) {

			//detect end
			var end int
			for end = i; data[end] != ')'; end++ {
			}

			var macro = data[i+4 : end]

			//skip to end
			i = end

			var scan scanner.Scanner
			var reader = bytes.NewReader(macro)
			scan.Init(reader)

			scan.Scan()
			switch scan.TokenText() {
			case "import":
				scan.Scan()
				var s = scan.TokenText()

				var path, err = strconv.Unquote(s)
				if err != nil {
					panic("import macro error: " + err.Error())
				}

				for _, c := range m.seeds {
					Require(path, imports[path]).AddTo(c)
					assets.New(path).AddTo(c)
				}
			default:
				panic("unknown macro command: " + scan.TokenText())
			}
		} else {
			m.w.WriteByte(b)
		}
	}
	return len(data), m.w.Flush()
}
