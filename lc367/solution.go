package solution

func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	current := 2
	for current*current <= num {
		current <<= 1
	}
	if current*current == num {
		return true
	}
	current >>= 1
	target := current
	println(current)
	for current > 0 {
		current >>= 1
		if (target+current)*(target+current) <= num {
			target += current
		}
	}
	println(target * target)
	return target*target == num
}
