package seed

import "math"

type Transition struct {
	In  *Animation
	Out *Animation

	Mirror bool
}

var Fade = Transition{
	Mirror: false,
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

func (page Page) SetTransition(trans Transition) {
	if trans.In != nil {
		page.OnPageEnter(func(q Script) {
			var Page = page.Script(q)

			Page.SetAnimation(trans.In)
			Page.SetAnimationDuration(q.Float(0.5))
			Page.SetAnimationIterations(q.Int(1))

			Page.Javascript(`let last=last_page; if (!last) return;`)
			Page.Javascript(`set(get(last), "display", "inline-flex");`)
			Page.Javascript(`set(` + Page.Element() + `, "z-index", "50");`)
			Page.Javascript(`animating = true;`)
			Page.Javascript(`setTimeout(function() { set(get(last), "display", "none"); set(` + Page.Element() + `, "z-index", ""); animation_complete(); }, 500);`)
		})
	}
	if trans.Mirror && trans.In != nil {
		page.OnPageExit(func(q Script) {
			var Page = page.Script(q)

			Page.SetAnimation(trans.In)
			Page.SetAnimationReverse()
			Page.SetAnimationDuration(q.Float(0.5))
			Page.SetAnimationIterations(q.Int(1))

			Page.Javascript(`set(` + Page.Element() + `, "display", "inline-flex");`)
			Page.Javascript(`set(` + Page.Element() + `, "z-index", "50");`)
			Page.Javascript(`animating = true;`)
			Page.Javascript(`setTimeout(function() { set(` + Page.Element() + `, "display", "none"); set(` + Page.Element() + `, "animation", ""); set(` + Page.Element() + `, "z-index", "");animation_complete(); }, 500);`)
		})
	} else if trans.Out != nil {
		page.OnPageExit(func(q Script) {
			var Page = page.Script(q)

			Page.SetAnimation(trans.Out)
			Page.SetAnimationDuration(q.Float(0.5))
			Page.SetAnimationIterations(q.Int(1))

			Page.Javascript(`set(` + Page.Element() + `, "display", "inline-flex");`)
			Page.Javascript(`set(` + Page.Element() + `, "z-index", "50");`)
			Page.Javascript(`animating = true;`)
			Page.Javascript(`setTimeout(function() { set(` + Page.Element() + `, "display", "none"); set(` + Page.Element() + `, "animation", ""); set(` + Page.Element() + `, "z-index", "");animation_complete(); }, 500);`)
		})
	}
}
