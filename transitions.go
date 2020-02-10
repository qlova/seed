package seed

import (
	"math"

	"github.com/qlova/seed/script"
)

//Transition is a transition between pages.
type Transition struct {
	In  *Animation
	Out *Animation

	When     Page
	WhenTag  string
	WhenBack bool

	Then, Else *Transition
}

//Fade is a transition.
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

//FadeIn is a transition.
var FadeIn = Transition{
	In: Fade.In,
}

//FadeOut is a transition.
var FadeOut = Transition{
	In: Fade.Out,
}

//Flip is a transition.
var Flip = Transition{
	Out: &Animation{
		0: func(frame Frame) {
			frame.RotateX(0)
			frame.CSS().Set("transform-origin", "top")
		},

		100: func(frame Frame) {
			frame.RotateX(-math.Pi / 2)
			frame.CSS().Set("transform-origin", "top")
		},
	},
}

//Stay is a transition, where the page stays visisble until the end of the transition.
var Stay = Transition{
	Out: &Animation{
		0: func(frame Frame) {
			frame.SetLayer(0)
		},

		100: func(frame Frame) {
			frame.SetLayer(0)
		},
	},
	In: &Animation{
		0: func(frame Frame) {
			frame.SetLayer(0)
		},

		100: func(frame Frame) {
			frame.SetLayer(0)
		},
	},
}

//FlipOut is a transition.
var FlipOut = Transition{
	Out: Flip.Out,
}

//SlideIn stores all slidein animations.
var SlideIn = struct {
	Up, Down, Left, Right *Animation
}{
	Up:    SlideUp.In,
	Down:  SlideDown.In,
	Left:  SlideLeft.In,
	Right: SlideRight.In,
}

//SlideOut stores all slideout animations.
var SlideOut = struct {
	Up, Down, Left, Right *Animation
}{
	Up:    SlideDown.Out,
	Down:  SlideUp.Out,
	Left:  SlideRight.Out,
	Right: SlideLeft.Out,
}

//SlideUp is a transition.
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

//SlideDown is a transition.
var SlideDown = Transition{
	In: &Animation{
		0: func(frame Frame) {
			frame.Translate(0, -100)
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
			frame.Translate(0, -100)
		},
	},
}

//SlideLeft is a transition.
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

//SlideRight is a transition.
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
	if (element.classList.contains("page")) {
		let last=last_page; if (!last || last == loading_page) return;
	}

	set(element, "animation-name", animation);
	set(element, "animation-direction", "normal");
	set(element, "animation-fill-mode", "forwards");
	set(element, "animation-duration", duration+"s");
	set(element, "animation-iteration-count", 1);

	set(element, "z-index", "50");
	animating = true;
	setTimeout(function() {
		set(element, "animation", ""); 
		set(element, "z-index", "");
		animation_complete();
	}, 1000*duration);
}
function beginOutTransition(element, animation, duration) {
	if (element.classList.contains("page")) {
		let last=last_page; if (!last || last == loading_page) return;
	}

	set(element, "animation-name", animation);
	set(element, "animation-direction", "normal");
	set(element, "animation-fill-mode", "forwards");
	set(element, "animation-duration", duration+"s");
	set(element, "animation-iteration-count", 1);
	set(element, "z-index", "50");
	animating = true;
	
	goto_exitpromise = Promise.resolve().delay(1000*duration).then(function() {
		set(element, "animation", ""); 
		set(element, "z-index", "");
		animation_complete(); 
	});
}
`

const duration = "0.5"

func SetTransitionIn(Page script.Seed, trans Transition) {
	var q = Page.Q

	if trans.WhenBack {
		q.If(q.Value("going_back").Bool(), func() {
			if trans.Then != nil {
				SetTransitionIn(Page, *trans.Then)
			}
		}, q.Else(func() {
			if trans.Else != nil {
				SetTransitionIn(Page, *trans.Else)
			}
		}))
		return
	}

	if !trans.When.Null() {
		q.If(q.LastPage().Equals(trans.When.Ctx(q)), func() {
			if trans.Then != nil {
				SetTransitionIn(Page, *trans.Then)
			}
		}, q.Else(func() {
			if trans.Else != nil {
				SetTransitionIn(Page, *trans.Else)
			}
		}))
		return
	}

	if trans.WhenTag != "" {
		q.Javascript(`if (` + q.LastPage().Element() + ` == null) return;`)
		q.If(q.Value(q.LastPage().Element()+".classList.contains('"+trans.WhenTag+"')").Bool(), func() {
			if trans.Then != nil {
				SetTransitionIn(Page, *trans.Then)
			}
		}, q.Else(func() {
			if trans.Else != nil {
				SetTransitionIn(Page, *trans.Else)
			}
		}))
		return
	}

	if trans.In != nil {
		q.Require(beginTransition)
		q.Javascript(`beginInTransition(` + Page.Element() + `, '` + q.Context.Animation(trans.In) + `', ` + duration + `);`)
	}
}

func SetTransitionOut(Page script.Seed, trans Transition) {
	var q = Page.Q

	if trans.WhenBack {
		q.If(q.Value("going_back").Bool(), func() {
			if trans.Then != nil {
				SetTransitionOut(Page, *trans.Then)
			}
		}, q.Else(func() {
			if trans.Else != nil {
				SetTransitionOut(Page, *trans.Else)
			}
		}))
		return
	}

	if !trans.When.Null() {
		q.If(q.NextPage().Equals(trans.When.Ctx(q)), func() {
			if trans.Then != nil {
				SetTransitionOut(Page, *trans.Then)
			}
		}, q.Else(func() {
			if trans.Else != nil {
				SetTransitionOut(Page, *trans.Else)
			}
		}))
		return
	}

	if trans.WhenTag != "" {
		q.If(q.Value(q.LastPage().Element()+".classList.contains('"+trans.WhenTag+"')").Bool(), func() {
			if trans.Then != nil {
				SetTransitionOut(Page, *trans.Then)
			}
		}, q.Else(func() {
			if trans.Else != nil {
				SetTransitionOut(Page, *trans.Else)
			}
		}))
		return
	}

	if trans.Out != nil {
		q.Require(beginTransition)
		q.Javascript(`beginOutTransition(` + Page.Element() + `, '` + q.Context.Animation(trans.Out) + `', ` + duration + `);`)
	}
}

//SetTransition sets a page transition for the page.
func (page Page) SetTransition(trans Transition) {
	if trans.In != nil || !trans.When.Null() || trans.WhenTag != "" || trans.WhenBack {
		page.OnPageEnter(func(q script.Ctx) {
			var Page = page.Ctx(q)
			SetTransitionIn(Page.Seed, trans)
		})
	}
	if trans.Out != nil || !trans.When.Null() || trans.WhenTag != "" || trans.WhenBack {
		page.OnPageExit(func(q script.Ctx) {
			var Page = page.Ctx(q)
			SetTransitionOut(Page.Seed, trans)
		})
	}
}
