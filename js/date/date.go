package date

import "github.com/qlova/seed/js"

func Now() js.Number {
	return js.Number{js.NewValue(`Date.now()`)}
}
