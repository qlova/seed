package main

import (
	"fmt"

	"github.com/qlova/seed/style/css"
)

func main() {
	var style = css.NewStyle()
	style.SetWidth(css.Number(100).Vw())
	style.SetHeight(css.Number(100).Vh())
	style.SetFloat(css.Left)

	fmt.Println(style.Float())

	fmt.Println(style)
}
