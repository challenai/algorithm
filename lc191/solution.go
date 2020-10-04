package solution

func hammingWeight(num int) int {
	rsu := 0
	for num > 0 {
		if num%2 == 1 {
			rsu++
		}
		num >>= 1
	}
	return rsu
}
