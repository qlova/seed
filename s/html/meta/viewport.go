package meta

import (
	"fmt"
	"strings"

	"github.com/qlova/seed"
	"github.com/qlova/seed/html/attr"
)

//DeviceWidth can be used inside Viewport.Width.
const DeviceWidth = -1

//DeviceHeight can be used inside Viewport.Height.
const DeviceHeight = -1

//Viewport is an HTML meta viewport specification.
type Viewport struct {
	Width, Height float32

	InitialScale, MinimumScale, MaximumScale float32

	UserScalable *bool
}

func (v Viewport) render() string {
	var builder strings.Builder

	if v.Width < 0 {
		builder.WriteString("width=device-width,")
	} else if v.Height > 0 {
		fmt.Fprintf(&builder, "width=%v,", v.Width)
	}

	if v.Height < 0 {
		builder.WriteString("height=device-height,")
	} else if v.Height > 0 {
		fmt.Fprintf(&builder, "height=%v,", v.Height)
	}

	if v.InitialScale > 0 {
		fmt.Fprintf(&builder, "initial-scale=%v,", v.InitialScale)
	}

	if v.MinimumScale > 0 {
		fmt.Fprintf(&builder, "minimum-scale=%v,", v.MinimumScale)
	}

	if v.MaximumScale > 0 {
		fmt.Fprintf(&builder, "minimum-scale=%v,", v.MaximumScale)
	}

	if v.UserScalable != nil {
		if !*v.UserScalable {
			fmt.Fprint(&builder, "user-scalable=no,")
		}
	}

	var result = builder.String()

	return result[:len(result)-1]
}

//Option returns the viewport's option.
func (v Viewport) option() seed.Option {
	return attr.Set("name", "viewport").And(attr.Set("content", v.render()))
}

//AddTo implements seed.Option.AddTo
func (v Viewport) AddTo(any seed.Any) {
	v.option().AddTo(any)
}

//Apply implements seed.Option.Apply
func (v Viewport) Apply(any seed.Ctx) {
	v.option().Apply(any)
}

//Reset implements seed.Option.Apply
func (v Viewport) Reset(any seed.Ctx) {
	v.option().Reset(any)
}

//And implements seed.Option.And
func (v Viewport) And(options ...seed.Option) seed.Option {
	return v.option().And(options...)
}
