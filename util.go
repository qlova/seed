package seed

import "math"

//Deg returns degrees to radians.
func Deg(deg float64) float64 {
	return deg * math.Pi / 180
}
