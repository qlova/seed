package animation

import (
	"bytes"
	"fmt"
	"sort"

	"qlova.org/seed"
	"qlova.org/seed/use/css"
)

type harvester struct {
	animations map[int]Animation
}

func newHarvester() harvester {
	return harvester{
		make(map[int]Animation),
	}
}

func (h harvester) harvest(c seed.Seed) map[int]Animation {
	var data data
	c.Load(&data)

	for _, anim := range data.animations {
		h.animations[anim.id] = anim
	}

	for _, child := range c.Children() {
		h.harvest(child)
	}

	return h.animations
}

func init() {
	css.RegisterRenderer(func(c seed.Seed) []byte {
		var harvested = newHarvester().harvest(c)
		var b bytes.Buffer

		//Deterministic render.
		keys := make([]int, len(harvested))
		for i := range harvested {
			keys = append(keys, i)
		}
		sort.Ints(keys)

		for _, key := range keys {
			anim := harvested[key]

			fmt.Fprintf(&b, `@keyframes a%v {`, anim.id)

			//Deterministic render.
			keys := make([]float64, len(anim.frames))
			for i := range anim.frames {
				keys = append(keys, i)
			}
			sort.Float64s(keys)

			for _, key := range keys {
				frame := anim.frames[key]

				var x string = "var(--x, 0)"
				var y string = "var(--y, 0)"
				var angle string = "var(--angle, 0)"
				var scale string = "var(--scale, 1)"

				if key == 0 {
					fmt.Fprint(&b, "from {")
				} else if key == 100 {
					fmt.Fprint(&b, "to {")
				} else {
					fmt.Fprintf(&b, "%v%% {", key)
				}
				for _, rule := range frame.Rules() {
					switch rule.Property() {
					case "--x":
						x = rule.Value()
						continue
					case "--y":
						y = rule.Value()
						continue
					case "--scale":
						scale = rule.Value()
						continue
					case "--angle":
						angle = rule.Value()
						continue
					}
					fmt.Fprintf(&b, `%v`, rule)
				}
				fmt.Fprintf(&b, `transform: translate(%v, %v) rotate(%v) scale(%v);`, x, y, angle, scale)
				fmt.Fprint(&b, "}")
			}
			fmt.Fprint(&b, `}`)
		}

		return b.Bytes()
	})
}
