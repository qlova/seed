package date

import "qlova.org/seed/js"

func Now() js.Number {
	return js.Number{js.NewValue(`Date.now()`)}
}
