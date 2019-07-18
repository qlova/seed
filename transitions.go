package seed

import "github.com/qlova/seed/script"

import "math"

type Transition struct {
	In  *Animation
	Out *Animation

	When       Page
	Then, Else *Transition
}

var Fade = Transition{
	In: &Animation{
		0: func(frame Frame) {
			frame.SetOpacity(0)
		},

		100: func(frame Frame) {
			frame.SetOpacity(1)
		},
	},
	Out: &Animation{
		0: func(frame Frame) {
			frame.SetOpacity(1)
		},

		100: func(frame Frame) {
			frame.SetOpacity(0)
		},
	},
}

var FadeIn = Transition{
	In: Fade.In,
}

var FadeOut = Transition{
	In: Fade.Out,
}

var Flip = Transition{
	Out: &Animation{
		0: func(frame Frame) {
			frame.RotateX(0)
			frame.Set("transform-origin", "top")
		},

		100: func(frame Frame) {
			frame.RotateX(-math.Pi / 2)
			frame.Set("transform-origin", "top")
		},
	},
}

var FlipOut = Transition{
	Out: Flip.Out,
}

var SlideUp = Transition{
	In: &Animation{
		0: func(frame Frame) {
			frame.Translate(0, 100)
		},

		100: func(frame Frame) {
			frame.Translate(0, 0)
		},
	},
	Out: &Animation{
		0: func(frame Frame) {
			frame.Translate(0, 0)
		},

		100: func(frame Frame) {
			frame.Translate(0, 100)
		},
	},
}

var SlideLeft = Transition{
	In: &Animation{
		0: func(frame Frame) {
			frame.Translate(100, 0)
		},

		100: func(frame Frame) {
			frame.Translate(0, 0)
		},
	},
	Out: &Animation{
		0: func(frame Frame) {
			frame.Translate(0, 0)
		},

		100: func(frame Frame) {
			frame.Translate(100, 0)
		},
	},
}

var SlideRight = Transition{
	In: &Animation{
		0: func(frame Frame) {
			frame.Translate(-100, 0)
		},

		100: func(frame Frame) {
			frame.Translate(0, 0)
		},
	},
	Out: &Animation{
		0: func(frame Frame) {
			frame.Translate(0, 0)
		},

		100: func(frame Frame) {
			frame.Translate(-100, 0)
		},
	},
}

func setTransitionIn(Page script.Page, trans Transition) {
	var q = Page.Q

	if !trans.When.Null() {
		q.If(q.LastPage().Equals(trans.When.Script(q)), func() {
			setTransitionIn(Page, *trans.Then)
			q.Return()
		})
		if trans.Else != nil {
			setTransitionIn(Page, *trans.Else)
			q.Return()
			return
		}
	}

	if trans.In != nil {
		Page.SetAnimation(trans.In)
		Page.SetAnimationDuration(q.Float(0.5))
		Page.SetAnimationIterations(q.Int(1))

		Page.Javascript(`set(get(last), "display", "inline-flex");`)
		Page.Javascript(`set(` + Page.Element() + `, "z-index", "50");`)
		Page.Javascript(`animating = true;`)
		Page.Javascript(`setTimeout(function() { set(get(last), "display", "none"); set(` + Page.Element() + `, "z-index", ""); animation_complete(); }, 500);`)
	}
}

func setTransitionOut(Page script.Page, trans Transition) {
	var q = Page.Q

	if !trans.When.Null() {
		q.If(q.NextPage().Equals(trans.When.Script(q)), func() {
			setTransitionOut(Page, *trans.Then)
			q.Return()
		})
		if trans.Else != nil {
			setTransitionOut(Page, *trans.Else)
			q.Return()
			return
		}
	}

	if trans.Out != nil {
		Page.SetAnimation(trans.Out)
		Page.SetAnimationDuration(q.Float(0.5))
		Page.SetAnimationIterations(q.Int(1))

		Page.Javascript(`set(` + Page.Element() + `, "display", "inline-flex");`)
		Page.Javascript(`set(` + Page.Element() + `, "z-index", "50");`)
		Page.Javascript(`animating = true;`)
		Page.Javascript(`setTimeout(function() { set(` + Page.Element() + `, "display", "none"); set(` + Page.Element() + `, "animation", ""); set(` + Page.Element() + `, "z-index", "");animation_complete(); }, 500);`)
	}
}

func (page Page) SetTransition(trans Transition) {
	if trans.In != nil || !trans.When.Null() {
		page.OnPageEnter(func(q Script) {
			var Page = page.Script(q)
			Page.Javascript(`let last=last_page; if (!last) return;`)
			setTransitionIn(Page, trans)
		})
	}
	if trans.Out != nil || !trans.When.Null() {
		page.OnPageExit(func(q Script) {
			var Page = page.Script(q)
			setTransitionOut(Page, trans)

		})
	}
}
