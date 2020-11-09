package solution

func maxSatisfied(customers, grumps []int, x int) int {
	sum := 0
	for i := 0; i < len(customers); i++ {
		if grumps[i] != 0 {
			sum += customers[i]
		}
	}
	if x >= len(grumps) {
		return sum
	}
	var currentMax, current int
	currentMax = 0
	for i := 0; i < len(grumps)-x; i++ {
		current = 0
		for j := i; j < x+i; j++ {
			if grumps[j] == 0 {
				current += customers[j]
			}
		}
		if current > currentMax {
			currentMax = current
		}
	}
	return sum + currentMax
}
