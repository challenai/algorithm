package solution

import (
	"math"
)

func ConstructRectangle(target int) []int {
	rsu := make([]int, 2)
	rsu[0] = int(math.Sqrt(float64(target)))
	for i := rsu[0]; i > 1; i-- {
		if target%rsu[0] == 0 {
			rsu[1] = target / rsu[0]
			return rsu
		}
	}
	rsu[0] = 1
	rsu[1] = target
	return rsu
}
