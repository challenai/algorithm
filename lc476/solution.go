package solution

func findComplement(num int) int {
	rsu := 0
	for num > 0 {
		if num%2 == 0 {
			rsu += 1
		}
		num >>= 1
		rsu <<= 1
	}
	rsu >>= 1
	return rsu
}
