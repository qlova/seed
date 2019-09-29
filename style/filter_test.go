package style

import (
	"fmt"
	"testing"
	"time"
)

func TestColors(t *testing.T) {

	var r, g, b float64 = 255, 0, 0

	var color = newRGB(r, g, b)
	var solver = newSolver(color)

	start := time.Now()
	var _, loss, filter = solver.Solve()
	elapsed := time.Since(start)
	fmt.Println("Binomial took ", elapsed)
	fmt.Println(loss)

	var lossMessage string
	if loss < 1 {
		lossMessage = "This is a perfect result."
	} else if loss < 5 {
		lossMessage = "The is close enough."
	} else if loss < 15 {
		lossMessage = "The color is somewhat off. Consider running it again."
	} else {
		lossMessage = "The color is extremely off. Run it again!"
	}

	fmt.Println(filter)
	fmt.Println(lossMessage)
}
