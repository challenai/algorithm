package solution

func findRadius(house, heater []int) int {
	minRadius := 0
	ptr := 0
	for i := 0; i < len(house); i++ {
		if ptr+1 < len(heater) && abs(heater[ptr+1]-house[i]) < abs(heater[ptr]-house[i]) {
			minRadius = max(minRadius, abs(heater[ptr+1]-house[i]))
		} else {
			minRadius = max(minRadius, abs(heater[ptr]-house[i]))
		}
	}
	return minRadius
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
