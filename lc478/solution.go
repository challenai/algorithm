package solution

import (
	"math"
	"math/rand"
	"time"
)

type Solution struct {
	r, x, y float64
	ra      *rand.Rand
}

func Constructor(radius float64, x_center float64, y_center float64) Solution {
	return Solution{
		r:  radius,
		x:  x_center,
		y:  y_center,
		ra: rand.New(rand.NewSource(time.Now().Unix())),
	}
}

func (s *Solution) RandPoint() []float64 {
	x := s.x + s.r
	y := s.y + s.r
	for math.Exp2(x-s.x)+math.Exp2(y-s.y) > math.Exp2(s.r) {
		x = s.x - s.r + s.ra.Float64()*s.r*2
		// println(rand.Float64())
		y = s.y - s.r + s.ra.Float64()*s.r*2
		// println(x, y)
	}
	return []float64{x, y}

	// rand.Seed(time.Now().UnixNano())
	// left := s.x - s.r
	// x := rand.Float64()*(s.r*2) + left
	// sqrt := math.Sqrt(math.Pow(s.r, 2) - math.Pow(x-s.x, 2))
	// y := rand.Float64()*(sqrt*2) - sqrt + s.y
	// return []float64{x, y}
}
