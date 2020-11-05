package solution

func arrangeCoins(n int) int {
	if n == 0 {
		return 0
	}
	index := 1
	var layer int
	for index*(index+1)>>1 < n {
		index <<= 1
	}
	if index*(index+1) == n {
		return index
	}
	index >>= 1
	layer = index
	for index > 0 {
		index >>= 1
		if (layer+index)*(layer+index+1)>>1 < n {
			layer += index
		}
	}
	return layer
}
