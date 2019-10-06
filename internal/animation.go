package internal

import (
	"github.com/qlova/seed/style"
	"github.com/qlova/seed/style/css"
)

//Frame is an animation frame.
type Frame struct {
	style.Style
}

//Animation is a mapping of float64 to frames as a ratio between 0 and 1 (start and end).
type Animation map[float64]func(Frame)

//Bytes returns the CSS keyframes of the animation.
func (animation Animation) Bytes() []byte {
	var Keyframes = make(css.Keyframes, len(animation))
	for time := range animation {
		var frame = Frame{Style: style.New()}
		animation[time](frame)
		Keyframes[time] = frame.Style.Style
	}
	return Keyframes.Bytes()
}
