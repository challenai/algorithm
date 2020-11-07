package solution

func hammingDistance(x, y int) int {
	rsu := 0
	for x > 0 || y > 0 {
		if x&1 != y&1 {
			rsu++
		}
		x >>= 1
		y >>= 1
	}
	return rsu
}
