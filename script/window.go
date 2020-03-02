package script

//ResizeWindowTo resizes the current window if possible to the requested dimensions.
func (q Ctx) ResizeWindowTo(w, h Int) {
	q.js.Run("window.resizeTo", w, h)
}
