package script

//SetAppColor sets the color of the app theme dynamically.
func (q Ctx) SetAppColor(c Color) {
	q.Javascript(`document.querySelector("meta[name=theme-color]").setAttribute("content", %v);`, c)
}
