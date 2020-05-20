package spinner

import (
	"fmt"
	"image/color"

	"github.com/qlova/seed"
	"github.com/qlova/seed/css"
	"github.com/qlova/seed/html"
	"github.com/qlova/seed/s/html/div"
)

func Ellipsis(c color.Color, options ...seed.Option) seed.Seed {
	r, g, b, _ := c.RGBA()
	rgb := fmt.Sprint("rgb(", r, ",", g, ",", b, ")")

	return div.New(seed.Options(options),

		html.AddClass(`lds-ellipsis`),
		html.Set(`<div style="background-color: `+rgb+`"></div><div style="background-color: `+rgb+`"></div><div style="background-color: `+rgb+`"></div>`),

		css.Add(`
			.lds-ellipsis {
				display: inline-block;
				position: relative;
				width: 80px;
				height: 80px;
			  }
			  .lds-ellipsis div {
				position: absolute;
				top: 33px;
				width: 13px;
				height: 13px;
				border-radius: 50%;
				background-color: inherit;
				animation-timing-function: cubic-bezier(0, 1, 1, 0);
			  }
			  .lds-ellipsis div:nth-child(1) {
				left: 8px;
				animation: lds-ellipsis1 0.6s infinite;
			  }
			  .lds-ellipsis div:nth-child(2) {
				left: 8px;
				animation: lds-ellipsis2 0.6s infinite;
			  }
			  .lds-ellipsis div:nth-child(3) {
				left: 32px;
				animation: lds-ellipsis2 0.6s infinite;
			  }
			  .lds-ellipsis div:nth-child(4) {
				left: 56px;
				animation: lds-ellipsis3 0.6s infinite;
			  }
			  @keyframes lds-ellipsis1 {
				0% {
				  transform: scale(0);
				}
				100% {
				  transform: scale(1);
				}
			  }
			  @keyframes lds-ellipsis3 {
				0% {
				  transform: scale(1);
				}
				100% {
				  transform: scale(0);
				}
			  }
			  @keyframes lds-ellipsis2 {
				0% {
				  transform: translate(0, 0);
				}
				100% {
				  transform: translate(24px, 0);
				}
			  }
		`),
	)
}
