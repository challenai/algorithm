package solution

func reverseBits(num int) int {
	resu := 0
	cntr := 31
	for cntr > 0 {
		resu += num % 2
		resu <<= 1
		num >>= 1
		cntr--
	}
	if num%2 == 1 {
		resu++
	}
	return resu
}
