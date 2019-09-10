package seed

import (
	"math"

	"github.com/qlova/seed/script"
)

type Transition struct {
	In  *Animation
	Out *Animation

	When    Page
	WhenTag string

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

var beginTransition = `function beginInTransition(element, animation, duration) {
	let last=last_page; if (!last || last == loading_page) return;

	set(element, "animation-name", animation);
	set(element, "animation-direction", "normal");
	set(element, "animation-fill-mode", "forwards");
	set(element, "animation-duration", duration);
	set(element, "animation-iteration-count", 1);

	set(element, "z-index", "50");
	animating = true;
	setTimeout(function() {
		set(element, "animation", ""); 
		set(element, "z-index", "");
		animation_complete();
	}, 500);
}
function beginOutTransition(element, animation, duration) {
	let last=last_page; if (!last || last == loading_page) return;

	set(element, "animation-name", animation);
	set(element, "animation-direction", "normal");
	set(element, "animation-fill-mode", "forwards");
	set(element, "animation-duration", duration);
	set(element, "animation-iteration-count", 1);
	
	goto_exitpromise = Promise.resolve().delay(500).then(function() {
		set(element, "animation", ""); 
		set(element, "z-index", "");
		animation_complete(); 
	});
}
`

func setTransitionIn(Page script.Page, trans Transition) {
	var q = Page.Q

	if !trans.When.Null() {
		q.If(q.LastPage().Equals(trans.When.Script(q)), func() {
			if trans.Then != nil {
				setTransitionIn(Page, *trans.Then)
			}
			q.Return()
		})
		if trans.Else != nil {
			setTransitionIn(Page, *trans.Else)
			q.Return()
			return
		}
		return
	}

	if trans.WhenTag != "" {
		q.Javascript(`if (` + q.LastPage().Element() + ` == null) return;`)
		q.If(q.Value(q.LastPage().Element()+".classList.contains('"+trans.WhenTag+"')").Bool(), func() {
			if trans.Then != nil {
				setTransitionIn(Page, *trans.Then)
			}
			q.Return()
		})
		if trans.Else != nil {
			setTransitionIn(Page, *trans.Else)
			q.Return()
			return
		}
		return
	}

	if trans.In != nil {
		q.Require(beginTransition)
		q.Javascript(`beginInTransition(` + Page.Element() + `, '` + q.Context.Animation(trans.In) + `', '0.5s');`)
	}
}

func setTransitionOut(Page script.Page, trans Transition) {
	var q = Page.Q

	if !trans.When.Null() {
		q.If(q.NextPage().Equals(trans.When.Script(q)), func() {
			if trans.Then != nil {
				setTransitionOut(Page, *trans.Then)
			}
			q.Return()
		})
		if trans.Else != nil {
			setTransitionOut(Page, *trans.Else)
			q.Return()
			return
		}
	}

	if trans.WhenTag != "" {
		q.If(q.Value(q.LastPage().Element()+".classList.contains('"+trans.WhenTag+"')").Bool(), func() {
			if trans.Then != nil {
				setTransitionOut(Page, *trans.Then)
			}
			q.Return()
		})
		if trans.Else != nil {
			setTransitionOut(Page, *trans.Else)
			q.Return()
			return
		}
	}

	if trans.Out != nil {
		q.Require(beginTransition)
		q.Javascript(`beginOutTransition(` + Page.Element() + `, '` + q.Context.Animation(trans.Out) + `', '0.5s');`)
	}
}

func (page Page) SetTransition(trans Transition) {
	if trans.In != nil || !trans.When.Null() || trans.WhenTag != "" {
		page.OnPageEnter(func(q Script) {
			var Page = page.Script(q)
			setTransitionIn(Page, trans)
		})
	}
	if trans.Out != nil || !trans.When.Null() || trans.WhenTag != "" {
		page.OnPageExit(func(q Script) {
			var Page = page.Script(q)
			setTransitionOut(Page, trans)
		})
	}
}
