package form

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/css"
	"github.com/qlova/seed/html/attr"
	"github.com/qlova/seed/s/html/form"
	"github.com/qlova/seed/script"
)

//New returns a new HTML form element.
func New(options ...seed.Option) seed.Seed {
	return form.New(
		css.SetDisplay(css.Flex),
		css.SetFlexDirection(css.Column),
		seed.Options(options),
	)
}

//ReportValidity reports the validity of the form.
func ReportValidity(form seed.Seed) script.Bool {
	return script.Bool{Value: script.Element(form).Call("reportValidity")}
}

//SetRequired sets the input element to be required for the form.
func SetRequired() seed.Option {
	return attr.Set("required", "")
}
