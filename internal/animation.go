package internal

import "github.com/qlova/seed/style"
import "github.com/qlova/seed/style/css"

type Frame struct {
	style.Style
}
type Animation map[float64]func(Frame)

type AnimationRecord struct {
	Animation Animation
	ID        string
}

func (animation Animation) Bytes() []byte {
	var Keyframes = make(css.Keyframes, len(animation))
	for time := range animation {
		var frame = Frame{Style: style.New()}
		animation[time](frame)

		println(frame.Style.Style.Left().String())

		Keyframes[time] = frame.Style.Style
	}
	return Keyframes.Bytes()
}
