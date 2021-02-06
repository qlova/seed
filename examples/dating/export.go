package dating

import "qlova.org/seed/use/wasm"

func init() {
	wasm.Export(AddCustom)
	wasm.Export(GetHolidays)
	wasm.Export(GetCustom)
	wasm.Export(SaveCustom)
	wasm.Export(LoadCustom)
}
