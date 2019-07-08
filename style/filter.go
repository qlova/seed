package style

import "fmt"
import "math"
import "math/rand"
import "time"

var seed = 1.0;
func random() float64 {
	seed++
    var x = math.Sin(seed) * 10000;
    return x - math.Floor(x);
}

type matrix [9]float64

type Color struct {
	R, G, B float64
}

func NewColor(r,g,b float64) *Color {
	var color = new(Color)
		color.set(r, g, b)
	return color
}

func (color *Color) String() string {
	return fmt.Sprintf("rgb(%v, %v, %v)", math.Round(color.R), math.Round(color.G), math.Round(color.B))
}

func (color *Color) set(r, g, b float64) {
	color.R = color.clamp(r)
	color.G = color.clamp(g)
	color.B = color.clamp(b)
}

func (color *Color) hueRotate(angle float64) {
	angle = angle / 180.0 * math.Pi
	var sin = math.Sin(angle)
	var cos = math.Cos(angle)
	
	color.multiply(matrix{
		0.213 + cos * 0.787 - sin * 0.213,
		0.715 - cos * 0.715 - sin * 0.715,
		0.072 - cos * 0.072 + sin * 0.928,
		0.213 - cos * 0.213 + sin * 0.143,
		0.715 + cos * 0.285 + sin * 0.140,
		0.072 - cos * 0.072 - sin * 0.283,
		0.213 - cos * 0.213 - sin * 0.787,
		0.715 - cos * 0.715 + sin * 0.715,
		0.072 + cos * 0.928 + sin * 0.072,
	})
}

func (color *Color) grayscale(value float64) {
	color.multiply(matrix{
		0.2126 + 0.7874 * (1 - value),
		0.7152 - 0.7152 * (1 - value),
		0.0722 - 0.0722 * (1 - value),
		0.2126 - 0.2126 * (1 - value),
		0.7152 + 0.2848 * (1 - value),
		0.0722 - 0.0722 * (1 - value),
		0.2126 - 0.2126 * (1 - value),
		0.7152 - 0.7152 * (1 - value),
		0.0722 + 0.9278 * (1 - value),
	})
}

func (color *Color) sepia(value float64) {
	color.multiply(matrix{
      0.393 + 0.607 * (1 - value),
      0.769 - 0.769 * (1 - value),
      0.189 - 0.189 * (1 - value),
      0.349 - 0.349 * (1 - value),
      0.686 + 0.314 * (1 - value),
      0.168 - 0.168 * (1 - value),
      0.272 - 0.272 * (1 - value),
      0.534 - 0.534 * (1 - value),
      0.131 + 0.869 * (1 - value),
	})
}

func (color *Color) saturate(value float64) {
	color.multiply(matrix{
      0.213 + 0.787 * value,
      0.715 - 0.715 * value,
      0.072 - 0.072 * value,
      0.213 - 0.213 * value,
      0.715 + 0.285 * value,
      0.072 - 0.072 * value,
      0.213 - 0.213 * value,
      0.715 - 0.715 * value,
      0.072 + 0.928 * value,
	})
}

func (color *Color) multiply(m matrix) {
	var newR = color.clamp(color.R * m[0] + color.G * m[1] + color.B * m[2])
    var newG = color.clamp(color.R * m[3] + color.G * m[4] + color.B * m[5])
    var newB = color.clamp(color.R * m[6] + color.G * m[7] + color.B * m[8])
	color.set(newR, newG, newB)
}

func (color *Color) brightness(value float64) {
	color.linear(value, 0)
}

func (color *Color) contrast(value float64) {
	color.linear(value, -(0.5 * value) + 0.5)
}

func (color *Color) linear(slope, intercept float64) {
	color.R = color.clamp(color.R * slope + intercept * 255)
	color.G = color.clamp(color.G * slope + intercept * 255)
	color.B = color.clamp(color.B * slope + intercept * 255)
}

func (color *Color) invert(value float64) {
	color.R = color.clamp((value + color.R / 255 * (1 - 2 * value)) * 255)
    color.G = color.clamp((value + color.G / 255 * (1 - 2 * value)) * 255)
    color.B = color.clamp((value + color.B / 255 * (1 - 2 * value)) * 255)
}

func (color *Color) hsl() struct { H, S, L float64 } {
	var r = color.R / 255
	var g = color.G / 255
	var b = color.B / 255
	var max = math.Max(math.Max(r, g), b)
	var min = math.Min(math.Min(r, g), b)
	var avg = (max+min) / 2
	var h, s, l = 0.0, 0.0, avg
	
	if (max == min) {
		h = 0
		s = 0
	} else {
		var d = max - min
		if l > 0.5 {
			s = d/(2 - max - min)
		} else {
			s = d / (max + min)
		}
		switch (max) {
			case r:
				h = (g - b) / d
				if g < b {
					h += 6
				}
			case g:
				h = (b - r) / d + 2;
			case b:
				h = (r - g) / d + 4;
			default:
				panic("this should not happen "+fmt.Sprint(max, r, g, b))
		}
		h /= 6;
	}
	
	return struct { H, S, L float64 }{
		H: h * 100,
		S: s * 100,
		L: l * 100, 
	}
}

func (color *Color) clamp(value float64) float64 {
	if (value > 255) {
		value = 255;
	} else if (value < 0) {
		value = 0;
	}
	return value;
}

type Solver struct {
	target *Color
	targetHSL struct { H, S, L float64 }
	reusedColor *Color
}

func NewSolver(target *Color) *Solver {
	return &Solver{
		target: target,
		targetHSL: target.hsl(),
		reusedColor: NewColor(0,0,0),
	}
}

func (solver *Solver) Solve() (values [6]float64, loss float64, filter string) {
	values, loss = solver.solveNarrow(solver.solveWide())
	filter = solver.css(values)
	
	return values, loss, filter
}

func (solver *Solver) solveWide() (values [6]float64, loss float64) {
	var A = 5.0
	var c = 15.0
	var a = [6]float64{50, 180, 18000, 600, 1.2, 1.2}
	
	var bestLoss = math.MaxFloat64
	
	for i := 0; bestLoss > 25 && i < 3; i++ {
		var initial = [6]float64{50, 20, 3750, 50, 100, 100};
		var result_values, result_loss = solver.spsa(A, a, c, initial, 1000)
		if result_loss < bestLoss {
			bestLoss = result_loss
			values = result_values
		}
	}
	return values, bestLoss
}

func (solver *Solver) solveNarrow(wide_values [6]float64, wide_loss float64) (values [6]float64, loss float64) {
	var A = wide_loss
	var c = 2.0
	var A1 = A + 1
	var a = [6]float64{0.25 * A1, 0.25 * A1, A1, 0.25 * A1, 0.2 * A1, 0.2 * A1};
	return solver.spsa(A, a, c, wide_values, 500)
}

func (solver *Solver) spsa(A float64, a [6]float64, c float64, values [6]float64, iters int) (best [6]float64, loss float64) {	
	var alpha = 1.0
	var gamma = 0.16666666666666666
	
	var bestLoss = math.MaxFloat64
	
	var deltas [6]float64
	var highArgs [6]float64
	var lowArgs [6]float64
	
	var fix = func(value float64, idx int) float64 {
      var max = 100.0
      if idx == 2  { /* saturate */
        max = 7500
      } else if (idx == 4 /* brightness */ || idx == 5 /* contrast */) {
        max = 200
      }

      if (idx == 3 /* hue-rotate */) {
        if (value > max) {
          value = math.Mod(value, max)
        } else if (value < 0) {
          value = max + math.Mod(value, max)
        }
      } else if (value < 0) {
        value = 0
      } else if (value > max) {
        value = max
      }
      return value
    }
	
	for k := 0; k < iters; k++ {
		var ck = c / math.Pow(float64(k+1), gamma)
		for i := 0; i < 6; i++ {
			if random() > 0.5 {
				deltas[i] =  1
			} else {
				deltas[i] = -1
			}
			highArgs[i] = values[i] + ck * deltas[i];
			lowArgs[i] = values[i] - ck * deltas[i];
		}
		
		var lossDiff = solver.loss(highArgs) - solver.loss(lowArgs)
		for i := 0; i < 6; i++ {
			var g = lossDiff / (2 * ck) * deltas[i]
			var ak = a[i] / math.Pow(A + float64(k) + 1, alpha)
			values[i] = fix(values[i] - ak * g, i)
		}
		
		var loss = solver.loss(values)
		if (loss < bestLoss) {
			bestLoss = loss;
			best = values
		}
	}
	
	return best, bestLoss
}

func (solver *Solver) loss(filters [6]float64) float64 {
	var color = solver.reusedColor
	color.set(0,0,0)
	
	color.invert(filters[0] / 100)
	color.sepia(filters[0] / 100)
	color.saturate(filters[2] / 100)
	color.hueRotate(filters[3] * 3.6)
	color.brightness(filters[4] / 100)
	color.contrast(filters[5] / 100)
	
	var colorHSL = color.hsl()
	return math.Abs(color.R - solver.target.R) +
		math.Abs(color.G - solver.target.G) +
		math.Abs(color.B - solver.target.B) +
		math.Abs(colorHSL.H - solver.targetHSL.H) +
		math.Abs(colorHSL.S - solver.targetHSL.S) +
		math.Abs(colorHSL.L - solver.targetHSL.L)
}

func (solver *Solver) css(filters [6]float64) string {

	var f = func(idx int, multiplier float64) float64 {
      return math.Round(filters[idx] * multiplier);
    }

    return fmt.Sprint(`brightness(0) invert(`,f(0, 1),`%) sepia(`,f(1, 1),`%) saturate(`,f(2, 1),`%) hue-rotate(`,f(3, 3.6),`deg) brightness(`,f(4, 1),`%) contrast(`,f(5, 1),`%);`);
}

func init() {
	rand.Seed(time.Now().UnixNano())
	seed = rand.ExpFloat64() / 2
}

//Test
func test_colors() {
	
	var r, g, b float64 = 255, 0, 0
	
	var color = NewColor(r, g, b);
	var solver = NewSolver(color)
	
	start := time.Now()
	var _, loss, filter = solver.Solve()
	elapsed := time.Since(start)
	 fmt.Println("Binomial took ", elapsed)
	fmt.Println(loss)
	
	var lossMessage string;
    if (loss < 1) {
      lossMessage = "This is a perfect result.";
    } else if (loss < 5) {
      lossMessage = "The is close enough.";
    } else if (loss < 15) {
      lossMessage = "The color is somewhat off. Consider running it again.";
    } else {
      lossMessage = "The color is extremely off. Run it again!";
    }
    
	fmt.Println(filter)
    fmt.Println(lossMessage)
}
